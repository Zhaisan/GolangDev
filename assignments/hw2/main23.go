package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func rot13(b byte) (res byte) {
	
    switch {
	case b >= 'A' && b <= 'Z':
		res = (b - 'A' + 13) % 26 + 'A'
	case b >= 'a' && b <= 'z':
		res = (b - 'a' + 13) % 26 + 'a'
	default:
		res = b  
	}
	return
}

func(x *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = x.r.Read(p)
    for i := range(p) {
      p[i] = rot13(p[i])
    }
    return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}