package querys

import (
	"time"

	"github.com/zhihu/norm/v3"
	"github.com/zhihu/norm/v3/constants"
	"github.com/zhihu/norm/v3/examples/models"
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
