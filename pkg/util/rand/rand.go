package rand

import (
	"math/rand"
	"time"
)

// GenerateUID 这个 UID 通过 math/rand 包和时间戳种子生成伪随机最大 10 位的 UID
func GenerateUID() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Int63n(9999999999)
	println(num)
}
