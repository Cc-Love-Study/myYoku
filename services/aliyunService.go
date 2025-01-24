package services

import (
	"fmt"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	vod "github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/gin-gonic/gin"
)

type AliyunService struct {
	Client *vod.Client
}

type RetrunJSON struct {
	RequestId     string
	UploadAddress string
	UploadAuth    string
	VideoId       string
}

func NewAliyunService(client *vod.Client) *AliyunService {
	return &AliyunService{Client: client}
}

// 创建阿里云句柄
func NewAliyunClient(accessKeyId string, accessKeySecret string) (*vod.Client, error) {
	config := sdk.NewConfig()
	config.AutoRetry = true
	// Please ensure that the environment variables ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set.
	credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
	/* use STS Token
	credential := credentials.NewStsTokenCredential(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"), os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN"))
	*/
	client, err := vod.NewClientWithOptions("cn-shanghai", config, credential)
	if err != nil {
		panic(err)
	}
	return client, err
}

func (a *AliyunService) GetAliyunUploadAddrAndAuth(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	fileName := c.PostForm("fileName")
	coverUrl := c.PostForm("coverUrl")
	tags := c.PostForm("tags")

	request := vod.CreateCreateUploadVideoRequest()
	request.Title = title
	request.Description = desc
	request.FileName = fileName
	request.CoverURL = coverUrl
	request.Tags = tags
	request.AcceptFormat = "JSON"
	request.Scheme = "https"
	response, err := a.Client.CreateUploadVideo(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	// fmt.Printf("response is %#v\n", response)
	data := &RetrunJSON{
		response.RequestId,
		response.UploadAddress,
		response.UploadAuth,
		response.VideoId,
	}
	c.JSON(http.StatusOK, data)
	return
}

func (a *AliyunService) ReflashAliyunUploadAuth(c *gin.Context) {
	videoId := c.PostForm("videoId")
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	request.Scheme = "https"
	response, err := a.Client.RefreshUploadVideo(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	// fmt.Printf("response is %#v\n", response)
	data := &RetrunJSON{
		response.RequestId,
		response.UploadAddress,
		response.UploadAuth,
		response.VideoId,
	}
	c.JSON(http.StatusOK, data)
	return
}

type ReturnJsonPlay struct {
	PlayAuth string
}

func (a *AliyunService) PlayAuth(c *gin.Context) {
	videoId := c.PostForm("videoId")

	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	request.Scheme = "https"

	response, err := a.Client.GetVideoPlayAuth(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	data := &ReturnJsonPlay{
		response.PlayAuth,
	}
	c.JSON(http.StatusOK, data)
	return
}
