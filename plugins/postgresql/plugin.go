package postgresql

import (
	"log"

	"github.com/spf13/viper"
)

type Plugin struct {
	config config
}

type config struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (p *Plugin) Init(env *viper.Viper) error {
	parsedConf := make(map[string]*config)
	err := env.UnmarshalKey("plugins", &parsedConf)
	if err != nil {
		return err
	}

	p.config = *parsedConf["postgres"]

	return nil
}

func (p *Plugin) Call() {
	log.Println("===========> Call from plugin")
}
