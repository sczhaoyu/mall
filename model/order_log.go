package model

import (
	"errors"
	"fmt"
	"time"
)

type OrderLog struct {
	Id            int64     `json:"id"`
	OperatingId   int       `json:"operatingId"` //操作人，系统操作为0
	OperatingName string    `json:"operatingName"`
	OrdersId      int       `json:"ordersId"`
	OrdersSn      string    `json:"ordersSn"`
	Content       string    `json:"content"` //内容
	CreateTime    time.Time `json:"createTime"`
}

func (o *OrderLog) Save() error {
	_, err := ShopDB.Insert(o)
	return err
}

func (o *OrderLog) Delete() error {
	_, err := ShopDB.Where("id=?", o.Id).Delete(OrderLog{})
	return err
}
func GetOrderLogById(id int64) (*OrderLog, error) {
	var ret OrderLog
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("OrderLog Not Found Value: %v", id))
	}
	return &ret, nil
}

func (o *OrderLog) Update() error {
	_, err := ShopDB.Where("id=?", o.Id).Update(o)
	return err
}
