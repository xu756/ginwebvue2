package models

import "time"

type UserRole struct {
	ID       int    `gorm:"primary_key"`
	RoleName string `gorm:"index;type:varchar(100)"`
	Role     []User `gorm:"foreignKey:Role;references:RoleName"`
}

type User struct {
	Id           int       `primaryKey:"true"`                     // 自增ID
	UserName     string    `gorm:"type:varchar(100);unique_index"` // 用户名
	Portrait     string    `gorm:"type:varchar(100)"`              // 头像
	Password     string    `gorm:"type:varchar(256)"`              // 密码
	Emial        string    `gorm:"type:varchar(100);"`             // 邮箱
	Token        string    `gorm:"type:varchar(256)"`              // token
	Verification string    `gorm:"type:varchar(100)"`              // 邮箱发送的验证码
	Role         string    `gorm:"type:varchar(100);default:'用户'"` // 角色
	Frequency    int       `gorm:"type:int(255)"`                  // 访问频率
	CreatedAt    time.Time `time_format:"2006-01-02 15:04:05"`     // 创建时间
	UpdatedAt    time.Time `time_format:"2006-01-02 15:04:05"`     // 更新时间
}

type Emial struct {
	Id        int       `primaryKey:"true"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"` // 发送时间
	To        string    //接收者
	Subject   string    //主题
	Body      string    //内容
}

type Menu struct {
	Id       int    `primaryKey:"true"`
	Name     string //菜单名称
	Path     string //页面路径
	Icon     string
	ParentId int    //父级菜单
	Role     string //用户角色
}
