package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

//互斥锁
var lock sync.Mutex

var rwlock sync.RWMutex

func main() {
	//testMap()

	//testMap2()
	testReadWrite()
}

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	var keys []int
	for k, _ := range a {
		keys = append(keys, k)

	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func testMap2() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			a[1] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}

func testReadWrite() {
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
			rwlock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(time.Millisecond * 10)
			rwlock.Unlock()
		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwlock.Lock()
				time.Sleep(time.Millisecond)
				//fmt.Println(a)
				rwlock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}
