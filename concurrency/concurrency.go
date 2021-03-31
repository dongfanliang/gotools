package concurrency

import (
	"sync"

	nsema "github.com/toolkits/concurrent/semaphore"
)

type Concurrency struct {
	wg      sync.WaitGroup
	errOnce sync.Once
	err     error

	q    chan interface{} // 结果队列
	sema *nsema.Semaphore
}

func NewConcurrency(n int) *Concurrency {
	if n < 1 {
		n = 1
	}
	return &Concurrency{q: make(chan interface{}, 1), sema: nsema.NewSemaphore(n)}
}

func (c *Concurrency) Do(f func(interface{}) (interface{}, error), paramSlice []interface{}) ([]interface{}, error) {
	result := []interface{}{}
	go func() {
		for item := range c.q {
			if item != nil {
				result = append(result, item)
			}
			c.wg.Done()
		}
	}()

	for i := 0; i < len(paramSlice); i++ {
		c.sema.Acquire()
		c.wg.Add(1)
		go func(item interface{}) {
			value, err := f(item)
			if err != nil {
				c.errOnce.Do(func() {
					c.err = err
				})
			} else {
				c.q <- value
			}

			c.sema.Release()
		}(paramSlice[i])
	}

	c.wg.Wait()
	close(c.q)
	return result, c.err
}
