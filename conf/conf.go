package conf

type AppConf struct {
	RunConf    `ini:"run"`
	DBConf     `ini:"db"`
	AliyunConf `ini:"aliyun"`
}

func NewAppConf() *AppConf {
	return &AppConf{}
}

type RunConf struct {
	Address string `ini:"address"`
	Md5Key  string `ini:"md5key"`
}

type DBConf struct {
	Address  string `ini:"address"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Dbname   string `ini:"dbname"`
}

type AliyunConf struct {
	AccessKeyId     string `ini:"accessKeyId"`
	AccessKeySecret string `ini:"accessKeySecret"`
}
