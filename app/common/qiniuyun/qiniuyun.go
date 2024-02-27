package qiniuyun

import (
	"context"
	"fmt"
	"miniproject/app/model"
	"os"
	"path/filepath"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gopkg.in/yaml.v2"
)

type QiniuConfig struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket    string `yaml:"bucket"`
	Domain    string `yaml:"domain"`
}

// 获得token
func Geturl(address string) string {
	// 读取配置文件
	config, err := ReadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return "error:打不开隐藏文件"
	}

	// 上传图片到七牛云
	localFilePath := address // 本地图片文件路径
	url, err := UploadFileToQiniu(config, localFilePath)
	if err != nil {
		fmt.Println("Error uploading file to Qiniu:", err)
		return "error:无法上传图片"
	}

	/*fmt.Println("File uploaded successfully. URL:", url)*/
	return url
	/*fmt.Println("Qiniu upload token:", token)*/
}

// 获得token
func Gettoken(address string) string {
	// 读取配置文件
	config, err := ReadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return "error:打不开隐藏文件"
	}
	token := GetQNToken(config)
	return token
}

//获得url

func ReadConfig(filename string) (*QiniuConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config QiniuConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func UploadFileToQiniu(config *QiniuConfig, localFilePath string) (string, error) {
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	uploader := storage.NewFormUploader(&cfg)
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	token := putPolicy.UploadToken(mac)

	ret := storage.PutRet{}
	remoteFileName := "images/" + time.Now().Format("2006-01-02") + "/" + filepath.Base(localFilePath)

	err := uploader.PutFile(context.Background(), &ret, token, remoteFileName, localFilePath, nil)
	if err != nil {
		return "", err
	}

	return config.Domain + "/" + ret.Key, nil
}

func GetQNToken(config *QiniuConfig) string {
	var maxInt uint64 = 1 << 32

	putPolicy := storage.PutPolicy{
		Scope:   config.Bucket,
		Expires: maxInt,
	}

	mac := qbox.NewMac(config.AccessKey, config.SecretKey)

	return putPolicy.UploadToken(mac)
}

// 存入植物的照片
func Plantimages() {
	var i int
	for i = 0; i < 12; i++ {
		model.Plantimages[i] = Geturl(model.Plantaddress[i])
	}
}

// 存入动物照片
func Animinalimages() {
	var i int
	for i = 0; i < 25; i++ {
		model.Animalsimages[i] = Geturl(model.Animaladdress[i])
	}
}

// 存入污染物照片
func Badbuildingimages() {
	var i int
	for i = 0; i < 4; i++ {
		model.Badbuildingimages[i] = Geturl(model.Goodbuildingaddress[i])
	}
}

// 存入环保建筑物照片
func Goodbuildingimages() {
	var i int
	for i = 0; i < 4; i++ {
		model.Goodbuildingimages[i] = Geturl(model.Goodbuildingaddress[i])
	}
}
