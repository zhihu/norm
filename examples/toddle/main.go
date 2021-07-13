package main

import (
	"time"

	"log"

	"github.com/zhihu/norm"
	"github.com/zhihu/norm/constants"
	"github.com/zhihu/norm/dialectors"
	"github.com/zhihu/norm/examples"
)

func main() {
	dalector := dialectors.MustNewNebulaDialector(dialectors.DialectorConfig{
		Addresses: []string{"127.0.0.1:9669"},
		Timeout:   time.Second * 5,
		Space:     "test",
		Username:  "test",
		Password:  "test",
	})
	db := norm.MustOpen(dalector, norm.Config{})
	// prepare(db)
	run(db)
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

func run(db *norm.DB) {
	insert(db)
	match(db)
}

func insert(db *norm.DB) {
	insertVertex(db)
	insertEdge(db)
}

func insertVertex(db *norm.DB) {
	user := &examples.User{
		VModel: norm.VModel{
			Vid:    "user_101",
			Policy: constants.PolicyHash,
		},
		ID:      101,
		Created: time.Now().Unix(),
	}
	err := db.Debug().InsertVertex(user)
	if err != nil {
		log.Fatalf("insert %+v error: %v", user, err)
		panic(err)
	}
}

func insertEdge(db *norm.DB) {
	vote := &examples.AnswerVoteUp{
		EModel: norm.EModel{
			Src:       "user_101",
			SrcPolicy: constants.PolicyHash,
			Dst:       "answer_102",
			DstPolicy: constants.PolicyHash,
		},
		VoteUpCnt: 101,
		Created:   time.Now().Unix(),
	}
	err := db.Debug().InsertEdge(vote)
	if err != nil {
		log.Fatalf("insert %+v error: %v", vote, err)
		panic(err)
	}
}

func match(db *norm.DB) {
	matchSingle(db)
	matchMulti(db)
}

func matchSingle(db *norm.DB) {
	nsql := "match(v:user) where id(v)==hash('user_101') return v.id as id,v.created as created"
	user := examples.User{}
	err := db.Debug().ExecuteAndParse(nsql, &user)
	if err != nil {
		log.Fatalf("exec %s error: %v", nsql, err)
		panic(err)
	}
	log.Printf("%+v", user)
}

func matchMulti(db *norm.DB) {
	nsql := "match(v:user) where id(v)==hash('user_101') or id(v)==hash('user_100')" +
		" return v.id as id,v.created as created"
	users := []examples.User{}
	err := db.Debug().ExecuteAndParse(nsql, &users)
	if err != nil {
		log.Fatalf("exec %s error: %v", nsql, err)
		panic(err)
	}
	log.Printf("%+v", users)
}

// feature support
func chainQuery(db *norm.DB) {
	// db.Debug().From(&examples.User{}).Over(&examples.AnswerVoteUp{}).Reversely().Row()
}
