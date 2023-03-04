package loginfoservice

import (
	"pp-bakcend/models"
)

type EquipmentInfo struct {
	Version  string
	Hostname string
}

type LoginInfoObj struct {
	Mid       string
	Name      string
	Avatar    string
	Equipment []EquipmentInfo
}

func GetAllLoginInfo() ([]LoginInfoObj, error) {
	loginfos := make([]LoginInfoObj, 0)
	users, err := models.GetAllUser()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		userEquipments, err := models.GetAllEquipmentsByMid(user.Mid)
		if err != nil {
			return nil, err
		}
		result := make([]EquipmentInfo, 0, len(userEquipments))
		for _, v := range userEquipments {
			result = append(result, EquipmentInfo{
				Version:  v.Version,
				Hostname: v.Hostname,
			})
		}
		loginfos = append(loginfos, LoginInfoObj{
			Mid:       user.Mid,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Equipment: result,
		})
	}
	return loginfos, nil

}
