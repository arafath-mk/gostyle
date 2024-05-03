package gcstyle

import "github.com/arafath-mk/gostyle/wcolor"

type StyledString struct {
	str        string
	applyStyle bool
}

// Applies foreground color.
func (s StyledString) Color(c wcolor.Color) StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = foreground(s.str, c)
	return s
}

// Applies background color.
func (s StyledString) Background(c wcolor.Color) StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = background(s.str, c)
	return s
}

// Applies bold style.
func (s StyledString) Bold() StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = bold(s.str)
	return s
}

func (s StyledString) Darken() StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = darken(s.str)
	return s
}

func (s StyledString) Lighten() StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = lighten(s.str)
	return s
}

// Applies italic style.
func (s StyledString) Italic() StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = italic(s.str)
	return s
}

// Applies underline style.
func (s StyledString) Underline() StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = underline(s.str)
	return s
}

// Applies strikethrough style.
func (s StyledString) Strikethrough() StyledString {
	if !s.applyStyle {
		return s
	}

	s.str = strikethrough(s.str)
	return s
}

// Returns string.
func (s StyledString) String() string {
	return s.str
}
