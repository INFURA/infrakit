package opts

import "sigs.k8s.io/controller-runtime/pkg/log/zap"

type LogConfig struct {
	Development bool `koanf:"dev"`
}

func FromLogConfig(opts LogConfig) zap.Options {
	var o zap.Options

	o.Development = opts.Development

	return o
}
