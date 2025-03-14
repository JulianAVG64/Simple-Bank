package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/JulianAVG64/Simple-Bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
