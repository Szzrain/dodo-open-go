package client

import (
	"context"
	"github.com/Szzrain/dodo-open-go/model"
	"github.com/Szzrain/dodo-open-go/tools"
)

// SendDirectMessage 发送私信
func (c *client) SendDirectMessage(ctx context.Context, req *model.SendDirectMessageReq) (*model.SendDirectMessageRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	req.MessageType = req.MessageBody.MessageType()
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(sendDirectMessageUri))
	if err != nil {
		return nil, err
	}

	result := &model.SendDirectMessageRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
