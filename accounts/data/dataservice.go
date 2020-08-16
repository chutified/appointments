package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chutified/appointments/accounts/config"
	"github.com/pkg/errors"
)

// databaseService is the
type databaseService struct{ db *sql.DB }

// New is the contructor for the Database.
// func New() *Database {
//     return &databaseService{}
// }

// Init initialize the databaseService connection to the database.
func (ds *databaseService) Init(cfg *config.DBConfig) error {

	// define database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBname)

	// open database connection
	var err error
	ds.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return errors.Wrap(err, "connecting to the db")
	}

	// test the connection
	for i := 0; i < 3; i++ {
		err = ds.db.Ping()
		if err == nil {
			break
		}

		time.Sleep(3 * time.Second)
	}
	if err != nil {
		return errors.Wrap(err, "db connection failed")
	}

	return nil
}