package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerTypeLogs struct {
	Id            int64     `json:"id"`
	CreateId      int       `json:"createId"`      //修改人
	CreateName    string    `json:"createName"`    //修改人账号
	Content       string    `json:"content"`       //修改内容
	CreateTime    time.Time `json:"createTime"`    //修改时间
	ProductCateId int       `json:"productCateId"` //商品分类id
}

func (s *SellerTypeLogs) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerTypeLogs) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerTypeLogs{})
	return err
}
func GetSellerTypeLogsById(id int64) (*SellerTypeLogs, error) {
	var ret SellerTypeLogs
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerTypeLogs Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerTypeLogs) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
