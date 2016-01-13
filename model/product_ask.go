package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductAsk struct {
	Id           int64     `json:"id"`
	SellerId     int       `json:"sellerId"`     //商家ID
	ProductId    int       `json:"productId"`    //商品ID
	UserId       int       `json:"userId"`       //咨询人ID
	UserName     string    `json:"userName"`     //咨询人账号
	AskContent   string    `json:"askContent"`   //咨询内容
	ReplyId      int       `json:"replyId"`      //回复人ID
	ReplyName    string    `json:"replyName"`    //回复人账号
	ReplyContent string    `json:"replyContent"` //回复内容
	CreateTime   time.Time `json:"createTime"`   //创建时间
	ReplyTime    time.Time `json:"replyTime"`    //回复时间
	State        int       `json:"state"`        //1、咨询；2、已经回答；3、前台显示；4、删除
}

func (p *ProductAsk) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductAsk) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductAsk{})
	return err
}
func GetProductAskById(id int64) (*ProductAsk, error) {
	var ret ProductAsk
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductAsk Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductAsk) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
