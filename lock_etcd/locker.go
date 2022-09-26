package lock_etcd

import (
    "context"
    "fmt"
    "github.com/Kephas73/go-lib/logger"
    "sync"
    "time"

    "go.etcd.io/etcd/client/v3/concurrency"
)

const (
    KDefaultLockTimeout    = 10 * time.Second
    KLockerNamespacePrefix = "lock"
)

// GLocker type
type GLocker struct {
    *GEtcd
    *concurrency.Mutex
    keyLocker         string
    LockerActionError error
    CancelFunc        context.CancelFunc
}

func (g *GEtcd) Locker(key string) sync.Locker {
    ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
    defer cancel()
    
    session, err := concurrency.NewSession(g.Client, concurrency.WithTTL(config.TTL), concurrency.WithContext(ctx))
    if err != nil {
        logger.Error("GEtcd::NewLocker - Error: %+v", err)
        return nil
    }
    
    
    concurrency.WithLease(session.Lease())
    
    keyLocker := fmt.Sprintf("/%s/%s/%s", config.ProjectName, KLockerNamespacePrefix, key)
    //lck := concurrency.NewMutex(session, keyLocker)

    return concurrency.NewLocker(session, keyLocker)
}

func (l *GLocker) Lock() *GLocker {
    ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
    l.CancelFunc = cancel

    l.LockerActionError = l.Mutex.Lock(ctx)
    return l
}

func (l *GLocker) Unlock() *GLocker {
    ctx, cancel := context.WithTimeout(context.Background(), KDefaultLockTimeout)
    l.CancelFunc = cancel

    l.LockerActionError = l.Mutex.Unlock(ctx)
    return l
}

func TestLocker(i int) int {
    // InstanceEtcdManger()

    //etcdG := GetEtcdDiscoveryInstance()
    //etcdG.Locker("AA").Lock()
    //defer etcdG.Locker("AA").Lock().Unlock()

    return i
}
