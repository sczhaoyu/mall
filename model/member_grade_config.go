package model

import (
	"errors"
	"fmt"
)

type MemberGradeConfig struct {
	Id     int64 `json:"id"`
	Grade1 int   `json:"grade1"` //注册会员经验值
	Grade2 int   `json:"grade2"` //铜牌会员经验值
	Grade3 int   `json:"grade3"` //银牌会员经验值
	Grade4 int   `json:"grade4"` //金牌会员经验值
	Grade5 int   `json:"grade5"` //钻石会员经验值
}

func (m *MemberGradeConfig) Save() error {
	_, err := ShopDB.Insert(m)
	return err
}

func (m *MemberGradeConfig) Delete() error {
	_, err := ShopDB.Where("id=?", m.Id).Delete(MemberGradeConfig{})
	return err
}
func GetMemberGradeConfigById(id int64) (*MemberGradeConfig, error) {
	var ret MemberGradeConfig
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("MemberGradeConfig Not Found Value: %v", id))
	}
	return &ret, nil
}

func (m *MemberGradeConfig) Update() error {
	_, err := ShopDB.Where("id=?", m.Id).Update(m)
	return err
}
