package user_detail

type UserDetail struct {
	UserId   uint64
	UserName string
	Sex      uint8
	Age      uint8
}

//获取表名
func GetTableName() string {
	return "user_detail"
}

func GetField() []string {
	return []string{
		"user_id", "user_name", "sex", "age",
	}
}
