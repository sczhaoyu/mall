package model

import (
	"errors"
	"fmt"
	"time"
)

type SellerTransport struct {
	Id               int64     `json:"id"`
	State            int       `json:"state"`            //状态 1-使用中 2-禁用
	SellerId         int       `json:"sellerId"`         //商家id
	TransName        string    `json:"transName"`        //模板名称
	TransType        int       `json:"transType"`        //计价方式 1-按件数 2-按重量
	TransExpress     int       `json:"transExpress"`     //快递
	TransExpressInfo string    `json:"transExpressInfo"` //快递模板信息
	TransEms         int       `json:"transEms"`         //EMS
	TransEmsInfo     string    `json:"transEmsInfo"`     //EMS模板信息
	TransMail        int       `json:"transMail"`        //平邮
	TransMailInfo    string    `json:"transMailInfo"`    //平邮模板信息
	Createtime       time.Time `json:"createtime"`       //创建时间
	TransTime        int       `json:"transTime"`        //发货时间
}

func (s *SellerTransport) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SellerTransport) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SellerTransport{})
	return err
}
func GetSellerTransportById(id int64) (*SellerTransport, error) {
	var ret SellerTransport
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SellerTransport Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SellerTransport) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
