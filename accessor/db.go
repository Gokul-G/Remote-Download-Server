package accessor

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Mysql Driver
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

//Datastore contanins the database Connection Object
type Datastore struct {
	db *sql.DB
}

var DS *Datastore

//InitDB initialies the DB Connection at start
func InitDB() error {

	db, err := sql.Open("mysql", "root@/remote_download?charset=utf8&parseTime=true&multiStatements=true")
	handleError(err)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	handleError(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	handleError(err)
	m.Steps(2)

	DS = &Datastore{db}
	return nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
