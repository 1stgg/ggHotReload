package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmdStr := "go run ./child.go"
	cmdArr := strings.Split(cmdStr, " ")
	cmdAfterArr := cmdArr[1:]
	c := exec.Command(cmdArr[0], cmdAfterArr...)
	si, err := c.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	so, err := c.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(so)

	err = c.Start()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			var buffer [512]byte

			_, err := os.Stdin.Read(buffer[:])
			if err != nil {

				fmt.Println("read error:", err)
				return

			}

			// fmt.Println("count:", n, ", msg:", string(buffer[:]))
			_, err = si.Write(buffer[:])
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	go func() {
		for {
			// sum := fmt.Sprintf("2+%d\n")
			// _, err = si.Write([]byte(sum))
			// if err != nil {
			// 	log.Fatal(err)
			// }
			answer, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(answer)
		}
	}()
	// Now do some maths
	// for {
	// }
	// Close the input and wait for exit
	// si.Close()
	// so.Close()
	c.Wait()
}
