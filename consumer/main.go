package main

import "sync"
import "github.com/romance-dev/try-word-gen"

func main() {
	g := wordgen.New(`学中文:n-v-a-n`)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println(g.String())
		}()
	}
	wg.Wait()
}
