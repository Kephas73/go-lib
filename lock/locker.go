package lock

import (
    "fmt"
    "github.com/samuel/go-zookeeper/zk"
)

const (
    KLockerNamespacePrefix = "lock"
    KNameProjectDir        = "go-lib"
)

func (g *Zookeeper) NewLocker(key string) *zk.Lock {
    keyLocker := fmt.Sprintf("/%s/%s", KLockerNamespacePrefix, key)
    return zk.NewLock(g.Conn, keyLocker, zk.WorldACL(zk.PermAll))

}
