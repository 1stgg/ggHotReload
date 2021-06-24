package main

import (
	"log"
	"os"
	"strings"

	"ggHotReload/lib/gg"

	"github.com/fsnotify/fsnotify"
)

var ggHotReloadVersion = "1.0.0"
var param = map[string]string{
	"sh":    "sh ./ggHR.sh",
	"watch": "./",
}

var runParent = gg.Debounce(func() {
	gg.Parent(param["sh"])
}, 500)

func main() {
	gg.PrintColor("[ggHR] "+ggHotReloadVersion, "warn")
	// gg.PrintColor("[ggHR] to restart at any time, enter `rs`","warn")
	gg.PrintColor("[ggHR] watching: "+param["watch"], "warn")
	// flag.Parse()
	if len(os.Args) > 0 {
		for _, arg := range os.Args {
			// fmt.Printf("args[%d]=%v\n", index, arg)
			// fmt.Println(index, arg)
			argArr := strings.Split(arg, "=")
			// fmt.Println(54, argArr[0])
			// param[argArr[0]] = argArr[1]
			switch argArr[0] {
			case "sh":
				// param["sh"] = argArr[1]
				param["sh"] = strings.Replace(argArr[1], ".$go", ".go", -1)
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
