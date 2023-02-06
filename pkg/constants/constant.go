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

package constants

const (
	//JWT
	SecretKey              = "secret key"
	Issuer                 = "douyin"
	Expire                 = 1200
	VideoTableName         = "video"
	UserTableName          = "user"
	UsersFavoriteTableName = "favorite"

	LimitQuery = 30

	IdentityKey                 = "id"
	Total                       = "total"
	Notes                       = "notes"
	NoteID                      = "note_id"
	ApiServiceName              = "demoapi"
	NoteServiceName             = "demonote"
	UserServiceName             = "user"
	VideoServiceName            = "video"
	FeedServiceName             = "feed"
	FavoriteServiceName         = "favorite"
	MySQLDefaultDSN             = "tiktok:tiktok@tcp(localhost:9910)/tiktok?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                 = "127.0.0.1:2379"
	UserServerHost              = "127.0.0.1:7777"
	VideoServerHost             = "127.0.0.1:7778"
	FeedServerHost              = "127.0.0.1:7779"
	FavoriteServerHost          = "127.0.0.1:7780"
	CPURateLimit        float64 = 80.0
	DefaultLimit                = 10

	TosAccessKey = "g2EfdHBkRt12iIUIOwWZ3_C5JXtx9OIQtL25JLdJ"
	TosSecretKey = "o36ZBDOqHGwymF-eDddH7Pg6stHvbpnTh9gemjPC"
	TosBucket    = "tiktok4396"
	TosURL       = "http://rp6phmpic.bkt.clouddn.com/"
)
