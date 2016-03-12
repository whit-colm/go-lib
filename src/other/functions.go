package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	m "math/rand"
	"os"
)

func Exit(message string, clear bool) {
	if clear == true {
		fmt.Println(c.CL + c.B1 + message)
	} else {
		fmt.Println(c.B1 + message)
	}
	os.Exit(-1)
}

func Spacer(times int) {
	repeat := 0
	for repeat < times {
		fmt.Println()
		repeat++
	}
}

func Go(testfor int) {
	fmt.Println()
	if testfor == 0 {
		fmt.Println(c.CL)
	} else if testfor == 1 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
	fmt.Println(c.B1 + "Press enter to continue")
	i.Prompt(c.B + ">" + c.M)
	fmt.Println(c.CL)
}

var errorResponse = []string{"Well...shebang.",
	"OH GOD EVERYTHING IS ON FIRE!",
	"10/10, 10/10, 100/100, BEST SCRIPT! BEST SCRIPT",
	"Hey look, death!",
	"Sure, why the hell not...",
	"Thank's CURRENT PRESEDENT",
	"Nope factor 5 Mr. Sulu",
	"Ooohh! a shiney!",
	"CODE BY L33T HAXX3R.",
	"So... How's your day going?",
}

func QuitAtError(problem error) {
	if problem != nil {
		fmt.Println(c.CL)
		fmt.Println(errorResponse[m.Intn(9)])
		fmt.Println()
		fmt.Println(c.V + "Details of the error are below:" + c.B00)
		fmt.Println(c.B01+problem+c.B00)
		os.Exit(1)
	} else {
		fmt.Println(c.V + "ERROR CHECK: " + c.B00 + "no errors found")
	}

}

func RandomInt(low int, high int) int {
	if low >= high || low < 0 || high < 0 {
		fmt.Println(c.B01 + "Sorry, please declare two integers both greater than 0 where the second number is larger than the first number." + c.X)
		Exit("Program stopped", false)
	}
	n := high - low
	unprocessedRandInt := m.Intn(n)
	randInt := unprocessedRandInt + low
	return randInt
}
