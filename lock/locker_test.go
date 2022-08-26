package lock

import (
    "fmt"
    "github.com/spf13/viper"
    "github.com/stretchr/testify/assert"
    "strings"
    "sync"
    "testing"
    "time"
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

func TestSingleLocker(t *testing.T) {
    InstanceZookeeper()
    mu := GetInstanceZookeeper().NewLocker("tien")
    err := mu.Lock()
    if err != nil {
        t.Logf("TestSingleLocker - lock Error: %v", err)
    }
    assert.Nil(t, err, err)

    time.Sleep(30 * time.Second)
    assert.Nil(t, err, err)
    if err != nil {
        t.Logf("TestSingleLocker - unlock Error: %v", err)
    }
}

func BenchmarkConnectZookeeper(b *testing.B) {
    fmt.Println("a")
}

func TestMultiLocker(t *testing.T) {
    InstanceZookeeper()
    //mu := GetInstanceZookeeper().NewLocker("tien")

    wg := sync.WaitGroup{}
    wg.Add(2)

    idx := []string{"init locker"}

    go func() {
        mu := GetInstanceZookeeper().NewLocker("tien")
        idx = append(idx, fmt.Sprintf("1 prepare lock: %s", time.Now()))
        defer wg.Done()

        err := mu.Lock()
        if err != nil {
            t.Logf("TestMultiLocker - lock Error: %v", err)
        }
        assert.Nil(t, err, err)
        idx = append(idx, fmt.Sprintf("1 lock err: %+v: %s", err, time.Now()))
        idx = append(idx, fmt.Sprintf("1 lock done: %s", time.Now()))

        //time.Sleep(30 * time.Second)

        idx = append(idx, fmt.Sprintf("1 unlock done: %s", time.Now()))

        err = mu.Unlock()
        idx = append(idx, fmt.Sprintf("1 unlock err: %+v: %s", err, time.Now()))
        assert.Nil(t, err, err)
    }()

    go func() {
        mu := GetInstanceZookeeper().NewLocker("tien")
        idx = append(idx, fmt.Sprintf("2 prepare lock: %s", time.Now()))
        defer wg.Done()

        //time.Sleep(3 * time.Second)
        err := mu.Lock()
        if err != nil {
            t.Logf("TestMultiLocker - lock Error: %v", err)
        }
        assert.Nil(t, err, err)
        idx = append(idx, fmt.Sprintf("2 lock err: %+v: %s", err, time.Now()))
        idx = append(idx, fmt.Sprintf("2 lock done: %s", time.Now()))

        idx = append(idx, fmt.Sprintf("2 unlock done: %s", time.Now()))

        err = mu.Unlock()
        assert.Nil(t, err, err)
        idx = append(idx, fmt.Sprintf("2 unlock err: %+v: %s", err, time.Now()))
    }()
    
    wg.Wait()

    t.Logf("Idx: %+v", strings.Join(idx, "\n"))
}
