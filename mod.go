package mod2

import "github.com/martende/mod3"

func HelloWorld() int {
	return 20 + mod3.HelloWorld()
}
