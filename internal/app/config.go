package apiserver

// grave accent is used for struct tags
// this config immediately allocated in memory with initialization
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// default config
func NewConfig() *Config {
	conf := Config{
		BindAddr: ":8080",
		LogLevel: "debug"}
	return &conf
}
