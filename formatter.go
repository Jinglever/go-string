package jgstr

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/tools/imports"
)

// 格式化golang代码，包括import
func FormatGo(src []byte) ([]byte, error) {
	formatted, err := imports.Process("", src, nil)
	if err != nil {
		logrus.Errorf("format source failed: %v", err)
		return nil, err
	}
	return formatted, nil
}
