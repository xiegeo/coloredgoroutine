package coloredgoroutine

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/xiegeo/coloredgoroutine/goid"
)

type writer struct {
	w io.Writer
}

func (w *writer) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	c := getColorForUID(goid.ID())
	if p[len(p)-1] == '\n' {
		p = p[:len(p)-1]
		return fmt.Fprintln(w.w, c.Sprintf("%s", p))
	}
	return c.Fprintf(w.w, "%s", p)
}

//Colors add color to the writer based on the current go routine id
func Colors(w io.Writer) io.Writer {
	return &writer{
		w: w,
	}
}

var colors []color.Color

func init() {
	same := color.BgBlack - color.FgBlack
	for f := color.FgBlack; f <= color.FgWhite; f++ {
		for b := color.BgBlack; b <= color.BgWhite; b++ {
			if b-f == same {
				continue
			}
			hi := color.FgHiBlack - color.FgBlack
			colors = append(colors, *color.New(f, b))
			colors = append(colors, *color.New(f+hi, b))
			colors = append(colors, *color.New(f, b+hi))
			colors = append(colors, *color.New(f+hi, b+hi))
		}
	}
}

func getColorForUID(id uint64) *color.Color {
	return getColorForID(int(id))
}

func getColorForID(id int) *color.Color {
	id = id % len(colors)
	if id < 0 {
		id += len(colors)
	}
	return &colors[id]
}
