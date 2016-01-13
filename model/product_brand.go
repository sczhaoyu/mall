package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductBrand struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`       //品牌名称
	NameFirst  string    `json:"nameFirst"`  //品牌首字母
	Image      string    `json:"image"`      //品牌图片
	LookMethod int       `json:"lookMethod"` //展示方式1、图片；2、文字
	Top        int       `json:"top"`        //推荐1、推荐；2、不推荐
	Sort       int       `json:"sort"`       //排序
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateId   int       `json:"updateId"`   //更新人
	UpdateTime time.Time `json:"updateTime"` //更新时间
	State      int       `json:"state"`      //状态 0、默认；1、提交审核；2、显示中；3、审核失败；4、删除
}

func (p *ProductBrand) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductBrand) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductBrand{})
	return err
}
func GetProductBrandById(id int64) (*ProductBrand, error) {
	var ret ProductBrand
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductBrand Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductBrand) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
