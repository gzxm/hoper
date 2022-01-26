package local

import (
	"fmt"
	"github.com/actliboy/hoper/server/go/lib/utils/configor"
	"github.com/actliboy/hoper/server/go/lib/utils/fs"
)

type Local struct {
	configor.Config
	LocalConfigName string
}

// 本地配置
func (cc *Local) HandleConfig(handle func([]byte)) error {
	localConfigName := cc.LocalConfigName
	if localConfigName != "" {
		adCongPath, err := fs.FindFile(localConfigName)
		localConfigName = adCongPath
		if err == nil {
			err := configor.New(&cc.Config).
				Handle(handle, adCongPath)
			if err != nil {
				return fmt.Errorf("配置错误: %v", err)
			}
		} else {
			return fmt.Errorf("找不到配置: %v", err)
		}
	}
	return nil
}