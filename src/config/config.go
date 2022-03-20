package config

type Config struct {
	ApiKey string
	ApiUrl string
	Format string
}

var Configuration = &Config{
	ApiKey: "valid_api_key",
	ApiUrl: "https://www.whoisxmlapi.com/whoisserver/WhoisService?",
	Format: "json",
}

func GenerateConfig() *Config {
	return Configuration
}
