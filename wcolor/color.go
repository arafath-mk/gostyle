package wcolor

import "fmt"

type Color struct {
	R, G, B byte
}

func (c Color) Clone() *Color {
	newColor := c
	return &newColor
}

func (c Color) Start(colorize bool) string {
	if !colorize {
		return ""
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c Color) End(colorize bool) string {
	if !colorize {
		return ""
	}
	return "\033[m"
}

func (c Color) StartBG(colorize bool) string {
	if !colorize {
		return ""
	}
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c Color) EndBG(colorize bool) string {
	if !colorize {
		return ""
	}
	return "\033[m"
}

func RGB(r, g, b byte) Color {
	return Color{r, g, b}
}

func HEX(css3Color string) Color {
	if css3Color[0] != '#' {
		return Color{} // Error: Invalid Format
	}

	var c Color
	switch len(css3Color) {
	case 7:
		c.R = hexCharToDecimalByte(css3Color[1])<<4 + hexCharToDecimalByte(css3Color[2])
		c.G = hexCharToDecimalByte(css3Color[3])<<4 + hexCharToDecimalByte(css3Color[4])
		c.B = hexCharToDecimalByte(css3Color[5])<<4 + hexCharToDecimalByte(css3Color[6])
	case 4:
		c.R = hexCharToDecimalByte(css3Color[1]) * 17
		c.G = hexCharToDecimalByte(css3Color[2]) * 17
		c.B = hexCharToDecimalByte(css3Color[3]) * 17
	default:
		return Color{} // Error: Invalid Format
	}
	return c
}

func hexCharToDecimalByte(hexChar byte) byte {
	switch {
	case hexChar >= '0' && hexChar <= '9':
		return hexChar - '0'
	case hexChar >= 'a' && hexChar <= 'f':
		return hexChar - 'a' + 10
	case hexChar >= 'A' && hexChar <= 'F':
		return hexChar - 'A' + 10
	}

	return 0 // Error: Invalid Format
}
