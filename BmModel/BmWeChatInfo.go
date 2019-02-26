package BmModel

type WeChatInfo struct {
	ServerName string `json:"server-name" bson:"server-name"`
	Code       string `json:"code" bson:"code"`
	AppId      string `json:"app-id" bson:"app-id"`
	Secret     string `json:"secret" bson:"secret"`
	OpenId     string `json:"open-id" bson:"open-id"`
	SessionKey string `json:"session-key" bson:"session-key"`
}
