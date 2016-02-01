package nil

import (
	"fmt"
	"os/exec"
)

func Func() {
	path := exec.Command("pwd")
	fmt.Println(path)
}
