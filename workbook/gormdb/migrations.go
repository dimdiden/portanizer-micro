package gormdb

import (
	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

func RunMigration(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "add Tag migration",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&workbook.Tag{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(&workbook.Tag{}).Error
			},
		},
		{
			ID: "add Post migration",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&workbook.Post{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(&workbook.Post{}).Error
			},
		},
	})
	return m.Migrate()
}
