package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clientID := 1
	client := Client{}
	row := db.QueryRow("SELECT * FROM clients WHERE id = :id", sql.Named("id", clientID))
	err = row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)
	if err != nil {
		fmt.Println(err)
        return
	}

	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.FIO, client.Login, client.Birthday, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")// настройте подключение к БД
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// clientID := -1

	// напиши тест здесь
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")// настройте подключение к БД
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// cl := Client{
	// 	FIO:      "Test",
	// 	Login:    "Test",
	// 	Birthday: "19700101",
	// 	Email:    "mail@mail.com",
	// }

	// напиши тест здесь
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")// настройте подключение к БД
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// cl := Client{
	// 	FIO:      "Test",
	// 	Login:    "Test",
	// 	Birthday: "19700101",
	// 	Email:    "mail@mail.com",
	// }

	// напиши тест здесь
}
