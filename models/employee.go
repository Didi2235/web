package models

import (
	"fmt"

	"net/http"

	"../db"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Firstname string `json:"first" form:"first"`
	Lastname  string `json:"last" form:"last"`
	User      string `json:"user" form:"user" query:"user"`
	Email     string `json:"email" form:"email"`
	Pass      string `json:"pass" form:"pass" query:"pass"`
}

type Employees struct {
	Data struct {
		Account interface {
		} `json:"member"`
		Status string
	}
}
type Status struct {
	Code   int
	Status struct {
		ResultCode int
		Info       string
	}
}

var total int

func Regist(Fname, Lname, User, Email, Pass string) Status {
	con := db.CreateCon()
	result := Status{}
	if Fname == "" || Lname == "" || User == "" || Email == "" || Pass == "" {
		result.Code = http.StatusBadRequest
		result.Status.ResultCode = 01
		result.Status.Info = "Value Error"
		return result
	}
	sqlStatement := "SELECT COUNT(*) FROM `user` where `Username` = ? "
	rows, _ := con.Query(sqlStatement, User)
	for rows.Next() {
		rows.Scan(&total)
		fmt.Println(total)
	}
	if total != 0 {
		result.Code = http.StatusInternalServerError
		result.Status.ResultCode = 03
		result.Status.Info = "User has exists"
		return result
	}

	sqlStatement = "INSERT INTO user VALUES (NULL, ?, ?, ?, ?, ?, NULL)"
	rows, err := con.Query(sqlStatement, Fname, Lname, User, Email, Pass)

	defer rows.Close()

	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Status.ResultCode = 02
		result.Status.Info = "FALSE"
	} else {
		result.Code = http.StatusOK
		result.Status.ResultCode = 00
		result.Status.Info = "Success Regist"
	}
	return result
}

func GetEmployee() Employees {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT*FROM user"
	employee := Employee{}
	var empty string

	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	con.Close()
	var test []interface{}
	for rows.Next() {

		err := rows.Scan(&empty, &employee.Firstname, &employee.Lastname, &employee.User, &employee.Email, &employee.Pass)
		if err != nil {
			fmt.Println(err)
			break
		}
		test = append(test, employee)
	}
	fmt.Println(test)
	result := Employees{}
	result.Data.Account = test
	result.Data.Status = "true"

	return result
}
