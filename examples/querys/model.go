package querys

import (
	"time"

	"github.com/zhihu/norm"
	"github.com/zhihu/norm/constants"
	"github.com/zhihu/norm/examples/models"
)

var (
	userExample = &models.User{
		VModel: norm.VModel{
			Vid:    "user_101",
			Policy: constants.PolicyHash,
		},
		ID:      101,
		Created: time.Now().Unix(),
	}
	voteExample = &models.AnswerVoteUp{
		EModel: norm.EModel{
			Src:       "user_101",
			SrcPolicy: constants.PolicyHash,
			Dst:       "answer_102",
			DstPolicy: constants.PolicyHash,
		},
		VoteUpCnt: 101,
		Created:   time.Now().Unix(),
	}
)
