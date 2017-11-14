package main

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	ErrGroup()
}

// ErrGroup errgroup sample
func ErrGroup() {
	eg := errgroup.Group{}
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {
			// Do someting
			err := DoSomething(i)
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
	return
}

// DoSomething sample func
func DoSomething(i int) error {
	fmt.Println(i)
	if i == 5 {
		err := errors.New("has error")
		return err
	}
	return nil
}
