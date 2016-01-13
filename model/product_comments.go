package model

import (
	"errors"
	"fmt"
	"time"
)

type ProductComments struct {
	Id              int64     `json:"id"`
	UserId          int       `json:"userId"`          //评价人ID
	UserName        string    `json:"userName"`        //评价人账号
	Grade           int       `json:"grade"`           //评分(1到5)
	Content         string    `json:"content"`         //评价内容
	SellerAttitude  string    `json:"sellerAttitude"`  //评价商家服务
	CreateTime      time.Time `json:"createTime"`      //评价时间
	ProductId       int       `json:"productId"`       //评价商品
	ProductGoodsId  int       `json:"productGoodsId"`  //货品ID
	SellerId        int       `json:"sellerId"`        //所属商家
	OrderSn         string    `json:"orderSn"`         //订单编号
	ReplyId         int       `json:"replyId"`         //回复人ID
	ReplyName       string    `json:"replyName"`       //回复人名称
	ReplyContent    string    `json:"replyContent"`    //回复内容
	State           int       `json:"state"`           //1、评价；2、审核通过，前台显示；3、删除
	AdminId         int       `json:"adminId"`         //审核上架人
	Description     int       `json:"description"`     //描述相符(1到5星)
	ServiceAttitude int       `json:"serviceAttitude"` //服务态度(1到5星)
	ProductSpeed    int       `json:"productSpeed"`    //发货速度(1到5星)
	LogisticsSpeed  int       `json:"logisticsSpeed"`  //物流态度（1到5星）
	ExpressAttitude int       `json:"expressAttitude"` //快递员态度(1到5星)
}

func (p *ProductComments) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *ProductComments) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(ProductComments{})
	return err
}
func GetProductCommentsById(id int64) (*ProductComments, error) {
	var ret ProductComments
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ProductComments Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *ProductComments) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
