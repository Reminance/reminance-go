/**
 * @File : demo04_error
 * @Author: Octal_H
 * @Date : 2019/11/8
 * @Desc :
 */

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err) //syntax error in pattern
		return
	}
	fmt.Println("files:", files)
}
