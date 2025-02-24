package main

import (
	"os"

	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/provider/manager"
	"github.com/hashicorp/go-hclog"
	hc_plugin "github.com/hashicorp/go-plugin"

	p "github.com/daytonaio/daytona-provider-hetzner/pkg/provider"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	hc_plugin.Serve(&hc_plugin.ServeConfig{
		HandshakeConfig: manager.ProviderHandshakeConfig,
		Plugins: map[string]hc_plugin.Plugin{
			"hetzner-provider": &provider.ProviderPlugin{Impl: &p.HetznerProvider{}},
		},
		Logger: logger,
	})
}
