package yinglish

import (
	"fmt"
	"testing"
)

func TestChs2yin(t *testing.T) {
	got := Chs2yin("XX君，那里不可以，要去了，那里要被塞满了", 0.5)
	fmt.Println(got)
}
