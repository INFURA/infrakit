package opts

type ProxyCmdConfig struct {
	ListenAddr           string `koanf:"listen_addr"`
	HealthListenAddr     string `koanf:"health_listen_addr"`
	EthBootstrapEndpoint string `koanf:"eth_bootstrap_endpoint"`
}

var (
	proxyConfig ProxyCmdConfig
)

func ProxyConfig() *ProxyCmdConfig {
	return &proxyConfig
}
