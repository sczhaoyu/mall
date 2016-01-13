package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemRoles struct {
	Id         int64     `json:"id"`
	RolesName  string    `json:"rolesName"` //角色名称
	Content    string    `json:"content"`   //角色描述
	UserId     int       `json:"userId"`    //创建人
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status"`   //1、未删除2、删除
	RoleCode   string    `json:"roleCode"` //角色code,唯一
}

func (s *SystemRoles) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemRoles) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemRoles{})
	return err
}
func GetSystemRolesById(id int64) (*SystemRoles, error) {
	var ret SystemRoles
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemRoles Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemRoles) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
