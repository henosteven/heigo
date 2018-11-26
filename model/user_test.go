package model

import (
	"github.com/henosteven/heigo/config"
	"github.com/henosteven/heigo/lib"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var mysqlConfig = config.MysqlConfig{
		"127.0.0.1",
		"3306",
		"root",
		"",
		"test",
		"tcp",
	}
	InitDb(mysqlConfig)

	var redisConfig = config.RedisConfig{
		3,
		100,
		"127.0.0.1",
		"6379",
		"",
		"",
	}
	lib.InitRedis(redisConfig)
	retCode := m.Run()
	TeardownDb()
	os.Exit(retCode)
}

func TestAddUser(t *testing.T) {
	caseList := []struct {
		userName     string
		expectResult bool
	}{
		{"heno", true},
		{"jinjing", true},
	}
	for _, val := range caseList {
		result, err := AddUser(val.userName)
		if err != nil {
			t.Errorf("addUser failed, error: %s, name: %s", val.userName, err.Error())
		}

		if result <= 0 {
			t.Errorf("addUser failed, name: %s, expect: %v  get: %v", val.userName, val.expectResult, result)
		}
	}
}

func TestGetUserNameByID(t *testing.T) {
	caseList := []struct {
		userID     int
		expectName string
	}{
		{1, "heno"},
	}
	for _, val := range caseList {
		name, err := GetUserNameByID(val.userID)
		if err != nil {
			t.Errorf("getUserNameByID failed, error: %s", err.Error())
		}

		if name != val.expectName {
			t.Errorf("getUserNameByID failed, expect: %s  get: %s", val.expectName, name)
		}
	}
}
