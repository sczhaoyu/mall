package model

import (
	"errors"
	"fmt"
	"time"
)

type MemberLoginLogs struct {
	Id         int64     `json:"id"`
	MemberId   int       `json:"memberId"`
	MemberName string    `json:"memberName"`
	LoginIp    string    `json:"loginIp"`    //登录IP
	CreateTime time.Time `json:"createTime"` //登录时间
	Source     int       `json:"source"`     //会员来源1、pc；2、H5；3、Android；4、IOS ;5、微信商城
}

func (m *MemberLoginLogs) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberLoginLogs) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberLoginLogs{})
	return err
}
func GetMemberLoginLogsById(id int64) (*MemberLoginLogs, error) {
	var ret MemberLoginLogs
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberLoginLogs Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberLoginLogs) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
