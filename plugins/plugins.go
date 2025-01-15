package plugins

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type PluginSubsystem struct {
}

type PluginInterface interface {
	Call() error
}

func New() (*PluginSubsystem, error) {
	return &PluginSubsystem{}, nil
}

func (ps *PluginSubsystem) Call() error {
	i := interp.New(interp.Options{
		GoPath: "./.plugins/",
		Env:    os.Environ(),
	})

	err := i.Use(stdlib.Symbols)
	if err != nil {
		return err
	}

	symbols := map[string]map[string]reflect.Value{
		"github.com/ibm/opentalaria/plugins": {
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
			"github.com/ibm/opentalaria/demo"
			"github.com/ibm/opentalaria/plugins"
		)

		func NewWrapper() (plugins.PluginInterface, error) {
			p, err := demo.New()
			var plugin plugins.PluginInterface = p

			return plugin, err
		}
	`)
	if err != nil {
		return err
	}

	v, err := i.Eval(`wrapper.NewWrapper`)
	if err != nil {
		return err
	}

	results := v.Call(nil)

	log.Printf("%T\n", results[0].Interface())

	plugin, ok := results[0].Interface().(PluginInterface)
	if !ok {
		return fmt.Errorf("invalid plugin type: %T", results[0].Interface())
	}

	log.Println(plugin)

	// _, err = i.Eval(string(plugin))
	// if err != nil {
	// 	return err
	// }

	// log.Println(string(plugin))

	// 	_, err = i.Eval(`d, _ := demo.New()

	// d.Call()`)
	// 	if err != nil {
	// 		return err
	// 	}

	// v, err := i.Eval(`demo.New`)
	// if err != nil {
	// 	return err
	// }

	// result := v.Call(nil)

	// log.Println(result)

	// p := result[0].Interface().(PluginInterface)
	// p.Call()

	return nil
}
