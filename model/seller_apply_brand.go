package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerApplyBrand struct {
	Id          int64     `json:"id"`
	SellerId    int       `json:"sellerId"`    //所属商家
	Name        string    `json:"name"`        //品牌名称
	NameFirst   string    `json:"nameFirst"`   //品牌首字母
	Image       string    `json:"image"`       //图片大小比例3:1
	CreateId    int       `json:"createId"`    //创建人
	CreateTime  time.Time `json:"createTime"`  //创建时间
	UpdateId    int       `json:"updateId"`    //修改人
	UpdateTime  time.Time `json:"updateTime"`  //修改时间
	ExplainInfo string    `json:"explainInfo"` //品牌说明
	State       int       `json:"state"`       //0、默认；1、提交审核；2、显示中；3、审核失败；4、删除
}

func (s *SellerApplyBrand) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerApplyBrand) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerApplyBrand{})
	return err
}
func GetSellerApplyBrandById(id int64) (*SellerApplyBrand, error) {
	var ret SellerApplyBrand
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerApplyBrand Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerApplyBrand) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
