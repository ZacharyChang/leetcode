package leetcode

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 3))
}

func TestConvert2(t *testing.T) {
	s := "A"
	fmt.Println(convert(s, 1))
}
