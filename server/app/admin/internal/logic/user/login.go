package user

import (
	"context"
	"time"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/internal/dao"
	"server/app/admin/internal/middleware"
	"server/app/admin/internal/model/entity"
	"server/utility"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mojocn/base64Captcha"
)

const (
	// TokenExpireTime Token过期时间
	TokenExpireTime = 2 * time.Hour
	// RefreshTokenExpireTime 刷新Token过期时间
	RefreshTokenExpireTime = 30 * 24 * time.Hour
)

type sLogin struct {
	captcha *base64Captcha.Captcha
}

func NewLogin() *sLogin {
	// 初始化验证码驱动
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	store := base64Captcha.DefaultMemStore
	return &sLogin{
		captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

// GetCaptcha 获取验证码
func (s *sLogin) GetCaptcha(ctx context.Context, req v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	res = &v1.CaptchaRes{}

	// 生成验证码
	driver := s.captcha.Driver.(*base64Captcha.DriverDigit)
	driver.Height = 80
	driver.Width = 240
	driver.Length = 4
	driver.MaxSkew = 0.7
	driver.DotCount = 80

	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, content, _, err := c.Generate()
	if err != nil {
		return nil, gerror.Wrap(err, "生成验证码失败")
	}

	res.CaptchaId = id
	res.CaptchaImg = content
	return
}

// Login 用户登录
func (s *sUser) Login(ctx context.Context, in v1.LoginReq) (out *v1.LoginRes, err error) {
	// 验证码校验
	if !base64Captcha.DefaultMemStore.Verify(in.CaptchaId, in.VerifyCode, true) {
		return nil, gerror.New("验证码错误")
	}
	out = &v1.LoginRes{}

	// 查询用户
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username, in.Username).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "用户名或密码错误")
	}

	// 检查用户状态
	if user.Status == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "用户已被禁用")
	}

	// 验证密码
	if err := utility.ComparePassword(user.Password, in.Password); err != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "用户名或密码错误")
	}

	// 生成访问令牌
	accessToken, err := s.generateToken(user.Id, user.Username, TokenExpireTime)
	if err != nil {
		return nil, gerror.Wrap(err, "生成访问令牌失败")
	}

	// 生成刷新令牌
	refreshToken, err := s.generateToken(user.Id, user.Username, RefreshTokenExpireTime)
	if err != nil {
		return nil, gerror.Wrap(err, "生成刷新令牌失败")
	}

	// 获取用户角色和权限
	var roles []string
	var permissions []string

	// 如果用户名是 developer，直接给予所有权限
	if in.Username == "developer" {
		roles = []string{"developer"}
		permissions = []string{"*:*:*"}
	} else {
		// 获取用户角色和权限
		var err error
		roles, permissions, err = s.getUserRolesAndPermissions(ctx, user.Id)
		if err != nil {
			return nil, gerror.Wrap(err, "获取用户权限失败")
		}
	}

	// 设置返回数据
	out.Avatar = user.Avatar
	out.Username = user.Username
	out.Nickname = user.Nickname
	out.Roles = roles
	out.Permissions = permissions
	out.AccessToken = accessToken
	out.RefreshToken = refreshToken
	out.Expires = gtime.New(time.Now().Add(TokenExpireTime))

	return
}

// RefreshToken 刷新令牌
func (s *sUser) RefreshToken(ctx context.Context, in v1.RefreshTokenReq) (out *v1.RefreshTokenRes, err error) {
	out = &v1.RefreshTokenRes{}

	// 验证刷新令牌
	claims, err := s.parseToken(in.RefreshToken)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "无效的刷新令牌")
	}

	// 检查用户是否存在且状态正常
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, claims.UserID).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "查询用户失败")
	}
	if user == nil || user.Status == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "用户不存在或已被禁用")
	}

	// 生成新的访问令牌
	accessToken, err := s.generateToken(user.Id, user.Username, TokenExpireTime)
	if err != nil {
		return nil, gerror.Wrap(err, "生成访问令牌失败")
	}

	// 生成新的刷新令牌
	refreshToken, err := s.generateToken(user.Id, user.Username, RefreshTokenExpireTime)
	if err != nil {
		return nil, gerror.Wrap(err, "生成刷新令牌失败")
	}

	out.AccessToken = accessToken
	out.RefreshToken = refreshToken
	out.Expires = gtime.New(time.Now().Add(TokenExpireTime))

	return
}

// generateToken 生成JWT令牌
func (s *sUser) generateToken(userID uint64, username string, expireTime time.Duration) (string, error) {
	claims := &middleware.JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(middleware.JwtSecretKey))
}

// parseToken 解析JWT令牌
func (s *sUser) parseToken(tokenString string) (*middleware.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &middleware.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*middleware.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, gerror.New("无效的令牌")
}

// getUserRolesAndPermissions 获取用户角色和权限
func (s *sUser) getUserRolesAndPermissions(ctx context.Context, userID uint64) ([]string, []string, error) {
	// 获取用户的角色列表
	roles, err := s.getUserRoles(ctx, userID)
	if err != nil {
		return nil, nil, gerror.Wrap(err, "获取用户角色失败")
	}

	// 获取角色对应的权限列表
	permissions, err := s.getRolePermissions(ctx, roles)
	if err != nil {
		return nil, nil, gerror.Wrap(err, "获取角色权限失败")
	}

	// 提取角色编码和权限标识
	var roleCodes []string
	var permissionCodes []string

	for _, role := range roles {
		roleCodes = append(roleCodes, role.Code)
	}

	for _, permission := range permissions {
		// 如果菜单有权限标识，则添加到权限列表
		if permission.Auths != "" {
			// 权限标识可能是多个，用逗号分隔
			auths := gstr.Split(permission.Auths, ",")
			for _, auth := range auths {
				auth = gstr.Trim(auth)
				if auth != "" {
					permissionCodes = append(permissionCodes, auth)
				}
			}
		}
	}

	// 去重权限标识
	permissionCodes = garray.NewStrArrayFrom(permissionCodes).Unique().Slice()

	// 如果是超级管理员，给予所有权限
	for _, roleCode := range roleCodes {
		if roleCode == "developer" {
			permissionCodes = []string{"*:*:*"}
			break
		}
	}

	return roleCodes, permissionCodes, nil
}

// getUserRoles 获取用户角色列表
func (s *sUser) getUserRoles(ctx context.Context, userID uint64) ([]entity.Role, error) {
	var roles []entity.Role

	// 通过用户角色关联表和角色表联查获取用户的角色信息
	err := dao.UserRole.Ctx(ctx).
		LeftJoin(dao.Role.Table(), dao.UserRole.Columns().RoleId+"="+dao.Role.Columns().Id).
		Where(dao.UserRole.Columns().UserId, userID).
		Where(dao.Role.Columns().Status, 1). // 只获取启用的角色
		Fields(dao.Role.Columns().Id, dao.Role.Columns().Name, dao.Role.Columns().Code, dao.Role.Columns().Status).
		Scan(&roles)

	if err != nil {
		return nil, gerror.Wrap(err, "查询用户角色失败")
	}

	return roles, nil
}

// getRolePermissions 获取角色对应的权限菜单列表
func (s *sUser) getRolePermissions(ctx context.Context, roles []entity.Role) ([]entity.Menu, error) {
	if len(roles) == 0 {
		return []entity.Menu{}, nil
	}

	// 提取角色ID列表
	var roleIds []uint64
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}

	var menus []entity.Menu

	// 通过角色菜单关联表和菜单表联查获取角色的权限菜单
	err := dao.RoleMenu.Ctx(ctx).
		LeftJoin(dao.Menu.Table(), dao.RoleMenu.Columns().MenuId+"="+dao.Menu.Columns().Id).
		WhereIn(dao.RoleMenu.Columns().RoleId, roleIds).
		Where(dao.Menu.Columns().Id+" IS NOT NULL"). // 确保菜单存在
		Fields(dao.Menu.Columns().Id, dao.Menu.Columns().Title, dao.Menu.Columns().Name, dao.Menu.Columns().Auths, dao.Menu.Columns().MenuType).
		Group(dao.Menu.Columns().Id). // 去重
		Scan(&menus)

	if err != nil {
		return nil, gerror.Wrap(err, "查询角色权限失败")
	}

	return menus, nil
}
