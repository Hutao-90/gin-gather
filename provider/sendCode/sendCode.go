package sendCode

import (
	sysConfig "debox/config"
	"debox/provider/cachepackage"
	"debox/provider/logger"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func GetCaptcha(phoneNumber string) bool {
	//系统配置
	globalConfig := sysConfig.GetConfig()
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		globalConfig.Sms.SecretId,
		globalConfig.Sms.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = globalConfig.Sms.Url
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)

	//生成验证码
	captcha := CreateCaptcha(4)
	//缓存验证码
	cache := cachepackage.GetCacheInstance()
	cacheKey := phoneNumber + captcha
	cache.Set(cacheKey, captcha, 300*time.Second) //五分钟

	fmt.Printf("验证码：", cacheKey)
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{phoneNumber})
	request.SmsSdkAppId = common.StringPtr("1400847174")
	request.SignName = common.StringPtr("北京骇狼科技有限公司")
	request.TemplateId = common.StringPtr("1896089")
	request.TemplateParamSet = common.StringPtrs([]string{captcha, "5"})
	request.SessionContext = common.StringPtr("测试")

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logger.LogInstance.WithFields(logrus.Fields{
			"message": "短信发送失败",
		}).Info(err)
		fmt.Printf("An API error has returned: %s", err)
		return false
	}
	if err != nil {
		logger.LogInstance.WithFields(logrus.Fields{
			"message": "短信发送失败",
		}).Info(err)
		return false
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
	return true
}

func CreateCaptcha(num int) string {
	str := "1"
	for i := 0; i < num; i++ {
		str += strconv.Itoa(0)
	}
	str10 := str
	int10, err := strconv.ParseInt(str10, 10, 32)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		j := int32(int10)
		return fmt.Sprintf("%0"+strconv.Itoa(num)+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(j))
	}
}
