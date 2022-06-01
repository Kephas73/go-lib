package opensearch

import (
    "fmt"
    "github.com/Kephas73/go-lib/constant"
    "github.com/spf13/viper"
    "strings"
)

type OpenSearchConfig struct {
    Hosts       []string
    Username    string
    Password    string
    IndexFormat string
}

var openSearchConf *OpenSearchConfig

func createConfigFromEnv(configKeys ...string) {
    configKey := "OpenSearch"
    for _, envKey := range configKeys {
        envKeyTrim := strings.TrimSpace(envKey)
        if envKeyTrim != "" {
            configKey = envKeyTrim
        }
    }

    openSearchConf = &OpenSearchConfig{}

    if err := viper.UnmarshalKey(configKey, openSearchConf); err != nil {
        err := fmt.Errorf("not found config name with env %q for open search with error: %+v", configKey, err)
        panic(err)
    }

    if len(openSearchConf.Hosts) == 0 {
        err := fmt.Errorf("not found hosts for open search with env %q", fmt.Sprintf("%s.Hosts", configKey))
        panic(err)
    }

    if openSearchConf.Username == "" {
        openSearchConf.Username = constant.UsernameDefault
    }

    if openSearchConf.Password == "" {
        openSearchConf.Password = constant.PasswordDefault
    }

    if openSearchConf.IndexFormat == "" {
        openSearchConf.IndexFormat = constant.IndexDefault
    }

}