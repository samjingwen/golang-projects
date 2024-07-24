package qns6_7

import (
	"math/rand"
	"sync"
	"time"
)

type population struct {
	boys, girls int
}

func Apocalypse() float64 {
	var wg sync.WaitGroup

	ch := make(chan population)

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			boys := 0
			girls := 0
			for girls == 0 {
				r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
				if r1.Intn(2) == 1 {
					boys++
				} else {
					girls++
				}
			}
			ch <- population{boys, girls}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	boys := 0
	girls := 0
	for p := range ch {
		boys += p.boys
		girls += p.girls
	}

	return float64(boys) / float64(girls)
}
