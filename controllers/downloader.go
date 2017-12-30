package controllers

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Gokul-G/Remote-Download-Server/models"

	"github.com/Gokul-G/Remote-Download-Server/accessor"
	_ "github.com/Gokul-G/Remote-Download-Server/models"
)

var bytesToMegaBytes = int64(1048576)

type PassThru struct {
	io.Reader
	curr         int64
	downloadItem models.DownloadItem
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	if err == nil || (err == io.EOF && n > 0) {
		printProgress(int64(pt.curr), pt.downloadItem.Size)
		var data = accessor.BroadcastMessage{DownloadIem: pt.downloadItem, DownloadedSize: pt.curr}
		accessor.BroadcastChannel <- data
	}

	return n, err
}

func printProgress(curr, total int64) {
	width := 40.0
	output := ""
	threshold := float64(curr/total) * float64(width)
	for i := 0.0; i < width; i++ {
		if i < threshold {
			output += "="
		} else {
			output += " "
		}
	}

	fmt.Printf("\r[%s] %.1f of %.1fMB", output, float64(curr/bytesToMegaBytes), float64(total/bytesToMegaBytes))
}

func download(data *models.DownloadData) (err error) {

	var item = data.ToDownloadItem()

	// Get the data
	resp, err := http.Get(data.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Get the Download File Name
	contentDisposition := resp.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	downloadFileName := params["filename"]

	item.Name = createNonExistingFileName(downloadFileName)
	item.URL = data.URL
	item.Size = resp.ContentLength

	// Create the file
	out, err := os.Create(item.Name)
	if err != nil {
		return err
	}
	defer out.Close()

	//DownloadProcess Started
	item.Status = models.InProgress
	accessor.CreateDownload(&item)

	// Writer the body to file
	src := &PassThru{Reader: resp.Body, downloadItem: item}
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	//DownloadProcess Ended
	accessor.UpdateDownloadStatus(&item, models.Completed)

	print("\n Completed =>" + item.Name)
	return nil
}

func createNonExistingFileName(downloadFileName string) string {
	localFileName := downloadFileName
	var i = 0
	for {
		if i > 0 {
			extension := filepath.Ext(downloadFileName)
			name := strings.TrimSuffix(downloadFileName, extension)
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

	itemName, _ := url.QueryUnescape(localFileName)
	return itemName
}
