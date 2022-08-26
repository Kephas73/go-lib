package lock_etcd

import (
    "fmt"
    "github.com/spf13/viper"
    "testing"
)

func init() {
    viper.SetConfigFile(`config.json`)

    err := viper.ReadInConfig()

    if err != nil {
        panic(err)
    }

    if viper.GetBool(`Debug`) {
        fmt.Println("Service RUN on DEBUG mode")
    } else {
        fmt.Println("Service RUN on PRODUCTION mode")
    }
}

func BenchmarkLocker(b *testing.B) {
    InstanceEtcdManger()
    for i := 0; i < b.N; i++ {
        TestLocker(i)
    }
}




//func TestSingleLocker(t *testing.T) {
//    // InstanceEtcdManger()
//
//    locker := GetEtcdDiscoveryInstance().NewLocker("key1")
//    err := locker.Lock()
//    if err != nil {
//        log.Printf("TestSingleLocker - Error: %v", err)
//    }
//    assert.Nil(t, err, err)
//    err = locker.Unlock()
//    if err != nil {
//        log.Printf("TestSingleLocker - Error: %v", err)
//    }
//    assert.Nil(t, err, err)
//}
//
//
//
//
//
//func TestMultiLocker(t *testing.T) {
//    InstanceEtcdManger()
//
//    locker := GetEtcdDiscoveryInstance().NewLocker("key1")
//    idx := []string{"init locker"}
//
//    wg := sync.WaitGroup{}
//    wg.Add(1)
//
//    /*    for i := 0; i < 2; i++ {
//          wg.Add(1)
//          go func(i int) {
//              idx = append(idx, fmt.Sprint(i), "1 prepare lock")
//              defer wg.Done()
//
//              locker.Lock()
//              idx = append(idx, fmt.Sprint(i),"1 lock done")
//              //time.Sleep(5 * time.Second)
//              idx = append(idx, fmt.Sprint(i),"1 unlock done")
//              locker.Unlock()
//          }(i)
//      }*/
//
//    go func(locker *GLocker) {
//        idx = append(idx, "1 prepare lock")
//        defer wg.Done()
//
//        timeNow := time.Now()
//        locker.Lock()
//        idx = append(idx, "1 lock done"+fmt.Sprint(time.Since(timeNow).Seconds()))
//        time.Sleep(5 * time.Second)
//        idx = append(idx, "1 unlock done")
//        locker.Unlock()
//    }(locker)
//
//    wg.Add(1)
//    go func(locker *GLocker) {
//        time.Sleep(1 * time.Second)
//        idx = append(idx, "2 prepare lock")
//        defer wg.Done()
//
//        timeNow := time.Now()
//        locker.Lock()
//        idx = append(idx, "2 lock done "+fmt.Sprint(time.Since(timeNow).Seconds()))
//        idx = append(idx, "2 unlock done")
//        locker.Unlock()
//    }(locker)
//
//    wg.Wait()
//    t.Logf("Idx: %+v", strings.Join(idx, "\n"))
//
//}
