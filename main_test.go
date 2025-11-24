package main

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД

	clientID := 1

	require.NotNil(t, clientID)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД

	clientID := -1

	require.GreaterOrEqual(t, clientID, 1)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err)
	res, err := db.Exec("INSERT INTO clients (FIO, Login, Birthday, Email) VALUES (:FIO, :Login, :Birthday, :Email)", sql.Named("FIO", cl.FIO), sql.Named("Login", cl.Login), sql.Named("Birthday", cl.Birthday), sql.Named("Email", cl.Email))
	require.Nil(t, err)
	id, err := res.LastInsertId()
	require.Nil(t, err)
	require.NotNil(t, id)
	require.NotNil(t, cl.ID)
	// напиши тест здесь

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
}
