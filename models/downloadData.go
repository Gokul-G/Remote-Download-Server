package models

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//DownloadData - Data Model for Creating DownloadItem
type DownloadData struct {
	URL string `json:"url"`
}

// ToDownloadItem - Returns DownloadItem for the DownloadData
func (data DownloadData) ToDownloadItem() DownloadItem {

	tokens := strings.Split(data.URL, "/")
	fileName := tokens[len(tokens)-1]
	localFileName := fileName
	var i = 0
	for {
		if i > 0 {
			extension := filepath.Ext(fileName)
			name := strings.TrimSuffix(fileName, extension)
			localFileName = name + "_" + strconv.Itoa(i) + extension
		}

		if _, err := os.Stat(localFileName); os.IsNotExist(err) {
			fmt.Println("File does not exist")
			break
		} else {
			fmt.Println("File exists")
		}
		i++
	}

	var downloadItem DownloadItem
	downloadItem.Name, _ = url.QueryUnescape(localFileName)
	downloadItem.URL = data.URL
	downloadItem.Status = NotStarted
	return downloadItem
}
