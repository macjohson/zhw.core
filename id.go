package zhwcore

import (
	"math/rand"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// 使用雪花算法获取唯一ID标识
func NewId() string {
	once.Do(func() {
		for retry := 0; retry < 10; retry++ {
			n, err := snowflake.NewNode(int64(rand.Int() % 1024))
			if err != nil {
				continue
			}
			node = n
			break
		}
		if node == nil {
			panic("snowflake id generate error")
		}
	})
	return string(node.Generate().Bytes())
}
