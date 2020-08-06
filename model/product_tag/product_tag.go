package product_tag

type ProductTag struct {
	TagId     uint64
	StoreId uint64
	Name   string
}

//获取表名
func GetTableName() string {
	return "product_tag"
}

func GetField() []string {
	return []string{
		"tag_id", "store_id", "name","created_by","updated_by","created_at","updated_at",
	}
}

/*
根据商户id和用户id，获取用户信息
注意：如果是简单的数据库操作或者是公共的方法，也可以封装在model中，
当然了也可以封装在servicelogic中，这个根据业务场景来决定。
*/
/*
func GetUserListByUserId(userName string, businessId uint64) (*Users, error) {
	userList := &Users{}
	err := db.Conn.Table(GetTableName()).Select(GetField()).
		Where("user_id = ? AND business_id = ?", userName, businessId).
		Find(userList).Error
	if err != nil {
		return nil, fmt.Errorf("select memeber info is fail, err: %v", err)
	}

	return userList, nil
}
*/
