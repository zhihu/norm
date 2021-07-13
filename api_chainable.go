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

func (db *DB) From(v IVertex) (tx *DB) {
	tx = db.getInstance()
	tx.sql += fmt.Sprintf("go from %s ", v.TagName())
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
	tx.sql += fmt.Sprintf("where %s ", sql)
	return
}
