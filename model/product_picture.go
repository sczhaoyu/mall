package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductPicture struct {
	Id          int64     `json:"id"`
	ProductId   int       `json:"productId"`   //商品ID
	ImagePath   string    `json:"imagePath"`   //图片路径
	Sort        int       `json:"sort"`        //排序
	CreateId    int       `json:"createId"`    //上传人
	CreateTime  time.Time `json:"createTime"`  //上传时间
	SellerId    int       `json:"sellerId"`    //商家ID
	State       int       `json:"state"`       //1、启用；2、不启用
	ProductLead int       `json:"productLead"` //商品主图1、主图；2、非主图，主图只能有一张
}

func (p *ProductPicture) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductPicture) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductPicture{})
	return err
}
func GetProductPictureById(id int64) (*ProductPicture, error) {
	var ret ProductPicture
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductPicture Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductPicture) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
