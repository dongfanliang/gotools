package main

import "fmt"

func WithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fn()
}
