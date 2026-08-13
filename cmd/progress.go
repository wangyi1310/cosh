package cmd

import (
	"fmt"
	"strings"
)

// ANSI color codes
const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorCyan    = "\033[36m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorWhite   = "\033[97m"
	colorDim     = "\033[2m"
)

type progressWriter struct {
	total   int64
	current int64
	label   string
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.current += int64(n)
	pw.print()
	return n, nil
}

func (pw *progressWriter) print() {
	if pw.total <= 0 {
		// Unknown total - show spinner style
		spinners := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		idx := int(pw.current/1024) % len(spinners)
		fmt.Printf("\r%s %s%s%s  %s%s%s downloaded",
			spinners[idx],
			colorCyan, pw.label, colorReset,
			colorYellow, formatSize(pw.current), colorReset)
		return
	}

	pct := float64(pw.current) / float64(pw.total) * 100
	barW := 30
	filled := int(pct / 100 * float64(barW))
	if filled > barW {
		filled = barW
	}

	// Gradient color based on progress
	barColor := colorCyan
	if pct > 60 {
		barColor = colorGreen
	} else if pct > 30 {
		barColor = colorBlue
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", barW-filled)
	fmt.Printf("\r%s %s%s%s %s%5.1f%% %s %s/%s%s",
		colorCyan+pw.label+colorReset,
		barColor, bar, colorReset,
		colorWhite, pct, colorReset,
		colorDim, formatSize(pw.current)+"/"+formatSize(pw.total), colorReset)
}

func (pw *progressWriter) finish() {
	fmt.Printf("\r%s %s✓ Complete%s  %s%s%s                    \n",
		colorCyan+pw.label+colorReset,
		colorGreen, colorReset,
		colorDim, formatSize(pw.total), colorReset)
}
