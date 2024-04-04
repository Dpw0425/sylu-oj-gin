package config

type Config struct {
	Mode  Mode  `mapstructure:"mode"`
	Http  Http  `mapstructure:"http"`
	Log   Log   `mapstructure:"log"`
	Mysql Mysql `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
	Gorm  Gorm  `mapstructure:"gorm"`
	Limit Limit `mapstructure:"limit"`
	Cors  Cors  `mapstructure:"cors"`
	Jwt   Jwt   `mapstructure:"jwt"`
	Email Email `mapstructure:"email"`
}

type Mode struct {
	RunMode string `mapstructure:"runmode"`
}

type Http struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Log struct {
	Level    int    `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Output   string `mapstructure:"output"`
	Lowfile  string `mapstructure:"lowfile"`
	Highfile string `mapstructure:"highfile"`
}

type Mysql struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	DBname     string `mapstructure:"dbname"`
	Parameters string `mapstructure:"parameters"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type Gorm struct {
	Debug  bool   `mapstructure:"debug"`
	DBType string `mapstructure:"dbtype"`
}

type Limit struct {
	Limit int64 `mapstructure:"limit"`
}

type Cors struct {
	Enable           bool          `mapstructure:"enable"`
	AllowOrigins     []interface{} `mapstructure:"alloworigins"`
	AllowMethods     []interface{} `mapstructure:"allowmethods"`
	AllowHeaders     []interface{} `mapstructure:"allowheaders"`
	AllowCredentials bool          `mapstructure:"allowcredentials"`
	MaxAge           int           `mapstructure:"maxage"`
}

type Jwt struct {
	Enable        bool   `mapstructure:"enable"`
	SigningMethod string `mapstructure:"signingmethod"`
	SigningKey    string `mapstructure:"signingkey"`
	Issuer        string `mapstructure:"issuer"`
	Expired       int    `mapstructure:"expired"`
	Store         string `mapstructure:"store"`
	FilePath      string `mapstructure:"filepath"`
	RedisDB       int    `mapstructure:"redisdb"`
	RedisPrefix   string `mapstructure:"redisprefix"`
}

type Email struct {
	Host     string `mapstructure:"host"`
	Smtp     string `mapstructure:"smtp"`
	Addr     string `mapstructure:"addr"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Expires  int    `mapstructure:"expires"`
}

func DSN(m Mysql) string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBname + "?" + m.Parameters
}
