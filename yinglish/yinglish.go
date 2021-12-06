package yinglish

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"html"
	"math/rand"
	"strconv"
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
			if random > yld/2 {
				return "💗"
			}
			return "…"
		}
		if in(s, []string{"!", "！"}) {
			return "❤️!"
		}
		if in(s, []string{"、", "；"}) {
			return "_"
		}
		if in(s, []string{"？", "？"}) {
			return "¿"
		}
		if len(s) > 1 && random > yld/2 {
			return fmt.Sprintf("%v～%v", strings.Split(s, "")[0], s)
		}
		if flag == "v" {
			if random > yld/2 {
				return fmt.Sprintf("👺%v", s)
			}
			return fmt.Sprintf("~%vx", s)
		} else if flag == "n" {
			if random > yld/2 {
				return fmt.Sprintf("🌟%v🌟", s)
			}
			return strings.Repeat("〇", len(s))
		} else if flag == "ul" || flag == "eng" || flag == "m" {
			return fmt.Sprintf("%v❤️~", s)
		} else if flag == "ez" || flag == "r" || flag == "d" {
			if random > yld/3 {
				return fmt.Sprintf("%v️👄", s)
			}
			return fmt.Sprintf("%v❕~", s)
		} else {
			if random > yld/3 {
				return fmt.Sprintf("%v️%v", s, randomEmoji())
			}
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
		//fmt.Println(a[0], a[1])
		result = fmt.Sprintf("%v%v", result, tans(a[0], a[1], yld))
	}
	return result
}

func randomEmoji() string {
	rand.Seed(time.Now().UnixNano())
	// http://apps.timwhitlock.info/emoji/tables/unicode
	emoji := [][]int{
		// Emoticons icons
		{128513, 128591},
		// Transport and map symbols
		{128640, 128704},
	}
	r := emoji[rand.Int()%len(emoji)]
	min := r[0]
	max := r[1]
	n := rand.Intn(max-min+1) + min
	return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
}
