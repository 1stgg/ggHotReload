package main

import "fmt"

func main() {
	// string、number、boolean、null、undefined；
	Log("abc", 1, false, nil, 1)
	// Object  Array
	Log(map[string]string{"a": "b", "b": "c"}, map[int]string{1: "b", 2: "b"}, map[string]int{"b": 1, "c": 1})
	Log([]int{1, 2}, []string{"a", "b"})

	// Function
}

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

			fmt.Println(i)
		}
	}
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
