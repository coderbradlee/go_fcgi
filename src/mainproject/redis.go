package main
import (
    "log"
    "fmt"
    "time"
    "strconv"
    "go_redis_cluster"
    "net/http"
)

const kNumOfRoutine = 2

func redisHandler(w http.ResponseWriter, r *http.Request) {
    cluster, err := redis.NewCluster(
    &redis.Options{
        StartNodes: configuration.RedisNodes,
        ConnTimeout: 50 * time.Millisecond,
        ReadTimeout: 50 * time.Millisecond,
        WriteTimeout: 50 * time.Millisecond,
        KeepAlive: 16,
        AliveTime: 60 * time.Second,
    })

    if err != nil {
    log.Fatalf("redis.New error: %s", err.Error())
    }
    start := time.Now()
    chann := make(chan int, kNumOfRoutine)
    for i := 0; i < kNumOfRoutine; i++ {
        go redisTest(cluster, i * 5, (i+1)*5, chann)
    }

    for i := 0; i < kNumOfRoutine; i++ {
        _ = <-chann
    }
    end := time.Now()
 
    //输出执行时间，单位为毫秒。
    // fmt.Println((end-start)/500000)
    fmt.Fprint(w, 10/(end.Sub(start)))
}

func redisTest(cluster *redis.Cluster, begin, end int, done chan int) {
    prefix := "{mykey}"
    for i := begin; i < end; i++ {
        key := prefix + strconv.Itoa(i)

        _, err := cluster.Do("set", key, i*10)
        if err != nil {
            //fmt.Printf("-set %s: %s\n", key, err.Error())
            //time.Sleep(100 * time.Millisecond)
            continue
        }
    
    }

    done <- 1
}


