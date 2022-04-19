// 整个 go 程序中只允许存在一个 main 包
package main

import "fmt"

/* main 函数是 go 程序的入口 */
func main() {
	// Println 大写字母开头，表示该标识符可以在声明的包外可见，如果是小写开头则仅在声明的包内可见
	fmt.Println("Hello, world!")
	fmt.Println("你好，世界！")
}
