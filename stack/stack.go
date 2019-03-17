package stack

import (
	"fmt"
	"strings"
)

// IsValid leetcode 20
// stack,but golang don't have stack implements, use array
func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, char := range s {
		//fmt.Println(reflect.TypeOf(char))
		if char == '(' || char == '{' || char == '[' {
			// 入栈
			stack = append(stack, char)
			// 循环中，stack不能为空
		} else if len(stack) > 0 && brackets[char] == stack[len(stack) - 1] {
			// 栈中有数据，且此元素与栈尾元素相同
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据则 false
	return len(stack) == 0
}

func splite(s rune) bool {
	if s == '/' {
		return true
	}
	return false
}

// SimplifyPath leetcode 71

func simplifyPath(path string) string {
	dirs := strings.FieldsFunc(path, splite)
	for i := 0; i < len(dirs); i++ {
		if dirs[i] == "." {
			dirs = append(dirs[:i], dirs[i+1:]...)
			i--
		} else {
			if dirs[i] == ".." {
				if i == 0 {
					dirs = dirs[1:]
					i-- // 去掉当前的"."
					continue
				}
				dirs = append(dirs[:i-1], dirs[i+1:]...)
				i -= 2 // 去掉前一个值和当前的".."
			}
		}
	}
	res := "/"
	for i := 0; i < len(dirs); i++ {
		res += dirs[i]
		if i != len(dirs)-1 {
			res += "/"
		}
	}
	return res
}



func main() {
	s := "{{{[()]}}}"
	fmt.Println(isValid(s))
	res := simplifyPath("/a//b////c/d//././/..")
	fmt.Println(res)
}
