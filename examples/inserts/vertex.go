package inserts

import (
	"log"
	"time"

	"github.com/zhihu/norm/v3"
	"github.com/zhihu/norm/v3/constants"
	"github.com/zhihu/norm/v3/examples/models"
)

func InsertVertex(db *norm.DB) {
	user := &models.User{
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
