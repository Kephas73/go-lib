package lock_etcd

import (
    "fmt"
    "github.com/Kephas73/go-lib/logger"
    "time"

    "github.com/spf13/viper"
    clientv3 "go.etcd.io/etcd/client/v3"
)

const (
    KNameProjectDir = "go_lib"
    LDefaultTLL     = 30
)

// ErrorWatcherClosed var
var ErrorWatcherClosed = fmt.Errorf("naming: watch closed")

type Config struct {
    Addrs       []string     `json:"addrs"`
    TTL         int          `json:"ttl"`
    ProjectName string       `json:"project_name"`
    Username    string       `json:"username"`
    Password    string       `json:"password"`
    LogLevel    EtcdLogLevel `json:"log_level"`
}

// Option type
type Option struct {
    Config      clientv3.Config
    RegistryDir string
    ServiceName string
    ServerID    string
    NodeData    NodeData
    TTL         time.Duration
}

// NodeData type
type NodeData struct {
    Addr     string
    Metadata interface{}
}

var config *Config

func createEtcdConfigFromEnv(envKey ...string) {
    config = &Config{}

    key := "Etcd"
    if len(envKey) > 0 {
        key = envKey[0]
    }

    if err := viper.UnmarshalKey(key, config); err != nil {
        logger.Error("createEtcdConfigFromEnv - Error: %v from key %q", err, key)
        panic(err)
    }

    if len(config.Addrs) == 0 {
        err := fmt.Errorf("not found any addr as host for etcd at %q", fmt.Sprintf("%s.addrs", key))
        panic(err)
    }

    if config.TTL <= 0 {
        config.TTL = LDefaultTLL
    }

    if config.ProjectName == "" {
        config.ProjectName = KNameProjectDir
    }
}
