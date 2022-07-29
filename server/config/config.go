package config

func Config() {
	CreateDir("./media/upload/user") // 创建用户头像文件夹·
	CreateDir("./media/upload/img")  // 创建文章图片文件夹·
}

// Redis 配置
type redisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// 发送邮件配置
type emailConfig struct {
	Host     string `json:"host"`     // smtp服务器地址
	Port     int    `json:"port"`     // smtp服务器端口
	Username string `json:"username"` // 邮箱用户名
	From     string `json:"from"`     // 发件人地址
	Password string `json:"password"` // 邮箱密钥

}
