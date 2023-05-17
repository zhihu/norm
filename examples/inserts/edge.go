package inserts

import (
	"log"
	"time"

	"github.com/zhihu/norm/v3"
	"github.com/zhihu/norm/v3/constants"
	"github.com/zhihu/norm/v3/examples/models"
)

func InsertEdge(db *norm.DB) {
	vote := &models.AnswerVoteUp{
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
