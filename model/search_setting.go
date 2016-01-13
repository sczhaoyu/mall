package model

import (
	"errors"
	"fmt"
	"time"
)

type SearchSetting struct {
	Id               int64     `json:"id"`
	Keyword          string    `json:"keyword"`          //搜索框下关键字设置
	KeywordFilter    int       `json:"keywordFilter"`    //关键字过滤1、不过滤；2、过滤
	IndexProductId   int       `json:"indexProductId"`   //索引处理到最大得商品ID，0为没有处理
	IndexProductTime time.Time `json:"indexProductTime"` //上次索引处理的时间
	UpdateTime       time.Time `json:"updateTime"`       //更新时间
}

func (s *SearchSetting) Save() error {
	_, err := ShopDB.Insert(s)
	return err
}

func (s *SearchSetting) Delete() error {
	_, err := ShopDB.Where("id=?", s.Id).Delete(SearchSetting{})
	return err
}
func GetSearchSettingById(id int64) (*SearchSetting, error) {
	var ret SearchSetting
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("SearchSetting Not Found Value: %v", id))
	}
	return &ret, nil
}

func (s *SearchSetting) Update() error {
	_, err := ShopDB.Where("id=?", s.Id).Update(s)
	return err
}
