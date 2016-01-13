package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductNormAttr struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`          //属性名称
	ProductNormId int       `json:"productNormId"` //规格ID
	Sort          int       `json:"sort"`          //排序
	CreateId      int       `json:"createId"`      //创建人
	CreateTime    time.Time `json:"createTime"`    //创建时间
	Image         string    `json:"image"`         //默认图片
}

func (p *ProductNormAttr) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductNormAttr) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductNormAttr{})
	return err
}
func GetProductNormAttrById(id int64) (*ProductNormAttr, error) {
	var ret ProductNormAttr
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductNormAttr Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductNormAttr) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
