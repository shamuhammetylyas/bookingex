package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB struct etmegimizin sebabi projemizde birden kop database ulanjak bolsak gerek bolyar
// birden kop driver ulanjak bolsak sheytsek gowy bolyar
// bashlangyc ucin biz postgres driver ulanyarys, yone son bashga database
// driver ulanjak bolsak gowy
// DB type-in 1-nji memberi postgres driver
// 2-nji memberi bashga database driver edip ulanyp bolyar
// shol driverlerin ikisem sql.DB pointeri ulanyar
type DB struct {
	SQL *sql.DB
}

// empty DB doredyar
var dbConn = &DB{}

// bir database connectionly projeler ucin ashakdaky constantlar gerek dal
// yone bir projede kop database ulanjak bolsak onda database pool conf ucin ashakdakylar gerek bolar
// maxOpenDbConn -> nace db connect edip bolyar
//
const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

//Creates database pool for postgres
//ConnectSQL database pool doredyar. Database havuzu manysyny beryar
//birden kop database ucin gerek bolyar
func ConnectSQL(dsn string) (*DB, error) {
	// NewDatabase taze bir database connection acyar
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	// NewDatabase funksiyasyndan gelyan database connection pooly dbConn structyn SQL-ne beryar
	dbConn.SQL = d
	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// tries to ping to database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

// creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
