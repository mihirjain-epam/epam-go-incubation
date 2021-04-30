package main

import (
	"bytes"
	"fmt"
)

type path []byte

// Any change made to the path variable p will not represent in the original variable as this is the copy
func (p path) NoPointerTruncateAtFinalSlash() {
	i := bytes.LastIndex(p, []byte("/"))
	if i >= 0 {
		p = (p)[0:i]
	}
	fmt.Printf("no pointer - %s\n", p)
}

// taking a pointer to the path variable it will make changes to the original variable
func (p *path) PointerTruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
	fmt.Printf("pointer - %s\n", *p)
}

// In this case doesn't matter if we pass value or pointer as both point to the same underlying array
func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	pathName.ToUpper()
	fmt.Printf("original - %s\n", pathName)
	pathName.NoPointerTruncateAtFinalSlash()
	fmt.Printf("original - %s\n", pathName)
	pathName.PointerTruncateAtFinalSlash()
	fmt.Printf("original - %s\n", pathName)
}
