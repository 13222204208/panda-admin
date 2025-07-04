import { http } from "@/utils/http";
import { baseUrlApi } from "./common/utils";
import type { BaseResponse } from "./common/types";

/** 用户登录响应数据 */
export interface LoginData {
  /** 头像 */
  avatar: string;
  /** 用户名 */
  username: string;
  /** 昵称 */
  nickname: string;
  /** 当前登录用户的角色 */
  roles: Array<string>;
  /** 按钮级别权限 */
  permissions: Array<string>;
  /** 访问令牌 */
  accessToken: string;
  /** 用于调用刷新`accessToken`的接口时所需的`token` */
  refreshToken: string;
  /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
  expires: Date;
}

/** 用户登录响应类型 */
export type UserResult = BaseResponse<LoginData>;

/** 刷新令牌响应数据 */
export interface RefreshTokenData {
  /** 访问令牌 */
  accessToken: string;
  /** 用于调用刷新`accessToken`的接口时所需的`token` */
  refreshToken: string;
  /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
  expires: Date;
}

/** 刷新令牌响应类型 */
export type RefreshTokenResult = BaseResponse<RefreshTokenData>;

/** 验证码响应数据 */
export interface CaptchaData {
  /** 验证码ID */
  captchaId: string;
  /** 验证码图片 */
  captchaImg: string;
}

/** 验证码响应类型 */
export type CaptchaResult = BaseResponse<CaptchaData>;

/** 登录接口 */
export const getLogin = (data?: object) => {
  return http.request<UserResult>("post", baseUrlApi("login"), { data });
};

/** 刷新`token` */
export const refreshTokenApi = (data?: object) => {
  return http.request<RefreshTokenResult>("post", baseUrlApi("refresh-token"), {
    data
  });
};

/** 获取验证码 */
export const getCaptcha = () => {
  return http.request<CaptchaResult>("get", baseUrlApi("captcha"));
};
