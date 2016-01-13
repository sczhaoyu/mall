package model

import (
	"errors"
	"fmt"
	"time"
)

type Settlement struct {
	Id                  int64     `json:"id"`
	SettleCycle         string    `json:"settleCycle"`         //结算周期
	SellerId            int       `json:"sellerId"`            //商家ID
	SellerName          string    `json:"sellerName"`          //商家名称
	MoneyOrder          float64   `json:"moneyOrder"`          //实际应收金额总计：orders.money_order总和
	MoneyPaidBalance    float64   `json:"moneyPaidBalance"`    //余额账户支付总金额：money_paid_balance总和
	MoneyPaidReality    float64   `json:"moneyPaidReality"`    //现金支付金额：money_paid_reality总和
	MoneyGiftcardAmount float64   `json:"moneyGiftcardAmount"` //礼品卡抵用金额：money_giftcard_amount总和
	MoneyCouponcode     float64   `json:"moneyCouponcode"`     //优惠码优惠金额：money_couponcode总和
	MoneyLowerAmount    float64   `json:"moneyLowerAmount"`    //下单立减金额：money_lower_amount总和
	MoneyIntegral       float64   `json:"moneyIntegral"`       //积分换算金额
	Integral            int       `json:"integral"`            //使用积分数量
	MoneyBack           float64   `json:"moneyBack"`           //当月退款的总金额：退货表back_money总和
	MoneyOther          float64   `json:"moneyOther"`          //其他金额，应支付给商家的其他金额
	MoneyOtherType      int       `json:"moneyOtherType"`      //其他金额加减状态：1、支付增加；2、支付扣减
	MoneyOtherReason    string    `json:"moneyOtherReason"`    //其他金额理由
	Commision           float64   `json:"commision"`           //佣金
	Payable             float64   `json:"payable"`             //应付金额：money_order-money_back-commision
	Status              int       `json:"status"`              //结算状态：1、账单生成；2、平台审核通过；3、商家核对通过；4、商家核对质疑；5、对账完成；6、支付完成
	SellerDoubt         string    `json:"sellerDoubt"`         //商家质疑
	PlatformExplain     string    `json:"platformExplain"`     //平台解释
	CreateTime          time.Time `json:"createTime"`
	UpdateTime          time.Time `json:"updateTime"`
	UpdateUserId        int       `json:"updateUserId"`   //修改者ID
	UpdateUserName      string    `json:"updateUserName"` //修改者名称
}

func (s *Settlement) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *Settlement) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(Settlement{})
	return err
}
func GetSettlementById(id int64) (*Settlement, error) {
	var ret Settlement
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Settlement Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *Settlement) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
