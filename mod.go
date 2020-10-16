package mod2

import "github.com/martende/mod3"

func HelloWorld() int {
	return 30 + mod3.HelloWorld()
}
