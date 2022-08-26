package lock

import (
    "fmt"
    "github.com/spf13/viper"
)

type Config struct {
    Addrs   []string `json:"addrs"`
    Timeout int
}

var config *Config

const (
    TimeoutDefault = 20
)

func createZookeeperConfigFromEnv(envKey ...string) {
    config = &Config{}

    configKey := "Zookeeper"
    if len(envKey) > 0 {
        configKey = envKey[0]
    }

    if err := viper.UnmarshalKey(configKey, config); err != nil {
        err := fmt.Errorf("not found config name with env %q for Zookeeper with error: %+v", configKey, err)
        panic(err)
    }

    if len(config.Addrs) == 0 {
        err := fmt.Errorf("not found any addr as host for Zookeeper at %q", fmt.Sprintf("%s.addrs", configKey))
        panic(err)
    }

    if config.Timeout <= 0 {
        config.Timeout = TimeoutDefault
    }
}

/*func locker() {
    lk, _, _ := zk.Connect([]string{"127.0.0.1"}, time.Second)

    zk.NewLock()

    lk.State()
    l := zk.NewLock()
    l.Lock()
}
*/