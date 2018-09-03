package coloredgoroutine

import (
	"fmt"
	"math"
	"os"
	"sync"
	"testing"

	"github.com/xiegeo/coloredgoroutine/goid"
)

func TestGetColorForID(t *testing.T) {
	t.Log(getColorForID(math.MinInt32).Sprintf("%d", math.MinInt32))
	t.Log(getColorForUID(math.MaxUint64).Sprintf("%d", uint64(math.MaxUint64)))

	for i := 0; i < len(colors); {
		s := []interface{}{}
		for {
			s = append(s, getColorForID(i).Sprintf("%3d", i))
			//c := &colors[i]
			//s = append(s, c.Sprintf("%3d%3d", c.fg, c.bg-10))
			i++
			if i%24 == 0 || i == len(colors) {
				break
			}
		}
		t.Log(s...)
	}
}
func TestPrintBannedColors(t *testing.T) {
	for i := 0; i < len(bannedColors); {
		s := []interface{}{}
		for i < len(bannedColors) {
			c := &bannedColors[i]
			s = append(s, c.Sprintf("%3d%3d", c.fg, c.bg-10))
			i++
			if i%13 == 0 {
				break
			}
		}
		t.Log(s...)
	}
}

func TestColorsInGoRoutines(t *testing.T) {

	c := Colors(os.Stdout)

	fmt.Fprintln(c, "Hi, I am go routine", goid.ID(), "from test routine")

	if testing.Short() {
		t.Skip("skip the longer version")
	}

	count := 100

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		i := i
		go func() {
			fmt.Fprintln(c, "Hi, I am go routine", goid.ID(), "from loop i =", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
