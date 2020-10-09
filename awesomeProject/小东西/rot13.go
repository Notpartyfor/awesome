package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(out byte) byte { //字母转换
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}
	return out
}

func (fz rot13Reader) Read(b []byte) (int, error) { //重写Read方法
	n, e := fz.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	//fmt.Println(&r)
	io.Copy(os.Stdout, &r)
	io.Copy(os.Stdout, strings.NewReader("hhhh"))
}
