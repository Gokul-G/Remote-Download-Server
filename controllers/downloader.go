package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gokul-G/Remote-Download-Server/accessor"

	"github.com/Gokul-G/Remote-Download-Server/models"
)

//Private Methods
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

func download(item *models.DownloadItem) (err error) {

	// Create the file
	out, err := os.Create(item.Name)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(item.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//DownloadProcess Started
	accessor.UpdateDownloadStatus(item, models.InProgress)

	// Writer the body to file
	src := &PassThru{Reader: resp.Body, total: float64(resp.ContentLength)}
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}

	//DownloadProcess Ended
	accessor.UpdateDownloadStatus(item, models.Completed)
	print("\n Completed =>" + item.Name)

	return nil
}
