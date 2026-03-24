package report

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UniPromotionDataGetRequest 获取商品全域推广数据明细 API Request
// 接口: v1.0/qianchuan/report/uni_promotion/data/get/
type UniPromotionDataGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// DataTopic 数据主题，可选值：
	// SITE_PROMOTION_PRODUCT_AD  商品全域-计划数据
	// SITE_PROMOTION_PRODUCT_PRODUCT  商品全域-商品数据
	DataTopic string `json:"data_topic"`
	// Dimensions 维度列表。可通过 uni_promotion/config/get/ 接口获取可用维度
	Dimensions []string `json:"dimensions"`
	// Metrics 指标列表。可通过 uni_promotion/config/get/ 接口获取可用指标
	Metrics []string `json:"metrics"`
	// StartTime 开始时间，格式：2022-08-01 00:00:00
	StartTime string `json:"start_time"`
	// EndTime 结束时间，格式：2022-08-01 23:59:59
	EndTime string `json:"end_time"`
	// Filters 过滤条件（可选）
	Filters []UniPromotionDataFilter `json:"filters,omitempty"`
	// OrderBy 排序（可选）
	OrderBy []UniPromotionDataOrderBy `json:"order_by,omitempty"`
	// Page 页码，默认 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认 20，最大 100
	PageSize int `json:"page_size,omitempty"`
}

// UniPromotionDataFilter 过滤条件
type UniPromotionDataFilter struct {
	// Field 过滤字段
	Field string `json:"field"`
	// Values 过滤值
	Values []string `json:"values"`
	// Type 字段类型：1-固定枚举值 2-固定输入值 3-数值类型
	Type int `json:"type"`
	// Operator 处理方式：7-包含
	Operator int `json:"operator"`
}

// UniPromotionDataOrderBy 排序条件
type UniPromotionDataOrderBy struct {
	// Field 排序字段（必须在 metrics 或 dimensions 中）
	Field string `json:"field"`
	// Type 排序类型：1-升序 2-降序
	Type int `json:"type"`
}

// Encode implement GetRequest interface
func (r UniPromotionDataGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("data_topic", r.DataTopic)
	values.Set("dimensions", string(util.JSONMarshal(r.Dimensions)))
	values.Set("metrics", string(util.JSONMarshal(r.Metrics)))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if len(r.Filters) > 0 {
		values.Set("filters", string(util.JSONMarshal(r.Filters)))
	}
	if len(r.OrderBy) > 0 {
		values.Set("order_by", string(util.JSONMarshal(r.OrderBy)))
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

// UniPromotionDataGetResponse 获取商品全域推广数据明细 API Response
type UniPromotionDataGetResponse struct {
	model.BaseResponse
	Data *UniPromotionDataGetResult `json:"data,omitempty"`
}

// UniPromotionDataGetResult 数据明细结果
type UniPromotionDataGetResult struct {
	// Rows 数据行
	Rows []UniPromotionDataRow `json:"rows,omitempty"`
	// Pagination 分页信息
	Pagination *model.PageInfo `json:"pagination,omitempty"`
}

// UniPromotionDataRow 数据行
type UniPromotionDataRow struct {
	// Dimensions 维度数据 key=维度字段, value=维度值
	Dimensions map[string]string `json:"dimensions,omitempty"`
	// Metrics 指标数据 key=指标字段, value=指标值
	Metrics map[string]json.Number `json:"metrics,omitempty"`
}
