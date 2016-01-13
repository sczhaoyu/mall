package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerUserLoginLog struct {
	Id         int64     `json:"id"`
	UserId     int       `json:"userId"`
	UserName   string    `json:"userName"`
	LoginIp    string    `json:"loginIp"`
	CreateTime time.Time `json:"createTime"`
}

func (s *SellerUserLoginLog) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerUserLoginLog) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerUserLoginLog{})
	return err
}
func GetSellerUserLoginLogById(id int64) (*SellerUserLoginLog, error) {
	var ret SellerUserLoginLog
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerUserLoginLog Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerUserLoginLog) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
