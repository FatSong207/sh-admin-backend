package config

type Config struct {
	Port string
	Redis
	Mysql
	Zap
	Jwt
}

// Redis Redis配置
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

// Mysql MySQL配置
type Mysql struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

// Zap Zap配置
type Zap struct {
	Level        string `mapstructure:"level"`          // 级别
	Prefix       string `mapstructure:"prefix"`         // 日志前缀
	Format       string `mapstructure:"format"`         // 输出
	Director     string `mapstructure:"director"`       // 日志文件夹
	LogInConsole bool   `mapstructure:"log-in-console"` // 输出控制台
}

// Jwt JWT配置
type Jwt struct {
	SigningKey  string `mapstructure:"signingKey"`
	ExpiredTime int    `mapstructure:"expiredTime"`
}
