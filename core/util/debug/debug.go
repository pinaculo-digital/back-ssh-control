package text

import (
	"fmt"

	"github.com/fatih/color"
)

func Println(args ...any) {
	boldGreen := color.New(color.FgGreen, color.Bold)

	fmt.Println()
	boldGreen.Println(args...)
	fmt.Println()

}
func Errorln(args ...any) {
	boldRed := color.New(color.FgRed, color.Bold)

	boldRed.Println(args...)
}
