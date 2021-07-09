package config

import (
	"HELLO-GO/constant"
	"HELLO-GO/utility/logger"
	"encoding/json"
	"strings"

	"github.com/spf13/viper"
)

var (
	// AppConfig set the data from property file
	AppConfig          applicationConfig
	HttpConfigProperty httpConfigProperty
	log                = logger.GetLogger()
	Environment        = constant.Development
	validEnvs          = map[string]constant.Environment{
		"dev":     constant.Development,
		"qa":      constant.QA,
		"staging": constant.Staging,
		"prod":    constant.Production,
	}
)

func Initialize(configPath, env string) {
	if val, ok := validEnvs[strings.ToLower(env)]; ok {
		Environment = val
	} else {
		log.Warningf("Environment variable '%s' is unset or is invalid.Please set it to one of DEV, QA, STAGING, PROD. Using 'DEV' as default.", env)
		Environment = constant.Development
	}
	if err := readConfig("yaml", configPath, "configuration."+strings.ToLower(string(Environment)), &AppConfig, &HttpConfigProperty); err != nil {
		panic(err)
	}
	log.Info("ConfigAPITimeoutMillis :", HttpConfigProperty.ConfigAPITimeoutMillis)
}

func readConfig(cfgType, pathOrURL, name string, appConfig interface{}, httpConfigProperty interface{}) error {

	//read in properties from configuration.x.yaml
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType(cfgType)
	v.AddConfigPath(pathOrURL)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(appConfig); err != nil {
		return err
	}
	if err := v.Unmarshal(httpConfigProperty); err != nil {
		return err
	}
	log.Warningf("AppConfig at startup %s", prettyJson(v.AllSettings()))
	return nil
}

func prettyJson(settings map[string]interface{}) string {
	b, _ := json.MarshalIndent(settings, "", "  ")
	return string(b)
}
