# goyinglish

yinglish go版本
创意来源于https://github.com/RimoChan/yinglish



## 样例

```golang
package main

import (
	"fmt"
	"github.com/Golemstorm/goyinglish/yinglish"
)
func main() {
    a := `不行，那里不可以，哒咩，葡萄，葡萄皮要爆掉了!`
    fmt.Println(yinglish.Chs2yin(a,0.8))
    //不行……那里!~不可……可以，…哒咩……葡……葡萄……葡……葡萄皮……皮要…爆掉了❤!
}
```

## TODO
 可以再色点，相比 RimoChan/yinglish 多了点词性判断，可以莫多涩涩
