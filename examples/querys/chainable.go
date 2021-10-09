package querys

import (
	"fmt"
	"log"
	"time"

	"github.com/zhihu/norm"
)

// 使用 chainable 必须保证函数调用顺序, 因为逻辑上是有序的.

func Chainable(db *norm.DB) {
	cnt := int64(0)
	queryWhere := fmt.Sprintf("%s.created > %d", voteExample.EdgeName(), time.Now().Unix())
	err := db.Debug().From(userExample).Over(voteExample).Bidirect().
		Where(queryWhere).Yield("'' as id").
		Group("id").Yield("count(1)").Return(&cnt)
	if err != nil {
		panic(err)
	}
	log.Printf("count: %d", cnt)
}

func ChainableCount(db *norm.DB) {
	cnt := int64(0)
	err := db.Debug().From(userExample).Over(voteExample).Bidirect().Count(&cnt)
	if err != nil {
		panic(err)
	}
	log.Printf("count: %d", cnt)
}
