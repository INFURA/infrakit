package opts

type ProxyCmdConfig struct {
	ListenAddr       string `koanf:"listen_addr"`
	HealthListenAddr string `koanf:"health_listen_addr"`
}

var (
	proxyConfig ProxyCmdConfig
)

func ProxyConfig() *ProxyCmdConfig {
	return &proxyConfig
}
