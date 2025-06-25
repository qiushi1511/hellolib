package hellolib

import (
	"fmt"

	"github.com/qiushi1511/hellolib/greet"
)

func Hello(name string) {
	fmt.Println(greet.Greet(name))
}

func WarmHello(name string) {
	fmt.Println(greet.WarmGreet(name))
}
