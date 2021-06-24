package gg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	si            io.WriteCloser
	so            io.ReadCloser
	err           error
	c             *exec.Cmd
	islistenWrite = false
	islistenRead  = false
)

func ParentClose() {

	if c != nil {
		PrintColor("restarting due to changes...", "success")
		// fmt.Println(19, "kill")
		var mutex sync.Mutex
		mutex.Lock()
		c.Process.Kill()
		mutex.Unlock()

	}

}
func Parent(cmdStr string) {
	ParentClose()
	cmdArr := strings.Split(cmdStr, " ")
	// fmt.Println(30, cmdArr, cmdStr)
	// fmt.Println(21, strings.Split("a 2", " "))
	cmdAfterArr := cmdArr[1:]

	PrintColor("[ggHR] starting `"+cmdStr+"`", "success")
	c = exec.Command(cmdArr[0], cmdAfterArr...)
	// c := exec.Command("echo", "666")
	si, err = c.StdinPipe()
	if err != nil {
		// fmt.Println(err)
		return
	}

	so, err = c.StdoutPipe()
	if err != nil {
		// fmt.Println(err)
		return
	}
	reader := bufio.NewReader(so)

	err = c.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if !islistenWrite {
		islistenWrite = true
		listenWrite()
	}

	if !islistenRead {
		islistenRead = true
		listenRead(reader)
	}

	// Close the input and wait for exit

	c.Wait()
	time.Sleep(Ms(1000))
	PrintColor("[ggHR] clean exit - waiting for changes before restart", "success")
}

func listenWrite() {
	go func() {
		for {
			var buffer [512]byte

			_, err := os.Stdin.Read(buffer[:])
			if err != nil {

				// fmt.Println("read error: ", err)
				return

			}

			// fmt.Println("count:", n, ", msg:", string(buffer[:]))
			_, err = si.Write(buffer[:])
			if err != nil {
				// fmt.Println(err)
				islistenWrite = false
				break
			}
		}
	}()

}

func listenRead(reader *bufio.Reader) {
	go func() {
		for {

			answer, err := reader.ReadString('\n')
			if err != nil {

				// fmt.Println(113, err)
				islistenRead = false
				break
			}
			fmt.Printf(answer)
		}
	}()

}
