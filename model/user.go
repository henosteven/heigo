package model

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE `user` (
  `UserID` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8
*/

var db *sql.DB

func InitDb() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}

func TeardownDb() {
	db.Close()
}

func AddUser(userName string) (bool, error){
	if len(userName) == 0 {
		return false, errors.New("user name empty")
	}

	stmtOut, err := db.Prepare("INSERT INTO User VALUES (NULL, ?)") // ? = placeholder
	if err != nil {
		return false, err
	}
	defer stmtOut.Close()

	result, err := stmtOut.Exec(userName)
	if err != nil {
		return false, err
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if affectedRow == 0 {
		return false, err
	}

	return true, nil
}

func GetUserNameByID(userID int) (string, error){
	if userID <= 0 {
		return "", errors.New("invalid userid")
	}

	stmtOut, err := db.Prepare("select UserName from User WHERE UserID=?") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(userID)
	if err != nil {
		panic(err.Error())
	}

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			panic(err.Error())
		}
	}
	return name, nil
}