package manager

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manager"
	managerReq "github.com/flipped-aurora/gin-vue-admin/server/model/manager/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/forgoer/openssl"
)

type AppManagementService struct{}

// CreateAppManagement 创建应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) CreateAppManagement(ctx context.Context, appManagement *manager.AppManagement) (err error) {
	// utils.RandomString(32) 返回的是一个临时字符串，而 Go 不允许直接对临时值取地址。先生成随机字符串并赋值给变量
	randomKey := utils.RandomString(32)
	randomIv := utils.RandomString(16)
	appManagement.Key = &randomKey
	appManagement.Iv = &randomIv
	err = global.GVA_DB.Create(appManagement).Error
	if err == nil {
		filePath := filepath.Join("./uploads/json", *appManagement.Title+".json")
		err = appManagementService.generateJSONFile(appManagement, filePath)
		if err != nil {
			return err // 返回创建文件时的错误
		}
	}
	return err
}

// 共享的生成 JSON 文件的函数
func (appManagementService *AppManagementService) generateJSONFile(appManagement *manager.AppManagement, filePath string) error {
	log.Println("存储目录" + filePath)
	// 对 Android 数组中的每个值进行 AES-CBC 加密
	var androidApps []string
	err := json.Unmarshal([]byte(appManagement.Android), &androidApps)
	if err != nil {
		return err
	}
	encryptedAndroid := make([]string, len(androidApps))
	log.Println("Android数组长度：", len(androidApps))
	for i, value := range androidApps {
		strValue := string(value) // 将 value 转换为 string
		log.Println("Android数组内容：" + strValue)
		encrypted, err := openssl.AesCBCEncrypt([]byte(strValue), []byte(*appManagement.Key), []byte(*appManagement.Iv), openssl.PKCS7_PADDING)
		if err != nil {
			return err
		}
		encryptedAndroid[i] = base64.StdEncoding.EncodeToString(encrypted)
	}

	// 对 IOS 数组中的每个值进行 AES-CBC 加密
	var isoApps []string
	ioserr := json.Unmarshal([]byte(appManagement.Ios), &isoApps)
	if ioserr != nil {
		return err
	}
	encryptedIos := make([]string, len(isoApps))
	for i, value := range isoApps {
		strValue := string(value) // 将 value 转换为 string
		log.Println("Android数组内容：" + strValue)
		encrypted, err := openssl.AesCBCEncrypt([]byte(strValue), []byte(*appManagement.Key), []byte(*appManagement.Iv), openssl.PKCS7_PADDING)
		if err != nil {
			return err
		}
		encryptedIos[i] = base64.StdEncoding.EncodeToString(encrypted)
	}
	// src := []byte("https://app.peidatz.com/apk/ruojjin-release.apk")
	// key := []byte("01587459")
	// iv := []byte("00058740")
	// dst, _ := openssl.AesCBCEncrypt(src, key, iv, openssl.PKCS7_PADDING)
	// 创建 JSON 数据
	jsonData := map[string]interface{}{
		"bg":      appManagement.Bg,
		"ico":     appManagement.Ico,
		"logo":    appManagement.Logo,
		"app":     appManagement.App,
		"title":   appManagement.Title,
		"version": appManagement.Version,
		"updata":  appManagement.UpdatedAt,
		"ios":     encryptedIos,
		// "android": base64.StdEncoding.EncodeToString(dst),
		"android": encryptedAndroid,
		"nav":     appManagement.Nav, // 直接使用 appManagement.Nav
	}

	// 创建目录
	folderPath := filepath.Dir(filePath)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err // 返回创建文件夹时的错误
	}

	// 创建或覆盖 JSON 文件
	file, err := os.Create(filePath) // 创建文件，如果存在则覆盖
	if err != nil {
		return err // 返回创建文件时的错误
	}
	defer file.Close() // 确保在函数结束时关闭文件

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")     // 设置缩进格式
	return encoder.Encode(jsonData) // 写入 JSON 数据
}

// DeleteAppManagement 删除应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) DeleteAppManagement(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&manager.AppManagement{}, "id = ?", ID).Error
	return err
}

// DeleteAppManagementByIds 批量删除应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) DeleteAppManagementByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]manager.AppManagement{}, "id in ?", IDs).Error
	return err
}

// UpdateAppManagement 更新应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) UpdateAppManagement(ctx context.Context, appManagement manager.AppManagement) (err error) {
	err = global.GVA_DB.Model(&manager.AppManagement{}).Where("id = ?", appManagement.ID).Updates(&appManagement).Error
	if err == nil {
		filePath := filepath.Join("./uploads/json", *appManagement.Title+".json")

		err = appManagementService.generateJSONFile(&appManagement, filePath)
		if err != nil {
			return err // 返回创建文件时的错误
		}
	}
	return err
}

// GetAppManagement 根据ID获取应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) GetAppManagement(ctx context.Context, ID string) (appManagement manager.AppManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&appManagement).Error
	return
}

// GetAppManagementInfoList 分页获取应用管理记录
// Author [yourname](https://github.com/yourname)
func (appManagementService *AppManagementService) GetAppManagementInfoList(ctx context.Context, info managerReq.AppManagementSearch) (list []manager.AppManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&manager.AppManagement{})
	var appManagements []manager.AppManagement
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.App != nil && *info.App != "" {
		db = db.Where("app LIKE ?", "%"+*info.App+"%")
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title = ?", *info.Title)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&appManagements).Error
	return appManagements, total, err
}
func (appManagementService *AppManagementService) GetAppManagementPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
