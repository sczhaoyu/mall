package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberGradeIntegralLogs struct {
	Id         int64     `json:"id"`
	MemberId   int       `json:"memberId"`   //会员ID
	MemberName string    `json:"memberName"` //会员名称
	Value      int       `json:"value"`      //经验值或积分
	CreateTime time.Time `json:"createTime"` //创建时间
	OptType    int       `json:"optType"`    //具体操作1、会员注册；2、会员登录；3、商品购买；4、商品评论；5、系统添加；6、系统减少
	OptDes     string    `json:"optDes"`     //操作描述，订单记录订单号，商品评论记录商品ID
	Type       int       `json:"type"`       //1、经验值；2、积分
}

func (m *MemberGradeIntegralLogs) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberGradeIntegralLogs) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberGradeIntegralLogs{})
	return err
}
func GetMemberGradeIntegralLogsById(id int64) (*MemberGradeIntegralLogs, error) {
	var ret MemberGradeIntegralLogs
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberGradeIntegralLogs Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberGradeIntegralLogs) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
