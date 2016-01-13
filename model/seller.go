package model

import (
	"errors"
	"fmt"
	"time"
)

type Seller struct {
	Id                int64     `json:"id"`
	MemberId          int       `json:"memberId"`          //用户ID
	Name              string    `json:"name"`              //用户名
	SellerName        string    `json:"sellerName"`        //店铺名称
	SellerLogo        string    `json:"sellerLogo"`        //店铺logo
	SellerGrade       int       `json:"sellerGrade"`       //店铺等级
	ScoreService      string    `json:"scoreService"`      //店铺评分服务
	ScoreDeliverGoods string    `json:"scoreDeliverGoods"` //店铺评分发货
	ScoreDescription  string    `json:"scoreDescription"`  //店铺评分描述
	ProductNumber     int       `json:"productNumber"`     //商品数量
	CollectionNumber  int       `json:"collectionNumber"`  //店铺收藏
	CreateTime        time.Time `json:"createTime"`        //创建时间
	SaleMoney         float64   `json:"saleMoney"`         //店铺总销售金额
	OrderCount        int       `json:"orderCount"`        //店铺总订单量
	OrderCountOver    int       `json:"orderCountOver"`    //店铺完成订单量
	SellerKeyword     string    `json:"sellerKeyword"`     //SEO关键字
	SellerDes         string    `json:"sellerDes"`         //SEO店铺描述
	AuditStatus       int       `json:"auditStatus"`       //审核状态 1、待审核；2、审核通过；3、冻结
	StoreSlide        string    `json:"storeSlide"`        //店铺轮播图
}

func (s *Seller) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *Seller) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(Seller{})
	return err
}
func GetSellerById(id int64) (*Seller, error) {
	var ret Seller
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Seller Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *Seller) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
