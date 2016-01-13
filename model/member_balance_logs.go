package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberBalanceLogs struct {
	Id          int64     `json:"id"`
	MemberId    int       `json:"memberId"`    //会员ID
	MemberName  string    `json:"memberName"`  //会员名称
	MoneyBefore float64   `json:"moneyBefore"` //变化之前的余额
	MoneyAfter  float64   `json:"moneyAfter"`  //变化之后的余额
	Money       float64   `json:"money"`       //变化金额
	CreateTime  time.Time `json:"createTime"`  //创建时间
	State       int       `json:"state"`       //1、充值；2、退款；3、消费；4、提款；5、系统添加；6、系统减少
	Remark      string    `json:"remark"`      //操作备注
	OptId       int       `json:"optId"`       //操作人ID
	OptName     string    `json:"optName"`     //操作人名称
}

func (m *MemberBalanceLogs) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberBalanceLogs) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberBalanceLogs{})
	return err
}
func GetMemberBalanceLogsById(id int64) (*MemberBalanceLogs, error) {
	var ret MemberBalanceLogs
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberBalanceLogs Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberBalanceLogs) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
