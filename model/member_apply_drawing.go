package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberApplyDrawing struct {
	Id           int64     `json:"id"`
	MemberId     int       `json:"memberId"`     //会员ID
	MemberName   string    `json:"memberName"`   //会员名称
	Code         string    `json:"code"`         //提现编号
	Money        float64   `json:"money"`        //提现金额
	CreateTime   time.Time `json:"createTime"`   //申请时间
	AuditingTime time.Time `json:"auditingTime"` //审核时间
	HandleTime   time.Time `json:"handleTime"`   //处理时间
	Bank         string    `json:"bank"`         //收款银行
	BankCode     string    `json:"bankCode"`     //收款账号
	State        int       `json:"state"`        //1、提交申请；2、申请通过；3、已打款；4、处理失败
	FailReason   string    `json:"failReason"`   //失败原因
	OptId        int       `json:"optId"`        //处理人ID
	OptName      string    `json:"optName"`      //处理人名称
}

func (m *MemberApplyDrawing) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberApplyDrawing) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberApplyDrawing{})
	return err
}
func GetMemberApplyDrawingById(id int64) (*MemberApplyDrawing, error) {
	var ret MemberApplyDrawing
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberApplyDrawing Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberApplyDrawing) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
