package querys

import (
	"log"

	"github.com/zhihu/norm"
)

func Count(db *norm.DB) {
	nsql := "match(v:user)-[e]->(v2) where id(v)==hash('user_101') return count(e)"
	cnt := 0
	err := db.Debug().ExecuteAndParse(nsql, &cnt)
	if err != nil {
		log.Fatalf("exec %s error: %v", nsql, err)
		panic(err)
	}
	log.Printf("count: %d", cnt)
}
