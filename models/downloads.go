package models

import (
	"time"
)

// DownloadStatus - Represents the Download Status of Each Download Process
type DownloadStatus int

const (
	NotStarted DownloadStatus = iota
	InProgress
	Paused
	Completed
)

//DownloadItem - Represents the each Download Item
type DownloadItem struct {
	ID        int64          `json:"id" form:"-"`
	Name      string         `json:"name" form:"name"`
	URL       string         `json:"url"`
	Status    DownloadStatus `json:"status"`
	Size      int64          `json:"size"`
	CreatedAt time.Time      `json:"created_at"`
}

//DownloadItems - Array of Download Item
type DownloadItems []DownloadItem
