package plugins

import (
	"fmt"
	"os"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type PluginSubsystem struct{}

type PluginInterface interface {
	Call() error
}

func New() (*PluginSubsystem, error) {
	return &PluginSubsystem{}, nil
}

func (ps *PluginSubsystem) Call() error {
	i := interp.New(interp.Options{
		GoPath: "./local_plugins/",
		Env:    os.Environ(),
	})

	err := i.Use(stdlib.Symbols)
	if err != nil {
		return err
	}

	symbols := map[string]map[string]reflect.Value{
		"github.com/ibm/opentalaria/plugins/plugins": {
			"PluginInterface": reflect.ValueOf((*PluginInterface)(nil)),
		},
	}

	err = i.Use(symbols)
	if err != nil {
		return err
	}

	_, err = i.Eval(`
		package wrapper

		import (
			"log"

			"github.com/ibm/opentalaria/demo"
			"github.com/ibm/opentalaria/plugins"
		)

		func NewWrapper() (plugins.PluginInterface, error) {
			p, err := demo.New()
			
			var pv plugins.PluginInterface = p
			return pv, err
		}
	`)
	if err != nil {
		return fmt.Errorf("error evaluating wrapper: %v", err)
	}

	// v, err := i.Eval(`wrapper.NewWrapper`)
	// if err != nil {
	// 	return fmt.Errorf("error calling NewWrapper: %v", err)
	// }

	// results := v.Call(nil)
	// log.Printf("%T", results[0].Interface())

	// plugin, ok := results[0].Interface().(PluginInterface)
	// if !ok {
	// 	return fmt.Errorf("invalid plugin type: %T", results[0].Interface())
	// }

	// log.Printf("Plugin: %T", plugin)

	return nil
}
