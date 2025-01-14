package client

import (
	"context"
	"github.com/Szzrain/dodo-open-go/model"
	"github.com/Szzrain/dodo-open-go/tools"
)

// GetMemberNFTStatus 获取成员数字藏品判断
func (c *client) GetMemberNFTStatus(ctx context.Context, req *model.GetMemberNFTStatusReq) (*model.GetMemberNFTStatusRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getMemberNFTStatusUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetMemberNFTStatusRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMemberUPowerchainInfo 取成员高能链数字藏品信息
func (c *client) GetMemberUPowerchainInfo(ctx context.Context, req *model.GetMemberUPowerchainInfoReq) (*model.GetMemberUPowerchainInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getMemberUPowerchainInfo))
	if err != nil {
		return nil, err
	}

	result := &model.GetMemberUPowerchainInfoRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
