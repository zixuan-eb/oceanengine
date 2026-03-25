package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdProductGetRequest 获取全域推广计划下商品列表 API Request
type AdProductGetRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// StartDate 开始日期，格式：YYYY-MM-DD
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束日期，格式：YYYY-MM-DD
	EndDate string `json:"end_date,omitempty"`
	// Fields 需要查询的指标，见返回参数
	Fields []string `json:"fields,omitempty"`
	// Filtering 过滤器
	Filtering *AdProductGetFilter `json:"filtering,omitempty"`
	// OrderField 排序指标字段，仅支持fields中所选指标
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码，1-100000，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，1-100，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// AdProductGetFilter 过滤器
type AdProductGetFilter struct {
	// SearchKeyWord 搜索关键词，支持商品ID和商品名称模糊搜索
	SearchKeyWord string `json:"search_key_word,omitempty"`
}

// Encode implements GetRequest interface
func (r AdProductGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	values.Set("start_date", r.StartDate)
	values.Set("end_date", r.EndDate)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdProductGetResponse 获取全域推广计划下商品列表 API Response
type AdProductGetResponse struct {
	model.BaseResponse
	Data *AdProductGetResult `json:"data,omitempty"`
}

// AdProductGetResult 响应数据
type AdProductGetResult struct {
	// ProductList 商品列表
	ProductList []AdProduct `json:"product_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// AdProduct 商品信息
type AdProduct struct {
	// ProductInfo 商品详情
	ProductInfo *ProductInfo `json:"product_info,omitempty"`
	// StatsInfo 消耗指标
	StatsInfo map[string]interface{} `json:"stats_info,omitempty"`
}
