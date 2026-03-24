package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UniPromotionConfigGetRequest 获取商品全域推广可用维度和指标 API Request
// 接口: v1.0/qianchuan/report/uni_promotion/config/get/
type UniPromotionConfigGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// DataTopics 数据主题列表，可选值：
	// SITE_PROMOTION_PRODUCT_AD  商品全域-计划数据
	// SITE_PROMOTION_PRODUCT_PRODUCT  商品全域-商品数据
	DataTopics []string `json:"data_topics"`
}

// Encode implement GetRequest interface
func (r UniPromotionConfigGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("data_topics", string(util.JSONMarshal(r.DataTopics)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// UniPromotionConfigGetResponse 获取商品全域推广可用维度和指标 API Response
type UniPromotionConfigGetResponse struct {
	model.BaseResponse
	Data *UniPromotionConfigGetResult `json:"data,omitempty"`
}

// UniPromotionConfigGetResult 配置结果
type UniPromotionConfigGetResult struct {
	// CustomConfigDatas 数据主题配置列表
	CustomConfigDatas []UniPromotionConfigData `json:"custom_config_datas,omitempty"`
}

// UniPromotionConfigData 数据主题配置
type UniPromotionConfigData struct {
	// DataTopic 数据主题
	DataTopic string `json:"data_topic,omitempty"`
	// Dimensions 维度列表
	Dimensions []UniPromotionConfigDimension `json:"dimensions,omitempty"`
	// Metrics 指标列表
	Metrics []UniPromotionConfigMetric `json:"metrics,omitempty"`
	// QueryLimit 查询限制
	QueryLimit *UniPromotionQueryLimit `json:"query_limit,omitempty"`
}

// UniPromotionConfigDimension 维度配置
type UniPromotionConfigDimension struct {
	// Field 维度字段
	Field string `json:"field,omitempty"`
	// Name 维度名称
	Name string `json:"name,omitempty"`
	// Description 维度描述
	Description string `json:"description,omitempty"`
	// Sortable 是否支持排序
	Sortable bool `json:"sort_able,omitempty"`
	// Filterable 是否支持筛选
	Filterable bool `json:"filterable,omitempty"`
	// FilterOnly 是否仅用于过滤（不作为查询维度）
	FilterOnly bool `json:"filter_only,omitempty"`
	// IsRequired 是否必填
	IsRequired bool `json:"is_required,omitempty"`
}

// UniPromotionConfigMetric 指标配置
type UniPromotionConfigMetric struct {
	// Field 指标字段
	Field string `json:"field,omitempty"`
	// Name 指标名称
	Name string `json:"name,omitempty"`
	// Description 指标描述
	Description string `json:"description,omitempty"`
	// Unit 单位：0-数量/比率 3-元 10-百分比
	Unit int `json:"unit,omitempty"`
	// Sortable 是否支持排序
	Sortable bool `json:"sort_able,omitempty"`
}

// UniPromotionQueryLimit 查询限制
type UniPromotionQueryLimit struct {
	// EarliestDay 最早可查询日期
	EarliestDay string `json:"earliest_day,omitempty"`
	// MaxDay 最大查询天数
	MaxDay string `json:"max_day,omitempty"`
	// MaxDayByDay 按天查询最大天数
	MaxDayByDay string `json:"max_day_by_day,omitempty"`
	// MaxDayByHour 按小时查询最大天数
	MaxDayByHour string `json:"max_day_by_hour,omitempty"`
	// MaxDimensionLimit 最大维度数
	MaxDimensionLimit string `json:"max_dimension_limit,omitempty"`
}
