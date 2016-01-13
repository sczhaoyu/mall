package model

import (
	"errors"
	"fmt"
	"time"
)

type NewsType struct {
	Id         int64     `json:"id"`
	Pid        int       `json:"pid"`
	ParentPath string    `json:"parentPath"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`  //序号越小，越靠前
	Image      string    `json:"image"` //分类图片
	CreateTime time.Time `json:"createTime"`
}

func (n *NewsType) Save() error {
	_, err := ShopDB.Insert(n)
	return err
}

func (n *NewsType) Delete() error {
	_, err := ShopDB.Where("id=?", n.Id).Delete(NewsType{})
	return err
}
func GetNewsTypeById(id int64) (*NewsType, error) {
	var ret NewsType
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("NewsType Not Found Value: %v", id))
	}
	return &ret, nil
}

func (n *NewsType) Update() error {
	_, err := ShopDB.Where("id=?", n.Id).Update(n)
	return err
}
