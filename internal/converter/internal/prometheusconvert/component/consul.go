package component

import (
	"time"

	"github.com/grafana/alloy/internal/component/discovery"
	"github.com/grafana/alloy/internal/component/discovery/consul"
	"github.com/grafana/alloy/internal/converter/diag"
	"github.com/grafana/alloy/internal/converter/internal/common"
	"github.com/grafana/alloy/internal/converter/internal/prometheusconvert/build"
	"github.com/grafana/alloy/syntax/alloytypes"
	prom_consul "github.com/prometheus/prometheus/discovery/consul"
)

func appendDiscoveryConsul(pb *build.PrometheusBlocks, label string, sdConfig *prom_consul.SDConfig) discovery.Exports {
	discoveryConsulArgs := toDiscoveryConsul(sdConfig)
	name := []string{"discovery", "consul"}
	block := common.NewBlockWithOverride(name, label, discoveryConsulArgs)
	pb.DiscoveryBlocks = append(pb.DiscoveryBlocks, build.NewPrometheusBlock(block, name, label, "", ""))
	return common.NewDiscoveryExports("discovery.consul." + label + ".targets")
}

func ValidateDiscoveryConsul(sdConfig *prom_consul.SDConfig) diag.Diagnostics {
	return common.ValidateHttpClientConfig(&sdConfig.HTTPClientConfig)
}

func toDiscoveryConsul(sdConfig *prom_consul.SDConfig) *consul.Arguments {
	if sdConfig == nil {
		return nil
	}

	return &consul.Arguments{
		Server:           sdConfig.Server,
		Token:            alloytypes.Secret(sdConfig.Token),
		Datacenter:       sdConfig.Datacenter,
		Namespace:        sdConfig.Namespace,
		Partition:        sdConfig.Partition,
		TagSeparator:     sdConfig.TagSeparator,
		Scheme:           sdConfig.Scheme,
		Username:         sdConfig.Username,
		Password:         alloytypes.Secret(sdConfig.Password),
		AllowStale:       sdConfig.AllowStale,
		Services:         sdConfig.Services,
		ServiceTags:      sdConfig.ServiceTags,
		NodeMeta:         sdConfig.NodeMeta,
		RefreshInterval:  time.Duration(sdConfig.RefreshInterval),
		HTTPClientConfig: *common.ToHttpClientConfig(&sdConfig.HTTPClientConfig),
	}
}
