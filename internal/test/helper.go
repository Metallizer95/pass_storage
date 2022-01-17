package test

import (
	"context"
	"fmt"
	"store_server/internal/storage/mongorepo"
	"testing"
)

func TestDatabase(t *testing.T) (mongorepo.Client, func()) {
	fmt.Println("In function TestDatabase")
	client, err := newClient()
	fmt.Println("Client is created")
	if err != nil {
		t.Fatal(err)
	}
	return client, func() {
		_ = client.db.Database(dbName).Drop(context.TODO())
	}
}
