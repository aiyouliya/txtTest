/*********************
defer

defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用
*********************/
package main
import "fmt"

func main(){
	defer fmt.Println("world")

	fmt.Println("hello")

	defer fmt.Println("你好")

	fmt.Println("世界！");
	fmt.Println("这里显示到哪块呢？")

}