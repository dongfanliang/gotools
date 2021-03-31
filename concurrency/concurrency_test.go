package concurrency

import "testing"

func f(p interface{}) (interface{}, error) {
	return p, nil
}

func TestEx(t *testing.T) {
	c := NewConcurrency(3)
	result, err := c.Do(f, []interface{}{1, 3, 4, 5, 6})
	t.Log(result, err)
}
