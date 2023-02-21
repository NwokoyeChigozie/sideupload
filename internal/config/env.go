package config

import "encoding/json"

type Configuration struct {
	Server ServerConfiguration
	App    App
	AWSS3  AWSS3
}
type BaseConfig struct {
	SERVER_PORT                      string  `mapstructure:"SERVER_PORT"`
	SERVER_SECRET                    string  `mapstructure:"SERVER_SECRET"`
	SERVER_ACCESSTOKENEXPIREDURATION int     `mapstructure:"SERVER_ACCESSTOKENEXPIREDURATION"`
	REQUEST_PER_SECOND               float64 `mapstructure:"REQUEST_PER_SECOND"`
	TRUSTED_PROXIES                  string  `mapstructure:"TRUSTED_PROXIES"`
	EXEMPT_FROM_THROTTLE             string  `mapstructure:"EXEMPT_FROM_THROTTLE"`

	APP_NAME string `mapstructure:"APP_NAME"`
	APP_KEY  string `mapstructure:"APP_KEY"`

	AWSS3_ACCESS_KEY_ID     string `mapstructure:"AWSS3_ACCESS_KEY_ID"`
	AWSS3_SECRET_ACCESS_KEY string `mapstructure:"AWSS3_SECRET_ACCESS_KEY"`
	AWSS3_DEFAULT_REGION    string `mapstructure:"AWSS3_DEFAULT_REGION"`
	AWSS3_BUCKET            string `mapstructure:"AWSS3_BUCKET"`
}

func (config *BaseConfig) SetupConfigurationn() *Configuration {
	trustedProxies := []string{}
	exemptFromThrottle := []string{}
	json.Unmarshal([]byte(config.TRUSTED_PROXIES), &trustedProxies)
	json.Unmarshal([]byte(config.EXEMPT_FROM_THROTTLE), &exemptFromThrottle)
	return &Configuration{
		Server: ServerConfiguration{
			Port:                      config.SERVER_PORT,
			Secret:                    config.SERVER_SECRET,
			AccessTokenExpireDuration: config.SERVER_ACCESSTOKENEXPIREDURATION,
			RequestPerSecond:          config.REQUEST_PER_SECOND,
			TrustedProxies:            trustedProxies,
			ExemptFromThrottle:        exemptFromThrottle,
		},
		App: App{
			Name: config.APP_NAME,
			Key:  config.APP_KEY,
		},

		AWSS3: AWSS3{
			AccessKeyId:     config.AWSS3_ACCESS_KEY_ID,
			SecretAccessKey: config.AWSS3_SECRET_ACCESS_KEY,
			DefaultRegion:   config.AWSS3_DEFAULT_REGION,
			Bucket:          config.AWSS3_BUCKET,
		},
	}
}
