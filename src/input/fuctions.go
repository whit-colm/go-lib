package input

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	"strconv"
	"strings"
)

func Prompt(p string) (t string, err error) {
	//P in Prompt is capitalised because it's a public function,
	//meaning it can be used in any go file if you know where to look.
	//
	//This is still here because a lot of my crap is dependant on this.
	fmt.Print(c.X + p)
	reader := bufio.NewReader(os.Stdin)
	t, err = reader.ReadString('\n')
	t = strings.TrimSpace(t)
	return t, nil
}

func StringInput(p string) string {
	//This function just returns whatever is typed as a string.
	fmt.Print(c.X + p)
	reader := bufio.NewReader(os.Stdin)
	t, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	t = strings.TrimSpace(t)
	t = strings.ToLower(t)
	return t
}

func IntInput(p string) int {
	//This function just returns whatever is typed as a string.
	fmt.Print(c.X + p)
	reader := bufio.NewReader(os.Stdin)
	t, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	t = strings.TrimSpace(t)
	i, err := strconv.Atoi(t)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return i
}

func DoubleInput(p string) float64 {
	//This function just returns whatever is typed as a string.
	fmt.Print(c.X + p)
	reader := bufio.NewReader(os.Stdin)
	t, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	t = strings.TrimSpace(t)
	f, err := strconv.ParseFloat(t, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return f
}
