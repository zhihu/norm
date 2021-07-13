package norm

import "github.com/zhihu/norm/constants"

// 预定名词必须实现的属性
type (
	ITag interface {
		TagName() string
	}
	IVertex interface {
		// 返回点的名称
		ITag
		GetVid() interface{}
		GetPolicy() constants.Policy
	}
	IEdge interface {
		// 返回边的名称
		EdgeName() string
		GetVidSrc() interface{}
		GetVidSrcPolicy() constants.Policy
		GetVidDst() interface{}
		GetVidDstPolicy() constants.Policy
	}
)

type (
	// Vid 是 int 或者 string 类型
	VModel struct {
		Vid    interface{}      `norm:"-"`
		Policy constants.Policy `norm:"-"`
	}
	EModel struct {
		Src       interface{}      `norm:"-"`
		SrcPolicy constants.Policy `norm:"-"`
		Dst       interface{}      `norm:"-"`
		DstPolicy constants.Policy `norm:"-"`
	}
)

// GetVid 获取 vid
func (v VModel) GetVid() interface{} {
	return v.Vid
}

func (v VModel) GetPolicy() constants.Policy {
	return v.Policy
}

func (v EModel) GetVidSrc() interface{} {
	return v.Src
}

func (v EModel) GetVidSrcPolicy() constants.Policy {
	return v.SrcPolicy
}

func (v EModel) GetVidDst() interface{} {
	return v.Dst
}

func (v EModel) GetVidDstPolicy() constants.Policy {
	return v.DstPolicy
}
