package main

import(
	"testing"
	"sync"
)

func TestSync(t *testing.T){
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}
	
		var wg sync.WaitGroup
		wg.Add(wantedCount)
	
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
	
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}