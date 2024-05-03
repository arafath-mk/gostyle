package gcstyle

import (
	"fmt"

	"github.com/arafath-mk/gostyle/wcolor"
)

type Style struct {
	Color         *wcolor.Color
	Background    *wcolor.Color
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	Darken        bool
	Lighten       bool
}

func (s Style) ApplyTo(str string, applyStyle bool) StyledString {
	styled := fmt.Sprintf("%s%s%s", s.Start(applyStyle), str, s.End(applyStyle))
	return StyledString{
		str:        styled,
		applyStyle: applyStyle,
	}
}

func (s Style) Start(applyStyle bool) string {
	if !applyStyle {
		return ""
	}

	var str string
	if s.Color != nil {
		str = fmt.Sprintf("%s\033[38;2;%d;%d;%dm", str, s.Color.R, s.Color.G, s.Color.B)
	}
	if s.Background != nil {
		str = fmt.Sprintf("%s\033[48;2;%d;%d;%dm", str, s.Color.R, s.Color.G, s.Color.B)
	}
	if s.Bold {
		str = fmt.Sprintf("%s\033[1m", str)
	}
	if s.Italic {
		str = fmt.Sprintf("%s\033[3m", str)
	}
	if s.Underline {
		str = fmt.Sprintf("%s\033[4m", str)
	}
	if s.Strikethrough {
		str = fmt.Sprintf("%s\033[9m", str)
	}
	if s.Darken {
		str = fmt.Sprintf("%s\033[2m", str)
	}
	if s.Lighten {
		str = fmt.Sprintf("%s\033[1m", str)
	}
	return str
}

func (s Style) End(applyStyle bool) string {
	if !applyStyle {
		return ""
	}

	var str string
	if s.Lighten {
		str = fmt.Sprintf("%s\033[0m", str)
	}
	if s.Darken {
		str = fmt.Sprintf("%s\033[0m", str)
	}
	if s.Strikethrough {
		str = fmt.Sprintf("%s\033[0m", str)
	}
	if s.Underline {
		str = fmt.Sprintf("%s\033[0m", str)
	}
	if s.Italic {
		str = fmt.Sprintf("%s\033[0m", str)
	}
	if s.Bold {
		str = fmt.Sprintf("%s\033[22m", str)
	}
	if s.Background != nil {
		str = fmt.Sprintf("%s\033[m", str)
	}
	if s.Color != nil {
		str = fmt.Sprintf("%s\033[m", str)
	}
	return str
}
