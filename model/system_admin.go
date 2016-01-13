package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemAdmin struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`     //登录名
	Password   string    `json:"password"` //密码
	RoleId     int       `json:"roleId"`   //角色id
	CreateTime time.Time `json:"createTime"`
	CreateUser int       `json:"createUser"` //创建人 只能是管理员
	Tel        string    `json:"tel"`        //电话
	Status     int       `json:"status"`     //状态 1-正常 2-冻结 3-删除
}

func (s *SystemAdmin) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemAdmin) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemAdmin{})
	return err
}
func GetSystemAdminById(id int64) (*SystemAdmin, error) {
	var ret SystemAdmin
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemAdmin Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemAdmin) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
