package config

// Redis 配置
type redisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// Mysql 配置
type mysqlConfig struct {
	Addr     string `json:"mysql"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

// 发送邮件配置
type emailConfig struct {
	Host     string `json:"host"`     // smtp服务器地址
	Port     int    `json:"port"`     // smtp服务器端口
	Username string `json:"username"` // 邮箱用户名
	From     string `json:"from"`     // 发件人地址
	Password string `json:"password"` // 邮箱密钥

}

// 微信公众号配置
type wxchatConfig struct {
	AppID          string `json:"appid"`          // 公众号appid
	AppSecret      string `json:"appsecret"`      // 应用密钥
	Token          string `json:"token"`          // 公众号token
	EncodingAESKey string `json:"encodingaeskey"` // 加解密密钥
}
