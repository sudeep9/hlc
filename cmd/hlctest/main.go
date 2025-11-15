package main

import (
	"fmt"
	"sync"

	"github.com/sudeep9/hlc"
)

func gentest() {
	h := hlc.NewHLC()

	wg := sync.WaitGroup{}
	for range 5 {
		wg.Go(func() {
			for range 10000 {
				t := h.Next()
				fmt.Printf("t=%d\n", t)
			}
		})
	}

	wg.Wait()
}

func main() {
	gentest()
}
