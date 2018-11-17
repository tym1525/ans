package module

import (
	"time"
)

type User struct {
	Id         int64     `orm:"auto"`         // 数据库自增Id
	UserId     uint64    `orm:"index;unique"` // 用户Id
	UserName   string    // 用户名
	NickName   string    // 昵称
	Password   string    // 密码
	Sex        int8      // 性别
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` // 创建时间
	UpdateTime time.Time // 更新时间
}
