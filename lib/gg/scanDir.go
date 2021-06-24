package gg

import (
	"io/ioutil"
)

func ScanDir(filePath string) []string {
	if filePath[len(filePath)-1:] != "/" {
		filePath = filePath + "/"
	}
	files, _ := walkDir(filePath)
	// fmt.Println(10)
	// for _, v := range files {
	// 	fmt.Println(v)
	// }
	// var ret []string
	// ret = append(ret, filePath)
	// ret := make([]string,1)
	return append([]string{filePath}, files...)
}

func walkDir(filepath string) ([]string, error) {
	files, err := ioutil.ReadDir(filepath) // files为当前目录下的所有文件名称【包括文件夹】
	if err != nil {
		return nil, err
	}

	var allfile []string
	for _, v := range files {
		fullPath := filepath + v.Name() + "/" // 全路径 + 文件名称
		if v.IsDir() {                        // 如果是目录
			// fmt.Println(30, filepath+v.Name())
			// fmt.Println(25, fullPath)
			allfile = append(allfile, fullPath)
			a, _ := walkDir(fullPath) // 遍历改路径下的所有文件

			allfile = append(allfile, a...)
		}
		//  else {
		// 	fmt.Println(28, fullPath)
		// 	continue
		// 	allfile = append(allfile, fullPath) // 如果不是文件夹，就直接追加到路径下
		// }
	}
	// fmt.Println(35, allfile)
	return allfile, nil
}
