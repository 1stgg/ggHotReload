package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"ggHotReload/lib/gg"

	"github.com/fsnotify/fsnotify"
)

var param = map[string]string{
	"sh":    "sh ./ggHR.sh",
	"watch": "./",
}

// func init() {
// 	flag.StringVar(&sh, "sh", "sh ./ggHR.sh", "reload shell code")
// 	flag.StringVar(&watch, "watch", "./", "watch file change path")
// }

func Debounce(fn func(int), ms float64) func(int) {
	prev := time.Unix(0, 0)

	return func(arg int) {
		curr := time.Now()
		delta := curr.Sub(prev).Seconds()
		if delta < ms {
			return
		}
		prev = curr // 可执行了之后，在刷新计时
		fn(arg)
	}
}

// func runParent() {
// 	fmt.Println(39)
// 	gg.ParentClose()
// 	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "testPrintColor233", 0x1B)
// 	fmt.Println(time.Now().Format("2019/10/29 16:47:38"))
// 	gg.Parent(param["sh"])

// }
var runParent = gg.Debounce(func() {
	fmt.Println(39)
	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "testPrintColor233", 0x1B)
	fmt.Println(time.Now().Format("06/01/02 15:04:05"))
	gg.Parent(param["sh"])
}, 500)

func main() {
	// flag.Parse()
	if len(os.Args) > 0 {
		for _, arg := range os.Args {
			// fmt.Printf("args[%d]=%v\n", index, arg)
			// fmt.Println(index, arg)
			argArr := strings.Split(arg, "=")
			fmt.Println(54, argArr[0])
			// param[argArr[0]] = argArr[1]
			switch argArr[0] {
			case "sh":
				param["sh"] = argArr[1]
			case "watch":
				param["watch"] = argArr[1]
			}
		}
	}

	// fmt.Println(19, param)
	// return
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	defer gg.ParentClose()

	done := make(chan bool)

	go func() {
		runParent()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// fmt.Println(event)
				// log.Println("event: 66", event)

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					runParent()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// 监控目录
	dirs := gg.ScanDir(param["watch"])

	for _, item := range dirs {
		err = watcher.Add(item)

		if err != nil {
			log.Fatal(err)
		}
	}
	// watcher.Add(param["watch"])
	<-done

}
