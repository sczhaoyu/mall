package model

import (
	"errors"
	"fmt"
	"time"
)

type BrowseLog struct {
	Id             int64     `json:"id"`
	SiteCookie     string    `json:"siteCookie"`     //cookie埋点
	SessionId      string    `json:"sessionId"`      //session标识
	UserAgent      string    `json:"userAgent"`      //客户端详细信息
	IpAddress      string    `json:"ipAddress"`      //IP地址
	AccessedPage   string    `json:"accessedPage"`   //访问的URL
	UrlReferer     string    `json:"urlReferer"`     //前一个URL
	CreateTime     time.Time `json:"createTime"`     //访问时间
	BrowseName     string    `json:"browseName"`     //浏览器名称
	BrowserVersion string    `json:"browserVersion"` //浏览器版本
	MemberId       int       `json:"memberId"`       //用户ID，用户没有登录ID为0
	Ebi            string    `json:"ebi"`
}

func (b *BrowseLog) Save() error {
	_, err := analysisDB.Insert(b)
	return err
}

func (b *BrowseLog) Delete() error {
	_, err := analysisDB.Where("id=?", b.Id).Delete(BrowseLog{})
	return err
}
func GetBrowseLogById(id int64) (*BrowseLog, error) {
	var ret BrowseLog
	b, err := analysisDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("BrowseLog Not Found Value: %v", id))
	}
	return &ret, nil
}

func (b *BrowseLog) Update() error {
	_, err := analysisDB.Where("id=?", b.Id).Update(b)
	return err
}
