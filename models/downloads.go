package models

import "time"

type DownloadStatus int

const (
	NotStarted DownloadStatus = iota
	InProgress
	Paused
	Completed
)

//Downloads Object
type Download struct {
	ID         int            `json:"id" form:"-"`
	Name       string         `json:"name" form:"name"`
	URL        string         `json:"url"`
	Status     DownloadStatus `json:"status"`
	Created_at time.Time      `json:"created_at"`
}

type Downloads []Download

//Data For Download
type DownloadData struct {
	URL string `json:"url"`
}
