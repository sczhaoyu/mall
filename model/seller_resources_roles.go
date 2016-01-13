package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerResourcesRoles struct {
	Id            int64     `json:"id"`
	ResourcesId   int       `json:"resourcesId"`
	SellerRolesId int       `json:"sellerRolesId"`
	CreateTime    time.Time `json:"createTime"`
}

func (s *SellerResourcesRoles) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerResourcesRoles) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerResourcesRoles{})
	return err
}
func GetSellerResourcesRolesById(id int64) (*SellerResourcesRoles, error) {
	var ret SellerResourcesRoles
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerResourcesRoles Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerResourcesRoles) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
