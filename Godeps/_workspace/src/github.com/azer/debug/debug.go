package debug

import (
	"fmt"
	. "github.com/mroth/emojitrack-gostreamer/Godeps/_workspace/src/github.com/azer/go-style"
	"os"
)

var (
	ctr   int    = 0
	reset string = Style("reset", "")
)

func Debug(text string, args ...interface{}) {
	if len(enabled) == 0 {
		return
	}

	name := caller()
	if !isEnabled(name) {
		return
	}

	color := colorOf(name)

	text = fmt.Sprintf("  %s%s %s %s%s\n", Style(color, name), reset, text, Style(color, diff()), reset)
	fmt.Fprintf(os.Stderr, text, args...)
}
