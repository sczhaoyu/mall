package model

import (
	"errors"
	"fmt"
	"time"
)

type SystemEmall struct {
	Id           int64     `json:"id"`
	SellerOrUser int       `json:"sellerOrUser"` //商家或者用户
	SellerUserId int       `json:"sellerUserId"`
	Emall        string    `json:"emall"`      //邮件
	Content      string    `json:"content"`    //内容
	State        int       `json:"state"`      //状态1、未发送；2、发送成功；3、发送失败
	ReturnCode   string    `json:"returnCode"` //返回状态码
	CreateTime   time.Time `json:"createTime"` //创建时间
	UpdateTime   time.Time `json:"updateTime"` //发生时间
}

func (s *SystemEmall) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SystemEmall) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SystemEmall{})
	return err
}
func GetSystemEmallById(id int64) (*SystemEmall, error) {
	var ret SystemEmall
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SystemEmall Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SystemEmall) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
