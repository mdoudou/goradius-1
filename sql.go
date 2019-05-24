package goradius

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type DbConfig struct {
	SqlType string
	SqlCmd  string
}

var a, b = GetDbConfig()

//判断是否正确连接数据库
func DbLive() {

	db, err := sql.Open(a, b)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	err1 := db.Ping()
	if err1 != nil {
		log.Fatal("Wrong connect to DB")

	}
}

//由客户端IP得到客户端对应得密钥
func GetSecretAndDiff(ipp string) (secret []byte) {
	db, _ := sql.Open(a, b)
	defer db.Close()
	res := db.QueryRow("select IP,secret from nas where IP = ?", ipp)
	err2 := res.Scan(&ipp, &secret)
	if err2 != nil {
		log.Println(err2)
	}
	if string(secret) == "" || string(ipp) == "" {
		log.Println("请把IP和密钥加入认证服务器白名单")
		return nil
	} else {
		secret = []byte(secret)
		return secret
	}

}

//从数据库中读取用户名和密码
func GetUserPasswd(u string) (p string) {
	db, _ := sql.Open(a, b)
	defer db.Close()
	var password string
	err2 := db.QueryRow("select password from users where Name = ?", u).Scan(&password)
	if err2 != nil {
		log.Println(err2)
		return ""
	} else {
		return password
	}

}
