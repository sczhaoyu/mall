package model

import (
	"errors"
	"fmt"
	"time"
)

type OrderPayLog struct {
	Id         int64     `json:"id"`
	OrdersId   int       `json:"ordersId"`
	OrdersSn   string    `json:"ordersSn"`
	PayStatus  string    `json:"payStatus"`  //支付方式
	PayContent string    `json:"payContent"` //支付返回的代码
	PayMoney   float64   `json:"payMoney"`   //支付金额
	MemberId   int       `json:"memberId"`
	MemberName string    `json:"memberName"`
	CreateTime time.Time `json:"createTime"` //支付时间
}

func (o *OrderPayLog) Save() error {
	_, err := ShopDB.Insert(o)
	return err
}

func (o *OrderPayLog) Delete() error {
	_, err := ShopDB.Where("id=?", o.Id).Delete(OrderPayLog{})
	return err
}
func GetOrderPayLogById(id int64) (*OrderPayLog, error) {
	var ret OrderPayLog
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("OrderPayLog Not Found Value: %v", id))
	}
	return &ret, nil
}

func (o *OrderPayLog) Update() error {
	_, err := ShopDB.Where("id=?", o.Id).Update(o)
	return err
}
