package gg

import (
	"fmt"
	"reflect"
	"runtime"
)

func Log(itf ...interface{}) {
	logs(itf...)
	fmt.Printf("\n")
}
func logs(itf ...interface{}) {
	for index, i := range itf {
		if index != 0 {
			fmt.Printf(" ")
		}
		switch i.(type) {
		case string:
			logStr(i.(string))
		case int:
			logInt(i.(int))
		case bool:
			logBool(i.(bool))
		case map[string]string:
			logMapStrStr(i.(map[string]string))
		case map[string]int:
			logMapStrInt(i.(map[string]int))
		case map[int]string:
			logMapIntStr(i.(map[int]string))
		case []int:
			logArrInt(i.([]int))
		case []string:
			logArrStr(i.([]string))
		// case []string:
		//     data := i.([]string)
		//     length := len(data)
		//     if length == 2 {
		//         return &person{
		//             Name: data[0],
		//             Sex:  data[1],
		//         }
		//     } else {
		//         return nil
		//     }
		default:

			if i != nil {
				typeOfA := reflect.TypeOf(i)
				// fmt.Println(typeOfA.Name(), typeOfA.Kind())
				// fmt.Printf("%d", typeOfA.Kind())
				// fmt.Print(typeOfA.Kind() == 19)
				// var funcType reflect.Kind
				// funcType :=
				switch typeOfA.Kind() {
				case 19: // func
					logFunc(i)
				default:
					fmt.Println(i)
				}
			} else {
				fmt.Println(i)
			}

		}
	}
}
func logFunc(arg interface{}) {
	funcName := runtime.FuncForPC(reflect.ValueOf(arg).Pointer()).Name()
	ret := "[Function: " + funcName + "]"
	fmt.Printf("%c[36m%s%c[0m", 0x1B, ret, 0x1B)
}
func logBool(arg bool) {
	fmt.Printf("%c[33m%t%c[0m", 0x1B, arg, 0x1B)
}
func logInt(arg int) {
	fmt.Printf("%c[33m%d%c[0m", 0x1B, arg, 0x1B)
}
func logStr(arg string) {
	fmt.Printf("%c[32m%s%c[0m", 0x1B, `"`+arg+`"`, 0x1B)
}
func logArrInt(arg []int) {
	fmt.Printf("%s", `[ `)
	for index, item := range arg {
		// logs(key)
		if index != 0 {
			fmt.Printf("%s", `, `)
		}

		logs(item)
	}
	fmt.Printf("%s", ` ]`)
}
func logArrStr(arg []string) {
	fmt.Printf("%s", `[ `)
	for index, item := range arg {
		// logs(key)
		if index != 0 {
			fmt.Printf("%s", `, `)
		}

		logs(item)
	}
	fmt.Printf("%s", ` ]`)
}
func logMapStrStr(arg map[string]string) {
	fmt.Printf("%s", `{ `)
	isFirst := true
	for key, item := range arg {

		if !isFirst {

			fmt.Printf("%s", `, `)
		}
		isFirst = false
		logs(key)
		fmt.Printf("%s", `: `)
		logs(item)
	}
	fmt.Printf("%s", ` }`)
}
func logMapStrInt(arg map[string]int) {
	fmt.Printf("%s", `{ `)
	isFirst := true
	for key, item := range arg {

		if !isFirst {

			fmt.Printf("%s", `, `)
		}
		isFirst = false
		logs(key)
		fmt.Printf("%s", `: `)
		logs(item)
	}
	fmt.Printf("%s", ` }`)
}
func logMapIntStr(arg map[int]string) {
	fmt.Printf("%s", `{ `)
	isFirst := true
	for key, item := range arg {

		if !isFirst {

			fmt.Printf("%s", `, `)
		}
		isFirst = false
		logs(key)
		fmt.Printf("%s", `: `)
		logs(item)
	}
	fmt.Printf("%s", ` }`)
}
