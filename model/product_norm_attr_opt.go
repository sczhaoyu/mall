package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductNormAttrOpt struct {
	Id            int64     `json:"id"`
	ProductNormId int       `json:"productNormId"` //规格ID
	ProductId     int       `json:"productId"`     //商品ID
	SellerId      int       `json:"sellerId"`      //商家ID
	TypeAttr      int       `json:"typeAttr"`      //1、商城属性 2、自定义属性
	Type          int       `json:"type"`          //1、文字；2、图片
	Name          string    `json:"name"`          //属性名称
	Image         string    `json:"image"`         //图片
	CreateId      int       `json:"createId"`      //创建人
	CreateTime    time.Time `json:"createTime"`    //创建时间
	AttrId        int       `json:"attrId"`        //属性id
}

func (p *ProductNormAttrOpt) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductNormAttrOpt) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductNormAttrOpt{})
	return err
}
func GetProductNormAttrOptById(id int64) (*ProductNormAttrOpt, error) {
	var ret ProductNormAttrOpt
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductNormAttrOpt Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductNormAttrOpt) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
