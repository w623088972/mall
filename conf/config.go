package conf

import (
	"log"
	"strings"

	"github.com/beijibeijing/viper"
	"github.com/radovskyb/watcher"
)

//Config 结构体
type Config struct {
	Name string
}

//ConfigInit 初始化
func ConfigInit(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	log.Printf("ConfigInit")
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("file") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatcherConfig()
	viper.OnWatcherConfigChange(func(e watcher.Event) {
		log.Printf("Config file changed: ", e.FileInfo)
	})
}

//BackslashDone 处理\\n问题
func BackslashDone(str string) string {
	return strings.Replace(str, "\\n", "\n", 1)
}
