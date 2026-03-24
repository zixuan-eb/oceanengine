package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// UniPromotionConfigGet 获取商品全域推广可用维度和指标
// 接口: v1.0/qianchuan/report/uni_promotion/config/get/
// 通过 data_topics 参数指定要查询的数据主题，返回该主题下可用的维度和指标列表
func UniPromotionConfigGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.UniPromotionConfigGetRequest) (*report.UniPromotionConfigGetResult, error) {
	var resp report.UniPromotionConfigGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/uni_promotion/config/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
