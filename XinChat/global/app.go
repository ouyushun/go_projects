package global

import (
	"XinChat/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Application 定义 Application 结构体，用来存放一些项目启动时的变量，便于调用
type Application struct {
	ConfigViper *viper.Viper
	Config config.Configuration
	Log *zap.Logger
}

var App = new(Application)


/*
作者：jassue
链接：https://juejin.cn/post/7018139911043678245
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */