package authservice

import (
	"fmt"
	"pp-bakcend/models"
	"pp-bakcend/pkg/gredis"
	"pp-bakcend/pkg/logging"
	"pp-bakcend/utils"
)

func CheckAuth(mid string, hostname string, jct string, encryptStr string, version string) bool {

	// auth fail
	if !utils.IsEqual(encryptStr, fmt.Sprintf("%s.%s.%s", mid, hostname, jct)) {
		return false
	}

	// update jct or add equipments
	equipment, err := models.GetEquipmentByHostnameMid(mid, hostname)
	if err != nil {
		logging.Error(err)
		return false
	}
	err = models.AddUserLoginTimes(mid, 1)
	if err != nil {
		logging.Error(err)
		return false
	}
	if equipment == nil {
		err = models.CreateEquipment(&models.Equipment{
			Mid:      mid,
			Hostname: hostname,
			Version:  version,
			Jct:      jct,
		})
		if err != nil {
			logging.Error(err)
			return false
		}
		return true
	}

	// update
	err = models.UpdateEquipmentJctVersion(hostname,mid, jct, version)
	if err != nil {
		logging.Error(err)
		return false
	}

	return true

}

func AddSession(mid string, pass string) (string, error) {
	session := utils.Md5(mid + pass)
	return session, gredis.Set(session, mid, 3600*24*7)
}

func CheckSession(mid string, session string) bool {
	//todo
	return gredis.Exists(session)
}

func CheckAuthWithHostname(mid string, cryptoStr string, hostname string) bool {
	equipment, err := models.GetEquipmentByHostnameMid(mid, hostname)
	if err != nil {
		logging.Error(err)
		return false
	}
	if equipment == nil {
		return false
	}
	return utils.IsEqual(cryptoStr, fmt.Sprintf("%s.%s.%s", mid, hostname, equipment.Jct))
}
