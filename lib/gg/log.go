package gg

import "fmt"

func main() {
	Log(1, "abc")
	// Log("abc")
}
func Log(itf ...interface{}) {
	for index, i := range itf {
		if index != 0 {
			fmt.Printf(" ")
		}
		switch i.(type) {
		case string:
			logStr(i.(string))
		case int:
			logInt(i.(int))
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
func logInt(arg int) {
	fmt.Printf("%c[33m%d%c[0m", 0x1B, arg, 0x1B)
}
func logStr(arg string) {
	fmt.Printf("%c[32m%s%c[0m", 0x1B, arg, 0x1B)
}
