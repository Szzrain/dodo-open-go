package model

import "errors"

// ChannelElement 取频道列表 list element
type ChannelElement struct {
	ChannelId   string `json:"channelId"`   // 频道号
	ChannelName string `json:"channelName"` // 频道名称
	ChannelType int    `json:"channelType"` // 频道类型，1：文字频道，2：语音频道，4：帖子频道，5：链接频道，6：资料频道
	DefaultFlag int    `json:"defaultFlag"` // 默认频道标识，0：否，1：是
	GroupId     string `json:"groupId"`     // 分组ID
	GroupName   string `json:"groupName"`   // 分组名称
}

// GetChannelListReq 取频道列表 request
type GetChannelListReq struct {
	IslandId string `json:"islandId" binding:"required"` // 群号
}

func (p *GetChannelListReq) ValidParams() error {
	if p.IslandId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	return nil
}

type (
	// GetChannelInfoReq 取频道信息 request
	GetChannelInfoReq struct {
		ChannelId string `json:"channelId" binding:"required"` // 频道号
	}

	// GetChannelInfoRsp 取频道信息 response
	GetChannelInfoRsp struct {
		ChannelElement

		IslandId string `json:"islandId"` // 群号
	}
)

func (p *GetChannelInfoReq) ValidParams() error {
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	return nil
}

type (
	// SendChannelMessageReq 发送频道消息 request
	SendChannelMessageReq struct {
		ChannelId           string       `json:"channelId" binding:"required"`   // 频道号
		MessageType         MessageType  `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
		MessageBody         IMessageBody `json:"messageBody" binding:"required"` // 消息内容
		ReferencedMessageId string       `json:"referencedMessageId,omitempty"`  // 回复消息ID
	}

	// SendChannelMessageRsp 发送频道消息 response
	SendChannelMessageRsp struct {
		MessageId string `json:"messageId"` // 消息 ID
	}
)

func (p *SendChannelMessageReq) ValidParams() error {
	if p.ChannelId == "" {
		return errors.New("invalid parameter ChannelId (empty detected)")
	}
	if p.MessageBody == nil {
		return errors.New("invalid parameter MessageBody (nil detected)")
	}
	return nil
}

type (
	// EditChannelMessageReq 编辑频道消息 request
	EditChannelMessageReq struct {
		MessageId   string       `json:"messageId" binding:"required"`   // 欲编辑的消息 ID
		MessageType MessageType  `json:"messageType" binding:"required"` // 消息类型，该参数会在SDK中重新赋值，所以无需开发者主动设值
		MessageBody IMessageBody `json:"messageBody" binding:"required"` // 消息内容
	}

	// EditChannelMessageRsp 发送频道消息 response
	EditChannelMessageRsp struct {
		MessageId string `json:"messageId"` // 消息 ID
	}
)

func (p *EditChannelMessageReq) ValidParams() error {
	if p.MessageId == "" {
		return errors.New("invalid parameter MessageId (empty detected)")
	}
	if p.MessageBody == nil {
		return errors.New("invalid parameter MessageBody (nil detected)")
	}
	return nil
}

// WithdrawChannelMessageReq 撤回频道消息 request
type WithdrawChannelMessageReq struct {
	MessageId string `json:"messageId" binding:"required"` // 消息ID
	Reason    string `json:"reason,omitempty"`             // 撤回原因
}

func (p *WithdrawChannelMessageReq) ValidParams() error {
	if p.MessageId == "" {
		return errors.New("invalid parameter MessageId (empty detected)")
	}
	return nil
}
