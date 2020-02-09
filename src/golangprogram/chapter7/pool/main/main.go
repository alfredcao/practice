package main

import (
	"fmt"
	"golangprogram/chapter7/pool"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	idCounter int64
	waitGroup sync.WaitGroup
)

const poolSize = 10
const goroutineCount = 30

func main() {
	p, err := pool.NewPool(NewDBConn, poolSize)
	if err != nil {
		fmt.Println("new pool failed, err :", err)
		return
	}

	waitGroup.Add(goroutineCount)
	for i := 1; i <= goroutineCount; i++ {
		go func(id int) {
			defer waitGroup.Done()
			dbConn, err := p.Acquire()
			if err != nil {
				fmt.Printf("task[%d] aquire dbconn failed, err : %v\n", id, err)
				return
			}
			defer p.Ret(dbConn)
			fmt.Printf("task[%d] aquire dbconn success\n", id)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Printf("task[%d] finish query\n", id)
		}(i)
	}

	waitGroup.Wait()
	p.Close()
	fmt.Println("pool close")
}

type DBConn struct {
	ID int64
}

func (dbConn *DBConn) Close() error {
	fmt.Println("Close DBConn=", dbConn.ID)
	return nil
}

func NewDBConn() (io.Closer, error) {
	id := atomic.AddInt64(&idCounter, 1)
	return &DBConn{
		ID: id,
	}, nil
}
