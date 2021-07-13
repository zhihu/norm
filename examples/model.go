package examples

import (
	"fmt"

	"github.com/zhihu/norm"
)

type (
	User struct {
		norm.VModel
		ID      int64 `norm:"id"`
		Created int64 `norm:"created"`
	}
	Answer struct {
		norm.VModel
		ID        int64 `norm:"id"`
		VoteUpCnt int64 `norm:"vote_up_cnt"`
	}
	AnswerVoteUp struct {
		norm.EModel
		VoteUpCnt int64 `norm:"vote_up_cnt"`
		Created   int64 `norm:"created"`
	}
)

var _ norm.IVertex = new(User)
var _ norm.IVertex = new(Answer)
var _ norm.IEdge = new(AnswerVoteUp)

func (*User) TagName() string {
	return "user"
}

func (*Answer) TagName() string {
	return "answer"
}

// GetVid 重定义获取 vid 的方法
func (a *Answer) GetVid() interface{} {
	return fmt.Sprintf("answer_%d", a.ID)
}

func (*AnswerVoteUp) EdgeName() string {
	return "answer_vote_up"
}
