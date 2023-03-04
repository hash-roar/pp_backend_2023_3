package models

import (
	"pp-bakcend/pkg/logging"

	"github.com/jinzhu/gorm"
)

type Equipment struct {
	gorm.Model
	Hostname string
	Version  string
	Jct      string
}

func GetEquipmentByHostname(hostname string) (*Equipment, error) {
	var equipment Equipment
	result := db.Where("hostname = ?", hostname).First(&equipment)
	logging.Trace(equipment)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return nil, result.Error
		} else {
			return nil, nil
		}
	}
	return &equipment, nil

}

func CreateEquipment(e *Equipment) error {
	logging.Trace(e)
	return db.Create(e).Error
}

func UpdateEquipmentJct(hostname string, newJct string) error {
	return db.Model(&Equipment{}).Where("hostname = ?", hostname).UpdateColumn("jct", newJct).Error
}

func GetAllEquipmentsByHostname(Hostnames []string) ([]Equipment, error) {
	var equipments []Equipment
	result := db.Where("hostname IN ?", Hostnames).Find(&equipments)
	return equipments, result.Error
}

func GetUserBlockWords(mid string) ([]BlockWords, error) {
	var words []BlockWords
	result := db.Model(&BlockWords{}).Where("mid = ?", mid).Find(&words)
	return words, result.Error
}
