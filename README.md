difftool
====
Very Simple package to launch difftool.

If you spent a lot of time trying to find difference between expected and actual output in you tests (just like I did), maybe this simple package will help you.

difftool will run program in DIFFTOOL environment variable with your input.

## Demo:

![Screen Cast](https://wix-shareable.s3.amazonaws.com/difftool.gif)


## Usage:
Olny one function exists:
```
func Open(a, b []byte) error
```

## Example:
```go
package exmple

import (
  "github.com/gregory-m/difftool"

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
      difftool.Open([]byte(s), []byte(tt.out))
    }
  }
}

```
