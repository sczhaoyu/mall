package model

import (
	"errors"
	"fmt"
	"time"
)

type Cart struct {
	Id             int64     `json:"id"`
	CreateTime     time.Time `json:"createTime"`     //创建时间
	MemberId       int       `json:"memberId"`       //会员ID
	Count          int       `json:"count"`          //数量
	SpecId         string    `json:"specId"`         //规格ID，多个规格用逗号分隔
	SpecInfo       string    `json:"specInfo"`       //规格详情
	ProductId      int       `json:"productId"`      //产品ID
	SellerId       int       `json:"sellerId"`       //商家ID
	ProductGoodsId int       `json:"productGoodsId"` //货品ID
}

func (c *Cart) Save() error {
	_, err := ShopDB.Insert(c)
	return err
}

func (c *Cart) Delete() error {
	_, err := ShopDB.Where("id=?", c.Id).Delete(Cart{})
	return err
}
func GetCartById(id int64) (*Cart, error) {
	var ret Cart
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Cart Not Found Value: %v", id))
	}
	return &ret, nil
}

func (c *Cart) Update() error {
	_, err := ShopDB.Where("id=?", c.Id).Update(c)
	return err
}
