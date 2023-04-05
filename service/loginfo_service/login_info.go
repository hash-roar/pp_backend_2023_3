package loginfoservice

import (
	"pp-bakcend/models"
	"time"
)

type EquipmentInfo struct {
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
}

type LoginInfoObj struct {
	Mid         string          `json:"mid"`
	Name        string          `json:"name"`
	Avatar      string          `json:"avatar"`
	Equipment   []EquipmentInfo `json:"equipment"`
	LatestLogin time.Time       `json:"latest_login"`
	Total       int64           `json:"total"`
	Sponsor     int64           `json:"sponsor"`
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
			Mid:         user.Mid,
			Name:        user.Name,
			Avatar:      user.Avatar,
			Equipment:   result,
			LatestLogin: user.UpdatedAt,
			Total:       user.TotalLogin,
			Sponsor:     user.Sponsor,
		})
	}
	return loginfos, nil

}
