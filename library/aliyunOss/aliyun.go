package aliyunOss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"go.uber.org/zap"
	"hash"
	"io"
	"service-all/app/init/global"
	"time"
)

// 请填写您的AccessKeyId。
//var accessKeyId string = "<yourAccessKeyId>"
//// 请填写您的AccessKeySecret。
//var accessKeySecret string = "<yourAccessKeySecret>"
//// host的格式为 bucketname.endpoint ，请替换为您的真实信息。
//var host string = "http://bucket-name.oss-cn-hangzhou.aliyuncs.com"
//// callbackUrl为 上传回调服务器的URL，请将下面的IP和Port配置为您自己的真实信息。
//var callbackUrl string = "http://88.88.88.88:8888";
//// 用户上传文件时指定的前缀。
//var upload_dir string = "upload/"
//var expire_time int64 = 30

type AliyunConf struct {
	AccessKeyId     string
	AccessKeySecret string
	Host            string
	UploadDir       string
	CallbackUrl     string
	ExpireTime      int64
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

func (conf AliyunConf) GetPolicyToken() (string, error) {
	now := time.Now().Unix()
	expire_end := now + conf.ExpireTime
	//var tokenExpire = get_gmt_iso8601(expire_end)

	//create post policy json
	var config ConfigStruct
	config.Expiration = get_gmt_iso8601(expire_end)
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, conf.UploadDir)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(conf.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = conf.CallbackUrl
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callback_str, err := json.Marshal(callbackParam)
	if err != nil {
		global.LOG.Error("callback json err", zap.Error(err))
		return "", err
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callback_str)

	var policyToken PolicyToken
	policyToken.AccessKeyId = conf.AccessKeyId
	policyToken.Host = conf.Host
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = conf.UploadDir
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)
	response, err := json.Marshal(policyToken)
	if err != nil {
		global.LOG.Error("json err", zap.Error(err))
		return "", err
	}
	return string(response), nil
}

func get_gmt_iso8601(expire_end int64) string {
	var tokenExpire = time.Unix(expire_end, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
