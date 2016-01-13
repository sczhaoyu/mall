package model

import (
	"errors"
	"fmt"
	"time"
)

type MIndexFloor struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`    //楼层名称
	OrderNo        int       `json:"orderNo"` //楼层排序号
	Status         int       `json:"status"`  //状态1显示0不显示
	Remark         string    `json:"remark"`  //楼层备注
	CreateUserId   int       `json:"createUserId"`
	CreateUserName string    `json:"createUserName"`
	CreateTime     time.Time `json:"createTime"`
	UpdateUserId   int       `json:"updateUserId"`
	UpdateUserName string    `json:"updateUserName"`
	UpdateTime     time.Time `json:"updateTime"`
}

func (m *MIndexFloor) Save() error {
	_, err := MDB.Insert(m)
	return err
}

func (m *MIndexFloor) Delete() error {
	_, err := MDB.Where("id=?", m.Id).Delete(MIndexFloor{})
	return err
}
func GetMIndexFloorById(id int64) (*MIndexFloor, error) {
	var ret MIndexFloor
	b, err := MDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MIndexFloor Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MIndexFloor) Update() error {
	_, err := MDB.Where("id=?", m.Id).Update(m)
	return err
}
