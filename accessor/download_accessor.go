package accessor

import (
	"fmt"

	"github.com/Gokul-G/Remote-Download-Server/models"
)

// GetDownloadListFromDB - retrives download list
func GetDownloadListFromDB() models.DownloadItems {
	var downloadItems models.DownloadItems
	var downloadItem models.DownloadItem
	rows, err := DS.db.Query("select * from downloads;")
	for rows.Next() {
		err = rows.Scan(&downloadItem.ID, &downloadItem.Name, &downloadItem.URL, &downloadItem.Status, &downloadItem.Created_at)
		downloadItems = append(downloadItems, downloadItem)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return downloadItems
}

// CreateDownload - adds a download item to DB
func CreateDownload(downloadObj *models.DownloadItem) {
	stmt, err := DS.db.Prepare("INSERT downloads SET name = ?, url = ?, status = ?")
	handleError(err)

	data, err := stmt.Exec(downloadObj.Name, downloadObj.URL, downloadObj.Status)
	id, _ := data.LastInsertId()
	downloadObj.ID = id
	handleError(err)
}

// UpdateDownloadStatus - updates the status of the downloadItem
func UpdateDownloadStatus(downloadObj *models.DownloadItem, status models.DownloadStatus) {
	stmt, err := DS.db.Prepare("Update downloads SET status=? WHERE id = ?")
	handleError(err)

	_, err = stmt.Exec(status, downloadObj.ID)
	handleError(err)
}
