package model

import (
	"errors"
	"fmt"
	"time"
)

type MIndexFloorData struct {
	Id             int64     `json:"id"`
	IndexFloorId   int       `json:"indexFloorId"` //所属楼层ID
	DataType       int       `json:"dataType"`     //数据类型：1商品2图片链接
	RefId          int       `json:"refId"`        //数据ID（data_type为1表示商品ID）
	Title          string    `json:"title"`        //图片链接标题
	Image          string    `json:"image"`        //图片地址
	LinkUrl        string    `json:"linkUrl"`      //图片链接地址
	OrderNo        int       `json:"orderNo"`      //排序号
	Remark         string    `json:"remark"`       //数据说明
	CreateUserId   int       `json:"createUserId"`
	CreateUserName string    `json:"createUserName"`
	CreateTime     time.Time `json:"createTime"`
	UpdateUserId   int       `json:"updateUserId"`
	UpdateUserName string    `json:"updateUserName"`
	UpdateTime     time.Time `json:"updateTime"`
}

func (m *MIndexFloorData) Save() error {
	_, err := MDB.Insert(m)
	return err
}

func (m *MIndexFloorData) Delete() error {
	_, err := MDB.Where("id=?", m.Id).Delete(MIndexFloorData{})
	return err
}
func GetMIndexFloorDataById(id int64) (*MIndexFloorData, error) {
	var ret MIndexFloorData
	b, err := MDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MIndexFloorData Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MIndexFloorData) Update() error {
	_, err := MDB.Where("id=?", m.Id).Update(m)
	return err
}
