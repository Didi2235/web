package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"../db"
	_ "github.com/go-sql-driver/mysql"
)

type Logging struct {
	Response int
	Log      struct {
		Status  int
		Message string `json:"msg"`
		Account struct {
			User_id  int    `json:"user_id" `
			Username string `json:"username"`
			Session  string `json:"session_id"`
		} `json:"account"`
	}
}

func Login(email string, pass string) Logging {
	con := db.CreateCon()
	var data Logging
	hasher := md5.New()

	sqlStatement := "SELECT `user_id`, `username` FROM `user` WHERE Email = ? AND Password = ?"
	err := con.QueryRow(sqlStatement, email, pass).Scan(&data.Log.Account.User_id, &data.Log.Account.Username)
	if err != nil {
		fmt.Println(err)
		data.Log.Status = 01
		data.Log.Message = "User Not EXISTS"
		return data
	}
	hasher.Write([]byte(data.Log.Account.Username))
	data.Log.Account.Session = hex.EncodeToString(hasher.Sum(nil))
	sqlStatement = "UPDATE `user` SET `session`= ? WHERE `Email` = ? AND Password = ?"
	_, err = con.Query(sqlStatement, data.Log.Account.Session, email, pass)
	if err != nil {
		fmt.Println(err)
		data.Response = http.StatusBadRequest
		data.Log.Account.User_id = 0
		data.Log.Account.Username = ""

		return data
	}
	data.Response = http.StatusOK
	data.Log.Status = 00
	data.Log.Message = "LOGIN Success"
	return data
}
