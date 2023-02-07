// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	UserNotExistErrCode        = 10005
	PublishErrCode             = 10006
	UploadErrCode              = 10007
	InsertErrCode              = 10008
	BadReqErrCode              = 10009
	ActionFavoriteErrCode      = 10010
	ListFavoriteErrCode        = 10011
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, "User not exist")
	PublishErr             = NewErrNo(PublishErrCode, "publish err")
	UploadErr              = NewErrNo(UploadErrCode, "tos upload err")
	InsertErr              = NewErrNo(InsertErrCode, "insert database err")
	BadReqErr              = NewErrNo(BadReqErrCode, "bad request err")
	ActionFavoriteErr      = NewErrNo(ActionFavoriteErrCode, "action favorite err")
	ListFavoriteErr        = NewErrNo(ListFavoriteErrCode, "list favorite err")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
