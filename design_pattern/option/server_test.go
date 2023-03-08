package option

import (
	"testing"
	"time"
)

func TestOption(t *testing.T) {
	s1, _ := NewServer("localhost", 8080,
		WithTimeout(60*time.Second),
		WithMaxConns(10000),
		WithProtocol("http"),
	)

	if s1.MaxConns != 10000 {
		t.Error("Option set maxConns failed")
	}
}
