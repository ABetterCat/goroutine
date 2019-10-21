package main

import (
	"fmt"
	"sync"
)
// 使用waitgroup控制主线程
func main()  {
	myList := []int{1,2,3}
	wg := sync.WaitGroup{}
	for _, v := range myList{
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(v)
		}(v)
	}
	wg.Wait()
}

