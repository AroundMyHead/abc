package abc

import (
	"abc/constant"

	"golang.org/x/example/hello/reverse"
)

func ReverseHello() string {
	return reverse.String(constant.Hello)
}