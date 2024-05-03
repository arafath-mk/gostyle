package gcstyle

import (
	"fmt"

	"github.com/arafath-mk/gostyle/wcolor"
)

func foreground(str string, c wcolor.Color) string {
	return fmt.Sprintf("\033[38;2;%sm%s\033[m", colorString(c), str)
}

func background(str string, c wcolor.Color) string {
	return fmt.Sprintf("\033[48;2;%sm%s\033[m", colorString(c), str)
}

func bold(str string) string {
	return fmt.Sprintf("\033[1m%s\033[22m", str)
}

func darken(str string) string {
	return fmt.Sprintf("\033[2m%s\033[0m", str)
}

func lighten(str string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", str)
}

func italic(str string) string {
	return fmt.Sprintf("\033[3m%s\033[0m", str)
}

func underline(str string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", str)
}

func strikethrough(str string) string {
	return fmt.Sprintf("\033[9m%s\033[0m", str)
}

func colorString(c wcolor.Color) string {
	return fmt.Sprintf("%d;%d;%d", c.R, c.G, c.B)
}
