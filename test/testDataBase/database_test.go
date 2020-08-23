package main

import (
	"testing"

	"firego/src/db"
)

func TestDataBase(t *testing.T) {
	client := db.GetInstance()

	key := "123"
	value := "hahaha"

	client.PutByKey(db.ChatRoomPrefix + key, value)

	var value2 string
	client.GetByKey(db.ChatRoomPrefix + key, &value2)


	t.Logf("Get value = %s", value2)
}
