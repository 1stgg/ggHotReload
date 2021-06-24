package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// fmt.Println(time)
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for {
			var buffer [512]byte

			_, err := os.Stdin.Read(buffer[:])
			if err != nil {

				fmt.Println("read error:", err)
				return

			}

			// fmt.Println("child msg:", string(buffer[:]))
			fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, string(buffer[:]), 0x1B)
		}
	}()
	go func() {
		for _ = range ticker.C {
			fmt.Printf("%v\n", time.Now())
		}
	}()
	time.Sleep(10 * time.Minute)
}
