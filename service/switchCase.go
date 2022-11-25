package service

import (
	"fmt"
	"time"
)

func SwitchCase() {
	s := "a"
	for {
		switch s {
		case "a":
			{
				fmt.Println("s1")
				s = "b"
				continue
				// time.Sleep(1*time.Second)
				// continue
			}
		case "b":
			{
				fmt.Println("s2")
				s = "c"
				time.Sleep(1 * time.Second)
				continue
			}
		case "c":
			{
				fmt.Println("s3")
				time.Sleep(1 * time.Second)
				break
			}
		}
		fmt.Println("=================")
	}
}
