package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Gokul-G/Remote-Download-Server/accessor"
	"github.com/Gokul-G/Remote-Download-Server/models"
)

func GetDownloadList(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if r.Method == "OPTIONS" {
		return
	}

	downloads := accessor.GetDownloadListFromDB()
	response, _ := json.Marshal(downloads)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func StartDownload(w http.ResponseWriter, r *http.Request) {

	var downloadData models.DownloadData
	json.NewDecoder(r.Body).Decode(&downloadData)
	go downloadFromURL(downloadData.URL)
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(downloadData)

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

var bytesToMegaBytes = 1048576.0

type PassThru struct {
	io.Reader
	curr  int64
	total float64
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	// last read will have EOF err
	if err == nil || (err == io.EOF && n > 0) {
		printProgress(float64(pt.curr), pt.total)
	}

	return n, err
}

func printProgress(curr, total float64) {
	width := 40.0
	output := ""
	threshold := (curr / total) * float64(width)
	for i := 0.0; i < width; i++ {
		if i < threshold {
			output += "="
		} else {
			output += " "
		}
	}

	fmt.Printf("\r[%s] %.1f of %.1fMB", output, curr/bytesToMegaBytes, total/bytesToMegaBytes)
}

func downloadFromURL(url string) (err error) {

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	localFileName := fileName
	var i = 0
	for {

		if i > 0 {
			extension := filepath.Ext(fileName)
			name := strings.TrimSuffix(fileName, extension)
			localFileName = name + "_" + strconv.Itoa(i) + extension
		}

		fmt.Println(localFileName)
		if _, err := os.Stat(localFileName); os.IsNotExist(err) {
			fmt.Println("File does not exist")
			break
		} else {
			fmt.Println("File exists")
		}
		i++
	}

	// Create the file
	out, err := os.Create(localFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	src := &PassThru{Reader: resp.Body, total: float64(resp.ContentLength)}

	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	return nil
}
