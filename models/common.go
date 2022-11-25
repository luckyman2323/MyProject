package models

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormMysql() *gorm.DB {
	dbuser := beego.AppConfig.String("Dbuser")
	dbpasswd := beego.AppConfig.String("Dbpasswd")
	dbip := beego.AppConfig.String("Dbip")
	dbport := beego.AppConfig.String("Dbport")
	dbname := beego.AppConfig.String("Dbname")

	dbKey := strings.Join(MYSQLKEYS, "")
	// 解密dbuser和dbpasswd
	dbpasswd, err := AesDecrypt(dbpasswd, []byte(dbKey))
	if err != nil {
		panic(fmt.Sprintf("Error AesDecrypt dbpasswd, err: %s", err))
	}
	dbuser, err = AesDecrypt(dbuser, []byte(dbKey))
	if err != nil {
		panic(fmt.Sprintf("Error AesDecrypt dbuser, err: %s", err))
	}
	connectstr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=Local&readTimeout=%s&writeTimeout=%s&parseTime=true", dbuser, dbpasswd, net.JoinHostPort(dbip, dbport), dbname, "utf8", "10s", "20s")

	if db, err := gorm.Open(mysql.Open(connectstr), gormConfig()); err != nil {
		panic(err.Error())
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(20)
		return db
	}
}

func gormConfig() *gorm.Config {
	var config = &gorm.Config{
		SkipDefaultTransaction:                   true, // 为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它
		DisableForeignKeyConstraintWhenMigrating: true, // 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
	}

	// 根据beego的日志级别设置gorm的日志级别，由于gorm只支持Info、Warn、Error、Silent四种日志级别，所以将beego的日志级别从7到0依次据范围设置gorm的日志级别
	switch logLevel, _ := beego.AppConfig.Int("logs::orm_log_level"); logLevel {
	case 1:
		config.Logger = logger.Default.LogMode(logger.Info)
	case 2:
		config.Logger = logger.Default.LogMode(logger.Warn)
	case 3:
		config.Logger = logger.Default.LogMode(logger.Error)
	case 4:
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return config
}

var MYSQLKEYS = []string{"sfe0", "23f_", "9fd&", "fwfl"}

// 解密
func AesDecrypt(crypted string, key []byte) (string, error) {
	if crypted == "" {
		return "", nil
	}
	decodeString, err := hex.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	return AESCBCPKCS7Decrypt(string(decodeString), key)
}

// AESCBCPKCS7Decrypt combines CBC decryption and PKCS7 unpadding
func AESCBCPKCS7Decrypt(src string, key []byte) (string, error) {
	// First decrypt
	pt, err := aesCBCDecrypt(key, []byte(src))
	if err == nil {
		res, err1 := pkcs7UnPadding(pt)
		return string(res), err1
	}
	return "", err
}

func aesCBCDecrypt(key, src []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(src) < aes.BlockSize {
		return nil, errors.New("invalid ciphertext. It must be a multiple of the block size")
	}
	iv := src[:aes.BlockSize]
	src = src[aes.BlockSize:]

	if len(src)%aes.BlockSize != 0 {
		return nil, errors.New("invalid ciphertext. It must be a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(src, src)

	return src, nil
}

func pkcs7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > aes.BlockSize || unpadding == 0 {
		return nil, errors.New("invalid pkcs7 padding (unpadding > aes.BlockSize || unpadding == 0)")
	}

	pad := src[len(src)-unpadding:]
	for i := 0; i < unpadding; i++ {
		if pad[i] != byte(unpadding) {
			return nil, errors.New("invalid pkcs7 padding (pad[i] != unpadding)")
		}
	}

	return src[:(length - unpadding)], nil
}

// 加密
func AesEncrypt(originData, key []byte) (string, error) {
	crypted, err := AESCBCPKCS7Encrypt(string(originData), key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(crypted), nil
}

// AESCBCPKCS7Encrypt combines CBC encryption and PKCS7 padding
func AESCBCPKCS7Encrypt(src string, key []byte) ([]byte, error) {
	// First pad
	tmp := pkcs7Padding([]byte(src))

	// Then encrypt
	return aesCBCEncrypt(key, tmp)
}

func pkcs7Padding(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func aesCBCEncrypt(key, s []byte) ([]byte, error) {
	return aesCBCEncryptWithRand(rand.Reader, key, s)
}

func aesCBCEncryptWithRand(prng io.Reader, key, s []byte) ([]byte, error) {
	if len(s)%aes.BlockSize != 0 {
		return nil, errors.New("invalid plaintext. It must be a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(s))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(prng, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], s)

	return ciphertext, nil
}
