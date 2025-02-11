package plugins

import "github.com/spf13/viper"

type PluginInterface interface {
	Init(env *viper.Viper) error
	Call()
}
