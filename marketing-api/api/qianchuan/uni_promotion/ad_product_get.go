package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AdProductGet 获取全域推广计划下商品列表
func AdProductGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AdProductGetRequest) (*unipromotion.AdProductGetResult, error) {
	var resp unipromotion.AdProductGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/uni_promotion/ad/product/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
