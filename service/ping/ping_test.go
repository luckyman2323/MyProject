package ping

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	res := ServerPing("192.168.9.171")
	fmt.Println("====", res)
	// pingTest("192.168.9.171")

}
