package difftool

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	var myTests = []struct {
		in  string
		out string
	}{
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
		{"Phasellus tristique lacinia eros, ut faucibus nisl lacinia non.", "Phasellus tristique lacinia eros, ut faucibus nisl lacinia non. "},
	}

	for _, tt := range myTests {
		s := strings.Replace(tt.in, "o", "a", -1)
		if s != tt.out {
			t.Errorf("Oy-Vey for: '%q'. get: '%q', want: '%q'", tt.in, s, tt.out)
			Open([]byte(s), []byte(tt.out))
		}
	}
}
