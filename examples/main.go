package main

import (
	"log"
	"time"

	"github.com/zhihu/norm/v3"
	"github.com/zhihu/norm/v3/dialectors"
	"github.com/zhihu/norm/v3/examples/querys"
)

func main() {
	db := newGdb()
	// prepare()

	// // 插入示例
	// inserts.InsertEdge(db)
	// inserts.InsertVertex(db)

	// // 查询 示例
	// querys.MatchSingle(db)
	// querys.MatchMulti(db)
	// querys.Count(db)
	querys.Chainable(db)
	// querys.ChainableCount(db)
}

func newGdb() *norm.DB {
	dalector := dialectors.MustNewNebulaDialector(dialectors.DialectorConfig{
		Addresses: []string{"127.0.0.1:9669"},
		Timeout:   time.Second * 5,
		Space:     "test",
		Username:  "test",
		Password:  "test",
	})
	db := norm.MustOpen(dalector, norm.Config{})
	return db
}

func prepare(db *norm.DB) {
	// 创建 tag
	createSchema := "" +
		"CREATE TAG IF NOT EXISTS user(id int, name string);" +
		"CREATE TAG IF NOT EXISTS answer(id int, vote_up_cnt int);" +
		"CREATE EDGE IF NOT EXISTS answer_vote_up(vote_up_cnt int, created timestamp);"
	_, err := db.Execute(createSchema)
	if err != nil {
		log.Fatalf("exec %s error: %v", createSchema, err)
		panic(err)
	}
}
