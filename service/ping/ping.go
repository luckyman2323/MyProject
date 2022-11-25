package ping

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/go-ping/ping"
)

func PingConn(addr string) (bool, error) {
	Command := fmt.Sprintf("ping -c 1 -W 3 %s > /dev/null && echo true || echo false", addr)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	return string(output) == "true\n", err
}

func PingTest(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		fmt.Println(err)
	}
	pinger.Debug = true
	pinger.OnFinish = func(statistics *ping.Statistics) {
		s := fmt.Sprintf("OnFinish: %#v\n", statistics)
		fmt.Println("======", s, "========")
	}
	pinger.OnRecv = func(packet *ping.Packet) {
		s := fmt.Sprintf("OnRecv: %#v\n", packet)
		fmt.Println("======", s, "========")
	}
	pinger.Timeout = time.Second * 3
	pinger.Count = 3
	err = pinger.Run() // blocks until finished
    if err != nil {
		panic(err)
	}
}

func ServerPing(target string) bool {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		panic(err)
	}

	pinger.Count = 5

	pinger.Timeout = time.Second * 10

	pinger.SetPrivileged(true)

	err = pinger.Run() // blocks until finished
	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()

	fmt.Println("======================", target)
	fmt.Println(stats)

	// 有回包，就是说明IP是可用的

	if stats.PacketsRecv >= 1 {

		return true

	}

	return false
}
