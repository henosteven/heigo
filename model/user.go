package model

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/henosteven/heigo/config"
	"github.com/henosteven/heigo/lib"
	"fmt"
	"strconv"
)

/*
CREATE TABLE `user` (
  `UserID` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8
*/

var db *sql.DB

func InitDb(config config.MysqlConfig) {
	var err error
	//root:@tcp(127.0.0.1:3306)/test
	db, err = sql.Open("mysql", fmt.Sprintf("%s:@%s(%s:%s)/%s", config.User, config.Protocol, config.Host, config.Port, config.Database))
	if err != nil {
		panic(err.Error())
	}
}

func TeardownDb() {
	db.Close()
}

func AddUser(userName string) (int, error){
	if len(userName) == 0 {
		return 0, errors.New("user name empty")
	}

	stmtOut, err := db.Prepare("INSERT INTO User VALUES (NULL, ?)") // ? = placeholder
	if err != nil {
		return 0, err
	}
	defer stmtOut.Close()

	result, err := stmtOut.Exec(userName)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	if userID == 0 {
		return 0, err
	}

	return int(userID), nil
}

func GetUserNameByID(userID int) (string, error){
	if userID <= 0 {
		return "", errors.New("invalid userid")
	}

	tmpUserName, err := lib.Get(getUserCacheKey(userID))
	if tmpUserName != "" {
		return tmpUserName, err
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

	lib.Set(getUserCacheKey(userID), name)
	return name, nil
}

func getUserCacheKey(userID int) string {
	return "user_" + strconv.Itoa(userID)
}