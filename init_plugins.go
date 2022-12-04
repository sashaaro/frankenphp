package frankenphp

import (
	"github.com/dunglas/frankenphp/plugins"
	"github.com/dunglas/frankenphp/plugins/httpclient" // TODO code generation
	"go.uber.org/zap"
)

func initPlugins(logger *zap.Logger, options ...Option) error {
	includedModules := []plugins.FrankenPhpPlugin{
		httpclient.Plugin{}, // TODO code generation while compile if module specified
	}

	for _, module := range includedModules {
		err := module.Init(options) // TODO
		if err != nil {
			return err
		}
		plugins.FrankenPhpPlugins[module.Name()] = module
	}

	return nil
}
