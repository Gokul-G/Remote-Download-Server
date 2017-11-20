package accessor

import (
	"database/sql"
	"fmt"

	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

type Datastore struct {
	db *sql.DB
}

var DS *Datastore

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

func GetDownloadListFromDB() models.Downloads {
	var downloads models.Downloads
	var download models.Download
	rows, err := DS.db.Query("select * from downloads;")
	for rows.Next() {
		err = rows.Scan(&download.ID, &download.Name, &download.URL, &download.Status, &download.Created_at)
		downloads = append(downloads, download)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return downloads
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
