package querys

import (
	"log"

	"github.com/zhihu/norm/v3"
	"github.com/zhihu/norm/v3/examples/models"
)

func MatchSingle(db *norm.DB) {
	nsql := "match(v:user) where id(v)==hash('user_101') return v.id as id,v.created as created"
	user := models.User{}
	err := db.Debug().ExecuteAndParse(nsql, &user)
	if err != nil {
		log.Fatalf("exec %s error: %v", nsql, err)
		panic(err)
	}
	log.Printf("%+v", user)
}

func MatchMulti(db *norm.DB) {
	nsql := "match(v:user) where id(v)==hash('user_101') or id(v)==hash('user_100')" +
		" return v.id as id,v.created as created"
	users := []models.User{}
	err := db.Debug().ExecuteAndParse(nsql, &users)
	if err != nil {
		log.Fatalf("exec %s error: %v", nsql, err)
		panic(err)
	}
	log.Printf("%+v", users)
}
