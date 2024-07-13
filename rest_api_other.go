package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) NewGetCurrentPositionMode() *GetGetCurrentPositionMode {
	return &GetGetCurrentPositionMode{c: c}
}

const (
	getCurrentPositionMode    = "/fapi/v1/positionSide/dual"
	changeCurrentPositionMode = "/fapi/v1/positionSide/dual"
)

type GetGetCurrentPositionMode struct {
	c *Client
}

type GetGetCurrentPositionModeResp struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

func (s *GetGetCurrentPositionMode) Do(ctx context.Context, opts ...RequestOption) (res *GetGetCurrentPositionModeResp, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getCurrentPositionMode,
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetGetCurrentPositionModeResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ChangePositionMode struct {
	c *Client
}

func (c *Client) NewChangePositionMode() *GetGetCurrentPositionMode {
	return &GetGetCurrentPositionMode{c: c}
}

type ChangePositionModeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *ChangePositionMode) Do(ctx context.Context, dualSidePosition bool, opts ...RequestOption) (res *ChangePositionModeResp, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: changeCurrentPositionMode,
		secType:  secTypeSigned,
	}
	if dualSidePosition {
		r.setParam("dualSidePosition", "true")
	} else {
		r.setParam("dualSidePosition", "false")
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ChangePositionModeResp)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
