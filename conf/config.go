package conf

import (
	"fmt"
	"os"
	"path"

	"myproject/utils"

	"github.com/spf13/viper"
)

var Config ConfigStruct

var configDir = "./conf/config.yaml"

func initConfig() {
	var err error

	_, err = os.Lstat(configDir)
	if os.IsNotExist(err) {
		dir, _ := path.Split(configDir)
		if err = os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
		if err = os.WriteFile(configDir, ConfigYaml, 0777); err != nil {
			panic(err)
		}

	}

	viper.SetConfigFile(configDir)
	//读取配置文件内容
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 如果是.yaml的配置文件 可以通过字段绑定env

	// viper.AutomaticEnv()
	// viper.BindEnv("port", "SYSTEM_PORT")
	// viper.BindEnv("dbType", "SYSTEM_DBTYPE")
	// viper.BindEnv("logger.logstorageday", "LOGGER_STORAGEDAY")
	// viper.BindEnv("logger.logmaxsize", "LOGGER_MAXSIZE")
	// viper.BindEnv("logger.logfilename", "LOGGER_FILENAME")
	// viper.BindEnv("logger.logdirpath", "LOGGER_DIRPATH")
	// viper.BindEnv("logger.loglevel", "LOGGER_LEVEL")
	// viper.BindEnv("rollingtime", "ROLLINGTIME")
	// viper.BindEnv("certPath", "CERTPATH")
	// viper.BindEnv("opengauss.logLevel", "OPENGAUSS_LOGLEVEL")
	// viper.BindEnv("opengauss.connMaxLifetime", "OPENGAUSS_CONNMAXLIFETIME")
	// viper.BindEnv("opengauss.maxOpenConns", "OPENGAUSS_MAXOPENCONNS")
	// viper.BindEnv("opengauss.maxIdleConns", "OPENGAUSS_MAXIDLECONNS")
	// viper.BindEnv("opengauss.dbname", "OPENGAUSS_DBNAME")
	// viper.BindEnv("opengauss.path", "OPENGAUSS_PATH")
	// viper.BindEnv("opengauss.password", "OPENGAUSS_PASSWORD")
	// viper.BindEnv("opengauss.username", "OPENGAUSS_USERNAME")

	if err = viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	// 解码
	switch Config.DbType {
	case "mysql":
		Config.Mysql.Username, err = utils.AesDecrypt(Config.Mysql.Username, utils.BASE64)
		if err != nil {
			panic(fmt.Errorf("error AesDecrypt mysql.username, err: %s", err))
		}
		Config.Mysql.Password, err = utils.AesDecrypt(Config.Mysql.Password, utils.BASE64)
		if err != nil {
			panic(fmt.Errorf("error AesDecrypt mysql.password, err: %s", err))
		}
	case "opengauss":
		Config.OpenGauss.Username, err = utils.AesDecrypt(Config.OpenGauss.Username, utils.BASE64)
		if err != nil {
			panic(fmt.Errorf("error AesDecrypt opengauss.username, err: %s", err))
		}
		Config.OpenGauss.Password, err = utils.AesDecrypt(Config.OpenGauss.Password, utils.BASE64)
		if err != nil {
			panic(fmt.Errorf("error AesDecrypt opengauss.password, err: %s", err))
		}
	}

	fmt.Printf("Config: %+v \n", Config)
}
