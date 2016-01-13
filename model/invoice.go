package model

import (
	"errors"
	"fmt"
	"time"
)

type Invoice struct {
	Id         int64     `json:"id"`
	Content    string    `json:"content"`    //发票抬头
	State      int       `json:"state"`      //状态1、显示；2、不显示
	CreateId   int       `json:"createId"`   //创建人
	CreateTime time.Time `json:"createTime"` //创建时间
}

func (i *Invoice) Save() error {
	_, err := ShopDB.Insert(i)
	return err
}

func (i *Invoice) Delete() error {
	_, err := ShopDB.Where("id=?", i.Id).Delete(Invoice{})
	return err
}
func GetInvoiceById(id int64) (*Invoice, error) {
	var ret Invoice
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Invoice Not Found Value: %v", id))
	}
	return &ret, nil
}

func (i *Invoice) Update() error {
	_, err := ShopDB.Where("id=?", i.Id).Update(i)
	return err
}
