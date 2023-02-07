// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	favorite "miniTikTok/kitex_gen/favorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ActionFavorite": kitex.NewMethodInfo(actionFavoriteHandler, newActionFavoriteArgs, newActionFavoriteResult, false),
		"ListFavorite":   kitex.NewMethodInfo(listFavoriteHandler, newListFavoriteArgs, newListFavoriteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func actionFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.DouyinFavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).ActionFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ActionFavoriteArgs:
		success, err := handler.(favorite.FavoriteService).ActionFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ActionFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newActionFavoriteArgs() interface{} {
	return &ActionFavoriteArgs{}
}

func newActionFavoriteResult() interface{} {
	return &ActionFavoriteResult{}
}

type ActionFavoriteArgs struct {
	Req *favorite.DouyinFavoriteActionRequest
}

func (p *ActionFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.DouyinFavoriteActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ActionFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ActionFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ActionFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ActionFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ActionFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.DouyinFavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ActionFavoriteArgs_Req_DEFAULT *favorite.DouyinFavoriteActionRequest

func (p *ActionFavoriteArgs) GetReq() *favorite.DouyinFavoriteActionRequest {
	if !p.IsSetReq() {
		return ActionFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ActionFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type ActionFavoriteResult struct {
	Success *favorite.DouyinFavoriteActionResponse
}

var ActionFavoriteResult_Success_DEFAULT *favorite.DouyinFavoriteActionResponse

func (p *ActionFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.DouyinFavoriteActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ActionFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ActionFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ActionFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ActionFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ActionFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.DouyinFavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ActionFavoriteResult) GetSuccess() *favorite.DouyinFavoriteActionResponse {
	if !p.IsSetSuccess() {
		return ActionFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ActionFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.DouyinFavoriteActionResponse)
}

func (p *ActionFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func listFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.DouyinFavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).ListFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListFavoriteArgs:
		success, err := handler.(favorite.FavoriteService).ListFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newListFavoriteArgs() interface{} {
	return &ListFavoriteArgs{}
}

func newListFavoriteResult() interface{} {
	return &ListFavoriteResult{}
}

type ListFavoriteArgs struct {
	Req *favorite.DouyinFavoriteListRequest
}

func (p *ListFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.DouyinFavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.DouyinFavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListFavoriteArgs_Req_DEFAULT *favorite.DouyinFavoriteListRequest

func (p *ListFavoriteArgs) GetReq() *favorite.DouyinFavoriteListRequest {
	if !p.IsSetReq() {
		return ListFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type ListFavoriteResult struct {
	Success *favorite.DouyinFavoriteListResponse
}

var ListFavoriteResult_Success_DEFAULT *favorite.DouyinFavoriteListResponse

func (p *ListFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.DouyinFavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.DouyinFavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListFavoriteResult) GetSuccess() *favorite.DouyinFavoriteListResponse {
	if !p.IsSetSuccess() {
		return ListFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.DouyinFavoriteListResponse)
}

func (p *ListFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ActionFavorite(ctx context.Context, Req *favorite.DouyinFavoriteActionRequest) (r *favorite.DouyinFavoriteActionResponse, err error) {
	var _args ActionFavoriteArgs
	_args.Req = Req
	var _result ActionFavoriteResult
	if err = p.c.Call(ctx, "ActionFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFavorite(ctx context.Context, Req *favorite.DouyinFavoriteListRequest) (r *favorite.DouyinFavoriteListResponse, err error) {
	var _args ListFavoriteArgs
	_args.Req = Req
	var _result ListFavoriteResult
	if err = p.c.Call(ctx, "ListFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
