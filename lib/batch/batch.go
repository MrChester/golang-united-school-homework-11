package main

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mx sync.Mutex
	var i int64
	users := make(chan struct{}, pool)

	for i = 0; i < n; i++ {
		wg.Add(1)
		users <- struct{}{}
		go func(userId int64) {
			user := getOne(userId)

			mx.Lock()
			res = append(res, user)
			mx.Unlock()

			_, ok := <-users
			if !ok {
				return
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	close(users)
	return
}
