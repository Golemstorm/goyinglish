package yinglish

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"math/rand"
	"strings"
	"time"
)

func in(a string, b []string) bool {
	for _, v := range b {
		if v == a {
			return true
		}
	}
	return false
}
func tans(s, flag string, yld float64) string {
	rand.Seed(time.Now().Unix() + rand.Int63())
	random := rand.Float64()
	if random > yld {
		return s
	} else {
		if in(s, []string{",", "，", ".", "。"}) {
			return "……"
		}
		if in(s, []string{"!", "！"}) {
			return "❤!"
		}
		if len(s) > 1 && random < yld/2 {
			return fmt.Sprintf("%v……%v", strings.Split(s, "")[0], s)
		}
		if flag == "v" {
			return fmt.Sprintf("~%vx", s)
		} else if flag == "n" {
			return strings.Repeat("〇", len(s))
		} else if flag == "ul" || flag == "eng" {
			return fmt.Sprintf("%v❤~", s)
		} else if flag == "ez" || flag == "r" {
			return fmt.Sprintf("%v!~", s)
		} else {
			return fmt.Sprintf("…%v", s)
		}
	}
}

func Chs2yin(s string, yld float64) string {
	x := gojieba.NewJieba()
	defer x.Free()
	words := x.Tag(s)
	result := ""
	for _, v := range words {
		a := strings.Split(v, "/")
		result = fmt.Sprintf("%v%v", result, tans(a[0], a[1], yld))
	}
	return result
}
