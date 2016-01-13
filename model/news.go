package model

import (
	"errors"
	"fmt"
	"time"
)

type News struct {
	Id          int64     `json:"id"`
	TypeId      int       `json:"typeId"`
	TypePath    string    `json:"typePath"`
	Title       string    `json:"title"` //新闻标题
	Content     string    `json:"content"`
	Author      string    `json:"author"`      //作者
	Source      string    `json:"source"`      //来源
	IsOut       int       `json:"isOut"`       //是否是外部链接0、不是；1、是
	OutUrl      string    `json:"outUrl"`      //外部链接的URL
	Status      int       `json:"status"`      //0、不显示；1、显示
	Sort        int       `json:"sort"`        //排序
	IsRecommend int       `json:"isRecommend"` //是否推荐文章0、不是推荐文章；1、推荐文章
	CreateId    int       `json:"createId"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

func (n *News) Save() error {
	_, err := ShopDB.Insert(n)
	return err
}

func (n *News) Delete() error {
	_, err := ShopDB.Where("id=?", n.Id).Delete(News{})
	return err
}
func GetNewsById(id int64) (*News, error) {
	var ret News
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("News Not Found Value: %v", id))
	}
	return &ret, nil
}

func (n *News) Update() error {
	_, err := ShopDB.Where("id=?", n.Id).Update(n)
	return err
}
