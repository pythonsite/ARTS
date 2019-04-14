package lib

import (
	"os"
	in "ARTS/2019031902/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
