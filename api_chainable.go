package norm

import (
	"fmt"
	"strings"

	"github.com/zhihu/norm/constants"
)

// Debug 执行时会输出 nsql 内容
func (db *DB) Debug() (tx *DB) {
	tx = db.getInstance()
	tx.debug = true
	return
}

func (db *DB) GoFrom(v IVertex, step int) (tx *DB) {
	tx = db.getInstance()

	vidStr := ""
	vid := v.GetVid()
	switch vid := vid.(type) {
	case int, int8, int32, int64, float32, float64:
		vidStr = fmt.Sprint(vid)
	case string:
		vidStr = "'" + vid + "'"
	default:
		vidStr += "'" + fmt.Sprint(vid) + "'"
	}
	switch v.GetPolicy() {
	case constants.PolicyHash:
		vidStr = "hash(" + vidStr + ")"
	}

	if step == 0 {
		tx.sql += fmt.Sprintf("go from %s ", vidStr)
	} else {
		tx.sql += fmt.Sprintf("go %d step from %s ", step, vidStr)
	}

	return
}

func (db *DB) Over(edges ...IEdge) (tx *DB) {
	tx = db.getInstance()
	names := make([]string, len(edges))
	for i, edge := range edges {
		names[i] = edge.EdgeName()
	}
	sql := strings.Join(names, ",")
	tx.sql += fmt.Sprintf("over %s ", sql)
	return
}

// Reversely 沿着边反向遍历
func (db *DB) Reversely() (tx *DB) {
	tx = db.getInstance()
	tx.sql += constants.DirectionReversely + " "
	return
}

// Bidirect 沿着边双向遍历
func (db *DB) Bidirect() (tx *DB) {
	tx = db.getInstance()
	tx.sql += constants.DirectionBidirect + " "
	return
}

func (db *DB) Limit(limit int) (tx *DB) {
	tx = db.getInstance()
	tx.sql += fmt.Sprintf("|limit %d ", limit)
	return
}

func (db *DB) Where(sql string) (tx *DB) {
	tx = db.getInstance()
	tx.sql += fmt.Sprintf("where %s ", sql)
	return
}

func (db *DB) Yield(sql string) (tx *DB) {
	tx = db.getInstance()
	tx.sql += fmt.Sprintf("yield %s ", sql)
	return
}

func (db *DB) Group(fields ...string) (tx *DB) {
	tx = db.getInstance()
	for i := range fields {
		fields[i] = "$-." + fields[i]
	}
	sql := strings.Join(fields, ",")
	tx.sql += fmt.Sprintf("|group by %s ", sql)
	return
}
