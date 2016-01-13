package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerManageCate struct {
	Id              int64     `json:"id"`
	Seller          int       `json:"seller"`          //商家id
	CreateId        int       `json:"createId"`        //申请人
	CreateTime      time.Time `json:"createTime"`      //申请时间
	ProductCateId   int       `json:"productCateId"`   //申请分类id
	ProductCateName string    `json:"productCateName"` //申请分类名称,提交类目组合
	State           int       `json:"state"`           //0-默认;1-提交审核;2-审核通过;3-审核失败;4-停用
	OptId           int       `json:"optId"`           //审核人
	OptTime         time.Time `json:"optTime"`         //审核时间
	StopId          int       `json:"stopId"`          //停用人
	StopTime        time.Time `json:"stopTime"`        //停用时间
	StopReason      string    `json:"stopReason"`      //停用原由
}

func (s *SellerManageCate) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerManageCate) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerManageCate{})
	return err
}
func GetSellerManageCateById(id int64) (*SellerManageCate, error) {
	var ret SellerManageCate
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerManageCate Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerManageCate) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
