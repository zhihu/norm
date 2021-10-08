package norm

import (
	"github.com/zhihu/norm/dialectors"
	"github.com/zhihu/norm/internal/converts"
)

// Execute 执行某条sql
func (db *DB) Execute(sql string) (*dialectors.ResultSet, error) {
	return db.execute(sql)
}

// ExecuteAndParse 执行某条sql, 然后解析为传入的结构.
// in 可以是 map[string]interface{}, *Strcut, *[]map, *[]struct
func (db *DB) ExecuteAndParse(sql string, in interface{}) error {
	nResult, err := db.execute(sql)
	if err != nil {
		return err
	}
	return converts.UnmarshalResultSet(nResult, in)
}

// InsertVertex 解析结构体, 然后插入一个点.
func (db *DB) InsertVertex(in IVertex) error {
	sql, err := converts.ConvertToCreateVertexSql(in, in.TagName(), in.GetVid(), in.GetPolicy())
	if err != nil {
		return err
	}
	_, err = db.execute(sql)
	return err
}

// InsertEdge 解析结构体, 然后插入一条边.
func (db *DB) InsertEdge(in IEdge) error {
	sql, err := converts.ConvertToCreateEdgeSql(in, in.EdgeName(), in.GetVidSrc(), in.GetVidDst(),
		in.GetVidSrcPolicy(), in.GetVidDstPolicy())
	if err != nil {
		return err
	}
	_, err = db.execute(sql)
	return err
}

// ReturnRow 返回 nsql 执行的原始结果, 不进行任何加工
func (db *DB) ReturnRow() (*dialectors.ResultSet, error) {
	return db.execute(db.sql)
}

// Return 返回数据并将结果反序列化到 out 结构体中
func (db *DB) Return(out interface{}) error {
	return db.ExecuteAndParse(db.sql, out)
}

// Count 计算符合条件的数据数量
func (db *DB) Count(out interface{}) error {
	// 不用 match 的 count 原因是部分情况下有 bug
	// 参考 https://github.com/vesoft-inc/nebula/issues/2934
	return db.Yield(" '' as id").Group("id").Yield("count(1)").Return(out)
}
