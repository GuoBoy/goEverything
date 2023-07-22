package main

import "time"

const (
	IS_FILE = iota // file type
	IS_DIR         // dir type
)

// Item name file/dir item
type Item struct {
	ID     string `json:"id" gorm:"id"`
	Path   string `json:"path" gorm:"path"`
	Name   string `json:"name" gorm:"name"`
	Size   int64  `json:"size" gorm:"size"`
	Type   uint   `json:"type" gorm:"type"`
	DiskID string `json:"disk_id" gorm:"disk_id"`
}

// Disk disk info
type Disk struct {
	ID        string    `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type DiskInfo struct {
	ID        string    `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	FileCount int64     `json:"file_count"`
	DirCount  int64     `json:"dir_count"`
}
