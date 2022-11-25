package conf

import (
	_ "embed"
	"fmt"
)

//go:embed config.yaml
var ConfigYaml []byte

type ConfigStruct struct {
	Port      int       `json:"port" yaml:"port"`
	DbType    string    `json:"dbType" yaml:"dbType"`
	Mysql     Mysql     `yaml:"mysql"`
	OpenGauss OpenGauss `yaml:"opengauss"`
	Logger    Logger    `yaml:"logger"`
}

type Mysql struct {
	Path         string `json:"path,omitempty" yaml:"path"`
	Dbname       string `json:"dbname,omitempty" yaml:"dbname"`
	Username     string `json:"username,omitempty" yaml:"username"`
	Password     string `json:"password,omitempty" yaml:"password"`
	MaxIdleConns int    `json:"maxIdleConns,omitempty" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns,omitempty" yaml:"maxOpenConns"`
	LogLevel     int    `json:"logLevel,omitempty" yaml:"logLevel"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local&readTimeout=10s&writeTimeout=20s&parseTime=true", m.Username, m.Password, m.Path, m.Dbname)
}

type OpenGauss struct {
	Path         string `json:"path,omitempty" yaml:"path"`
	Dbname       string `json:"dbname,omitempty" yaml:"dbname"`
	Username     string `json:"username,omitempty" yaml:"username"`
	Password     string `json:"password,omitempty" yaml:"password"`
	MaxIdleConns int    `json:"maxIdleConns,omitempty" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns,omitempty" yaml:"maxOpenConns"`
	LogLevel     int    `json:"logLevel,omitempty" yaml:"logLevel"`
}

func (o *OpenGauss) Dsn(schema string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?search_path=%s", o.Username, o.Password, o.Path, o.Dbname, schema)
}

type Logger struct {
	LogLevel      string
	LogDirPath    string
	LogFileName   string
	LogMaxSize    int64
	LogStorageDay int64
}
