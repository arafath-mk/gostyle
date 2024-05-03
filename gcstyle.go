package gcstyle

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

func ApplyTo(str string, applyStyle bool) StyledString {
	return StyledString{
		str:        str,
		applyStyle: applyStyle,
	}
}

func CanApplyStyle(writer io.Writer) bool {
	w, ok := writer.(*os.File)
	if !ok {
		return false
	}

	if os.Getenv("TERM") == "dumb" {
		return false
	}

	if !isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd()) {
		return false
	}

	return true
}
