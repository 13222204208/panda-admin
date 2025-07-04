package curd

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"golang.org/x/tools/imports"
)

func GetTempGeneratePath() string {
	return gfile.Abs(gfile.Temp() + "/hotgo-generate/")
}

func FormatGo(name, code string) (string, error) {
	path := GetTempGeneratePath() + "/" + name
	if err := gfile.PutContents(path, code); err != nil {
		return "", err
	}
	res, err := imports.Process(path, []byte(code), nil)
	if err != nil {
		err = gerror.Newf(`FormatGo error format "%s" go files: %v`, path, err)
		return "", err
	}
	return string(res), nil
}
