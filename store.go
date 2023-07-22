package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB

func InitDb() {
	dB, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Fatal(err)
	}
	db = dB
	db.AutoMigrate(&Disk{}, &Item{})
}

type Store interface {
	Add() error
	Delete()
	Find(key string) any
}

type DiskStore struct {
	D *Disk
}

func (d *DiskStore) Add() error {
	var disk *Disk
	if err := db.Find(&disk, "id=?", d.D.ID).Error; err != nil {
		return err
	}
	if disk.ID == "" {
		d.D.UpdatedAt = time.Now()
		return db.Create(&d.D).Error
	} else {
		return db.Model(&Disk{}).Where("id=?", d.D.ID).Update("updated_at", time.Now()).Error
	}
}

func (d *DiskStore) Delete() {

}

func (d *DiskStore) Find() ([]*DiskInfo, error) {
	var (
		disks []*Disk
		res   []*DiskInfo
	)

	if err := db.Find(&disks).Error; err != nil {
		return nil, err
	}
	for _, disk := range disks {
		var (
			fileCount int64
			dirCount  int64
			diskInfo  = DiskInfo{}
		)
		db.Model(&Item{}).Where("disk_id = ? and type = ?", disk.ID, IS_FILE).Count(&fileCount)
		db.Model(&Item{}).Where("disk_id = ? and type = ?", disk.ID, IS_DIR).Count(&dirCount)
		diskInfo.ID, diskInfo.Name, diskInfo.UpdatedAt, diskInfo.FileCount, diskInfo.DirCount = disk.ID, disk.Name, disk.UpdatedAt, fileCount, dirCount
		res = append(res, &diskInfo)
	}
	return res, nil
}

// ItemStore //
type ItemStore struct {
	D *Item
}

func (i *ItemStore) Add(items []*Item) {
	for _, itm := range items {
		db.Create(&itm)
	}
}

func (i *ItemStore) Delete(id string) error {
	return db.Delete(&Item{}, "disk_id = ?", id).Error
}

func (i *ItemStore) Find(key string) ([]*Item, error) {
	var res []*Item
	if err := db.Where("path like ?", "%"+key+"%").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
