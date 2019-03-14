/**
 * Created with Goland.
 * Description: 
 * User: Insomnia
 * Date: 2019-03-14
 * Time: 下午9:07
 */
/*
	Package main the service entry
*/
package main

import (
	"GinBlog/service/pkg/setting"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	s := setting.LoadServer()
	if s != nil {
		fmt.Println("wtf")
	}
}
