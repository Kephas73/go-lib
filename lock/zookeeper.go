package lock

import (
    "fmt"
    "github.com/samuel/go-zookeeper/zk"
    "time"
)

type Zookeeper struct {
    *zk.Conn
}

var zookeeper *Zookeeper

func ConnectZookeeper() *Zookeeper {

    conn, _, err := zk.Connect(config.Addrs, time.Second, zk.WithLogInfo(true))
    if err != nil {
        panic(err)
    }

    if conn == nil {
        err = fmt.Errorf("need initialization zookeeper global instance connection first")
        panic(err)
    }

    zookeeper = &Zookeeper{
        conn,
    }

    return zookeeper
}

func InstanceZookeeper(envKey ...string) {
    if config == nil {
        createZookeeperConfigFromEnv(envKey...)
    }

    if config == nil {
        err := fmt.Errorf("not found config for zookeeper")
        panic(err)
    }

    ConnectZookeeper()
}

func GetInstanceZookeeper() *Zookeeper {
    if zookeeper == nil {
        zookeeper = ConnectZookeeper()
    }

    return zookeeper
}
