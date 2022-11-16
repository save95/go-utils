package maker

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// AppID 生成 APP ID，除 prefix 外长度15位，组成结构参考 AppIDWith
func AppID(prefix string) string {
	return AppIDWith(prefix, 15)
}

// AppIDWith 生成 APP ID，长度12-20位，受第二参数 length 约束，组成结构如下：
// ｜00|111|22222|33333|
//   ｜  |    |     |
//   ｜  |    |     `-------- 当日的秒数（5位）
//   ｜  |    `-------------- 随机数补位（2-10位）
//   ｜  `------------------- 当日是一年中的第几日，[000, 366]（3位）
//    `---------------------- 年份最后2位（2位）
func AppIDWith(prefix string, length uint) string {
	baseLen := 10 // 除随机字符串之外的长度
	rnLen := int(math.Min(20, math.Max(12, float64(length)))) - baseLen

	bytes := make([]byte, 20)
	rand.Seed(time.Now().UnixNano())
	rand.Read(bytes)
	symbolsByteLength := byte(len(_symbols))
	for i, b := range bytes {
		bytes[i] = _symbols[b%symbolsByteLength]
	}
	rn := string(bytes)[0:rnLen]

	now := time.Now()

	// 当日的秒数
	sec := (now.Hour()*60+now.Minute())*60 + now.Second()

	return fmt.Sprintf(
		"%s%02d%03d%s%05d",
		prefix,
		now.Year()-2000,
		now.YearDay(),
		rn,
		sec,
	)
}
