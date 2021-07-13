package converts

type User struct {
	ID   int    `norm:"id"`
	Name string `norm:"name"`
}

// func TestConvertToInsertSql(t *testing.T) {
// 	user := &User{
// 		ID:   100,
// 		Name: "nebula",
// 	}
// 	sql, _ := ConvertToInsertSql(user, "user", "user_100")
// 	t.Log(sql)
// }
