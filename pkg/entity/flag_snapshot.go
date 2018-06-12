//go:generate goqueryset -in flag_snapshot.go

package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// FlagSnapshot is the snapshot of a flag
// Any change of the flag will create a new snapshot
// gen:qs
type FlagSnapshot struct {
	gorm.Model
	FlagID    uint `gorm:"index:idx_flagsnapshot_flagid"`
	UpdatedBy string
	FlagJSON  string `sql:"type:text"` // JSON string repr of flag
	Flag      *Flag  `gorm:"-"`
}

// Scan implements scanner interface
func (f *Flag) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s := cast.ToString(value)
	if err := json.Unmarshal([]byte(s), f); err != nil {
		return fmt.Errorf("cannot scan %v into Flag type. err: %v", value, err)
	}
	return nil
}

// Value implements valuer interface
func (f *Flag) Value() (driver.Value, error) {
	bytes, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}

// BeforeSave converts FlagSnapshot's Flag field into FlagJSON
func (fs *FlagSnapshot) BeforeSave() (err error) {
	flagJSON, err := fs.Flag.Value()
	if err != nil {
		return err
	}
	fs.FlagJSON = flagJSON.(string)
	return nil
}

// AfterFind converts FlagSnapshot's FlagJSON field into Flag object
func (fs *FlagSnapshot) AfterFind() (err error) {
	fs.Flag = &Flag{}
	return fs.Flag.Scan(fs.FlagJSON)
}

// SaveFlagSnapshot saves the Flag Snapshot
func SaveFlagSnapshot(db *gorm.DB, flagID uint, updatedBy string) {
	tx := db.Begin()

	f := &Flag{}
	q := NewFlagQuerySet(tx).IDEq(flagID)
	if err := q.One(f); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"flagID": flagID,
		}).Error("failed to find the flag when SaveFlagSnapshot")
		return
	}
	f.Preload(tx)

	fs := FlagSnapshot{FlagID: f.ID, UpdatedBy: updatedBy, Flag: f}
	if err := tx.Create(&fs).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"err":    err,
			"flagID": f.ID,
		}).Error("failed to save FlagSnapshot")
		tx.Rollback()
		return
	}

	if err := q.GetUpdater().SetUpdatedBy(updatedBy).SetSnapshotID(fs.ID).Update(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":            err,
			"flagID":         f.ID,
			"flagSnapshotID": fs.ID,
		}).Error("failed to save Flag's UpdatedBy and SnapshotID")
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
	}
}
