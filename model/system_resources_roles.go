package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemResourcesRoles struct {
	Id          int64     `json:"id"`
	ResourcesId int       `json:"resourcesId"`
	RolesId     int       `json:"rolesId"`
	CreateTime  time.Time `json:"createTime"`
}

func (s *SystemResourcesRoles) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemResourcesRoles) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemResourcesRoles{})
	return err
}
func GetSystemResourcesRolesById(id int64) (*SystemResourcesRoles, error) {
	var ret SystemResourcesRoles
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemResourcesRoles Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemResourcesRoles) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
