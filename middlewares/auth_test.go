package middlewares

import (
	"sync"
	"testing"
)

func TestAddSessionConcurrent(t *testing.T) {
	auth := NewAuthMiddleware()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			auth.AddSession("session", "1111111111")
			if _, ok := auth.SessionPhone("session"); !ok {
				t.Error("expected session phone to be stored")
			}
		}(i)
	}
	wg.Wait()

	phone, ok := auth.SessionPhone("session")
	if !ok || phone != "1111111111" {
		t.Fatalf("SessionPhone() = %q %v", phone, ok)
	}
}
