package pprof

import (
	"fmt"
	"testing"

	"github.com/wj-stack/fuda/pkg/httpx"
)

func TestPPROF(t *testing.T) {
	fmt.Println(Start(":8000"))
	v, _ := httpx.Get("http://127.0.0.1:8000/debug/pprof")
	fmt.Println(string(v))
	fmt.Println(Stop())
}

func TestAuto(t *testing.T) {
	fmt.Println(Auto())
	fmt.Println(Stop())
}
