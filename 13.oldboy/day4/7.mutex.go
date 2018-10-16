package main

import (
	"sync"
	"math/rand"
	"time"
	"sync/atomic"
	"fmt"
)

var lock sync.Mutex
var rwLock sync.RWMutex

func testMutexLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				lock.Lock()
				time.Sleep(time.Millisecond)
				//fmt.Println(a)
				lock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}


func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				time.Sleep(time.Millisecond)
				//fmt.Println(a)
				rwLock.RUnlock()

				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

//读写锁是互斥锁性能的100倍
func main() {
	//testMutexLock() //1598
	testRWLock() 	  //183258
}