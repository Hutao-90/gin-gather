package cosService

import (
	"context"
	sysConfig "debox/config"
	"debox/provider/logger"
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 上传文件至腾讯cos
func UploadCos(locImg string, imgUrl string) string {

	sysConfig := sysConfig.GetConfig()
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(sysConfig.Cos.Url)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(sysConfig.Cos.SecretId),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv(sysConfig.Cos.SecretKey), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//name := imgUrl
	// 1.通过字符串上传对象
	//f := strings.NewReader("test")
	//
	//_, err := c.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	panic(err)
	//}
	//// 2.通过本地文件上传对象
	//_, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
	//if err != nil {
	//	panic(err)
	//}
	// 3.通过文件流上传对象
	fd, err := os.Open(locImg)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"message": "上传至cos",
			"error":   err,
		}).Error(err)
		return ""
	}
	defer fd.Close()
	_, err = c.Object.Put(context.Background(), imgUrl, fd, nil)
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"message": "上传至cos",
			"error":   err,
		}).Error(err)
		return ""
	}

	return sysConfig.Cos.Url + "/" + imgUrl
}
