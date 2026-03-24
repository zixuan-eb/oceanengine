package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// UniPromotionDataGet 获取商品全域推广数据明细
// 接口: v1.0/qianchuan/report/uni_promotion/data/get/
// 支持按计划(SITE_PROMOTION_PRODUCT_AD)或商品(SITE_PROMOTION_PRODUCT_PRODUCT)维度查询
func UniPromotionDataGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.UniPromotionDataGetRequest) (*report.UniPromotionDataGetResult, error) {
	var resp report.UniPromotionDataGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/uni_promotion/data/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
