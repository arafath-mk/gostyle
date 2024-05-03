package gcstyle

import (
	"github.com/arafath-mk/gostyle/wcolor"
)

func AppendColorStart(dst []byte, c wcolor.Color) []byte {
	dst = append(dst, "\033[38;2;"...)
	dst = append(dst, c.R, ';', c.G, ';', c.B)
	dst = append(dst, 'm')
	return dst
}

func AppendColorEnd(dst []byte) []byte {
	dst = append(dst, "\033[m"...)
	return dst
}

func AppendBackgroundStart(dst []byte, c wcolor.Color) []byte {
	dst = append(dst, "\033[48;2;"...)
	dst = append(dst, c.R, ';', c.G, ';', c.B)
	dst = append(dst, 'm')
	return dst
}

func AppendBackgroundEnd(dst []byte) []byte {
	dst = append(dst, "\033[m"...)
	return dst
}

func AppendBoldStart(dst []byte) []byte {
	dst = append(dst, "\033[1m"...)
	return dst
}

func AppendBoldEnd(dst []byte) []byte {
	dst = append(dst, "\033[22m"...)
	return dst
}

func AppendDarkenStart(dst []byte) []byte {
	dst = append(dst, "\033[2m"...)
	return dst
}

func AppendDarkenEnd(dst []byte) []byte {
	dst = append(dst, "\033[0m"...)
	return dst
}

func AppendLightenStart(dst []byte) []byte {
	dst = append(dst, "\033[1m"...)
	return dst
}

func AppendLightenEnd(dst []byte) []byte {
	dst = append(dst, "\033[0m"...)
	return dst
}

func AppendItalicStart(dst []byte) []byte {
	dst = append(dst, "\033[3m"...)
	return dst
}

func AppendItalicEnd(dst []byte) []byte {
	dst = append(dst, "\033[0m"...)
	return dst
}

func AppendUnderlineStart(dst []byte) []byte {
	dst = append(dst, "\033[4m"...)
	return dst
}

func AppendUnderlineEnd(dst []byte) []byte {
	dst = append(dst, "\033[0m"...)
	return dst
}

func AppendStrikethroughStart(dst []byte) []byte {
	dst = append(dst, "\033[9m"...)
	return dst
}

func AppendStrikethroughEnd(dst []byte) []byte {
	dst = append(dst, "\033[0m"...)
	return dst
}

func AppendStyleStart(dst []byte, s Style) []byte {
	if s.Color != nil {
		dst = append(dst, AppendColorStart(dst, *s.Color)...)
	}
	if s.Background != nil {
		dst = append(dst, AppendBackgroundStart(dst, *s.Background)...)
	}
	if s.Bold {
		dst = append(dst, AppendBoldStart(dst)...)
	}
	if s.Italic {
		dst = append(dst, AppendItalicStart(dst)...)
	}
	if s.Underline {
		dst = append(dst, AppendUnderlineStart(dst)...)
	}
	if s.Strikethrough {
		dst = append(dst, AppendStrikethroughStart(dst)...)
	}
	if s.Darken {
		dst = append(dst, AppendDarkenStart(dst)...)
	}
	if s.Lighten {
		dst = append(dst, AppendLightenStart(dst)...)
	}
	return dst
}

func AppendStyleEnd(dst []byte, s Style) []byte {
	if s.Lighten {
		dst = append(dst, AppendLightenEnd(dst)...)
	}
	if s.Darken {
		dst = append(dst, AppendDarkenEnd(dst)...)
	}
	if s.Strikethrough {
		dst = append(dst, AppendStrikethroughEnd(dst)...)
	}
	if s.Underline {
		dst = append(dst, AppendUnderlineEnd(dst)...)
	}
	if s.Italic {
		dst = append(dst, AppendItalicEnd(dst)...)
	}
	if s.Bold {
		dst = append(dst, AppendBoldEnd(dst)...)
	}
	if s.Background != nil {
		dst = append(dst, AppendBackgroundEnd(dst)...)
	}
	if s.Color != nil {
		dst = append(dst, AppendColorEnd(dst)...)
	}
	return dst
}
