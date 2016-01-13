package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductNorm struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`       //规格名称
	Sort       int       `json:"sort"`       //排序
	Type       int       `json:"type"`       //规格类型1、文字；2、图片
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateId   int       `json:"updateId"`   //更新人
	UpdateTime time.Time `json:"updateTime"` //更新时间
	State      int       `json:"state"`      //状态 0：删除 1:正常
}

func (p *ProductNorm) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductNorm) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductNorm{})
	return err
}
func GetProductNormById(id int64) (*ProductNorm, error) {
	var ret ProductNorm
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductNorm Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductNorm) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
