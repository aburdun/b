package b

import (
        "fmt"
        "github.com/aburdun/a"
)

func PrintMe() {
        fmt.Println("I am b v3 and am using a v4? " + a.GetVersion())
}

