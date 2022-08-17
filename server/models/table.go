package models

import (
	"time"
)

type UserRole struct {
	ID       int    `gorm:"primary_key"`
	RoleName string `gorm:"index;type:varchar(100)"`
	RoleUser []User `gorm:"foreignKey:Role;references:ID"`
	RoleMenu []Menu `gorm:"foreignKey:Role;references:ID"`
	RoleLogs []Log  `gorm:"foreignKey:Role;references:ID"`
}

type User struct {
	Id           int       `primaryKey:"true"`                     // 自增ID
	UserName     string    `gorm:"type:varchar(100);unique_index"` // 用户名
	Portrait     string    `gorm:"type:varchar(100)"`              // 头像
	Password     string    `gorm:"type:varchar(256)"`              // 密码
	Emial        string    `gorm:"type:varchar(100);"`             // 邮箱
	Token        string    `gorm:"type:varchar(256)"`              // token
	Verification string    `gorm:"type:varchar(100)"`              // 邮箱发送的验证码
	Role         uint      // 角色ID
	Frequency    int       `gorm:"type:int(255)"`              // 访问频率
	CreatedAt    time.Time `time_format:"2006-01-02 15:04:05"` // 创建时间
	UpdatedAt    time.Time `time_format:"2006-01-02 15:04:05"` // 更新时间
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
	Icon     string //图标
	ParentId int    `gorm:"default:0"` //父级菜单ID
	Role     uint   //角色
}

type Upload struct {
	Id        int       `primaryKey:"true"`
	Name      string    //文件名称
	Path      string    `json:"path"`                       //文件路径
	Size      int64     `json:"size"`                       //文件大小
	Type      string    `json:"type"`                       //文件类型
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"` // 创建时间
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"` // 更新时间
}

// Log 日志表
type Log struct {
	Id        int       `primaryKey:"true"`
	User      string    //用户名
	UserId    int       //用户ID
	Role      uint      //角色
	Ip        string    //IP
	Catetory  string    //日志类型  登录、操作、异常
	Type      string    //日志类型
	Content   string    //日志内容
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"` // 创建时间
}

// ArticleCategory 文章类型
type ArticleCategory struct {
	Id       int       `primaryKey:"true"`
	Name     string    `gorm:"index;type:varchar(100)"`
	Category []Article `gorm:"foreignKey:Category;references:Name"` //文章分类
}
type ArticleTag struct {
	Id int `primaryKey:"true" json:"id"`
	// 不重复
	Name string `gorm:"unique;type:varchar(100)" json:"name"`
}

// Article 文章
type Article struct {
	Id        int          `gorm:"primary_key" json:"id"`
	Title     string       `json:"title" ` // 标题
	Tag       []ArticleTag `gorm:"many2many:tags" json:"tag"`
	Category  string       `gorm:"type:varchar(100) " json:"category"`
	Author    string       `gorm:"type:varchar(100) " json:"author"`
	Content   string       `json:"content"`
	Show      bool         `gorm:"default:true" json:"show"`   //是否显示
	CreatedAt time.Time    `time_format:"2006-01-02 15:04:05"` // 创建时间
	UpdatedAt time.Time    `time_format:"2006-01-02 15:04:05"` // 更新时间
}
