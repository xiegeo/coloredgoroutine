package coloredgoroutine

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/xiegeo/coloredgoroutine/goid"
)

func TestGetColorForID(t *testing.T) {
	for i := -28; i < len(colors); {
		s := []interface{}{}
		for {
			s = append(s, getColorForID(i).Sprintf("%3d", i))
			i++
			if i%28 == 0 {
				break
			}
		}
		t.Log(s...)
	}
}

func TestColorsInGoRoutines(t *testing.T) {

	count := 100

	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < count; i++ {
		i := i
		go func() {
			fmt.Fprintln(Colors(os.Stdout), "Hi, I am go routine", goid.ID(), "from loop i =", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
