package model

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	Id                int64     `json:"id"`
	ProductCateId     int       `json:"productCateId"`     //分类ID
	Name1             string    `json:"name1"`             //商品名称建议50个字符
	Name2             string    `json:"name2"`             //商品促销信息（建议100个字符）
	Keyword           string    `json:"keyword"`           //商品关键字，用于检索商品，用逗号分隔
	ProductBrandId    int       `json:"productBrandId"`    //品牌ID
	IsSelf            int       `json:"isSelf"`            //1、自营；2、商家
	CostPrice         float64   `json:"costPrice"`         //成本价
	ProtectedPrice    float64   `json:"protectedPrice"`    //保护价，最低价格不能低于
	MarketPrice       float64   `json:"marketPrice"`       //市场价
	MallPcPrice       float64   `json:"mallPcPrice"`       //商城价
	MalMobilePrice    float64   `json:"malMobilePrice"`    //商城价Mobile
	VirtualSales      int       `json:"virtualSales"`      //虚拟销量
	ActualSales       int       `json:"actualSales"`       //实际销量
	ProductStock      int       `json:"productStock"`      //商品库存
	IsNorm            int       `json:"isNorm"`            //1、没有启用规格；2、启用规格
	NormIds           string    `json:"normIds"`           //规格ID集合
	NormName          string    `json:"normName"`          //规格属性值集合 空
	State             int       `json:"state"`             //1、刚创建；2、提交审核；3、审核通过；4、申请驳回；5、商品删除；6、上架；7、下架
	IsTop             int       `json:"isTop"`             //1、不推荐；2、推荐
	UpTime            time.Time `json:"upTime"`            //商品上架时间
	Description       string    `json:"description"`       //商品描述信息
	Packing           string    `json:"packing"`           //包装清单
	SellerId          int       `json:"sellerId"`          //商家ID
	CreateId          int       `json:"createId"`          //创建人
	CreateTime        time.Time `json:"createTime"`        //创建时间
	UpdateTime        time.Time `json:"updateTime"`        //更新时间
	SellerCateId      int       `json:"sellerCateId"`      //商家分类ID
	SellerIsTop       int       `json:"sellerIsTop"`       //商品推荐，1、不推荐；2、推荐  推荐的商品会显示到店铺的首页
	SellerState       int       `json:"sellerState"`       //店铺状态：1、店铺正常；2、店铺关闭 默认1
	CommentsNumber    int       `json:"commentsNumber"`    //评价总数
	ProductCateState  int       `json:"productCateState"`  //平台商品分类状态：1、分类正常；2、分类关闭 默认1
	IsInventedProduct int       `json:"isInventedProduct"` //是否是虚拟商品：1、实物商品；2、虚拟商品
	TransportId       int       `json:"transportId"`       //运费模板id
	MasterImg         string    `json:"masterImg"`         //主图
}

func (p *Product) Save() error {
	_, err := ShopDB.Insert(p)
	return err
}

func (p *Product) Delete() error {
	_, err := ShopDB.Where("id=?", p.Id).Delete(Product{})
	return err
}
func GetProductById(id int64) (*Product, error) {
	var ret Product
	b, err := ShopDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Product Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *Product) Update() error {
	_, err := ShopDB.Where("id=?", p.Id).Update(p)
	return err
}
