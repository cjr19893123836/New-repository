package dao

import (
	"Supply_and_Demand/global"
	"Supply_and_Demand/http_models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserHomepageAlbumImageDao struct {
	Orm *gorm.DB
}

func NewUserHomepageAlbumImageDao() *UserHomepageAlbumImageDao {
	return &UserHomepageAlbumImageDao{
		Orm: global.DB,
	}
}

// CreateImage 创建相册图片
func (m *UserHomepageAlbumImageDao) CreateImage(albumId int64, imageUrl, imageDesc string, sortNum int8) error {
	if albumId == 0 {
		return errors.New("创建图片失败：相册ID不能为空")
	}
	if imageUrl == "" {
		return errors.New("创建图片失败：图片URL不能为空")
	}

	// 获取该相册下最大的排序号
	var maxSortNum int8
	m.Orm.Model(&http_models.UserHomepageAlbumImage{}).
		Where("album_id = ?", albumId).
		Select("COALESCE(MAX(sort_num), 0)").
		Scan(&maxSortNum)

	// 如果没有指定排序号，自动递增
	if sortNum == 0 {
		sortNum = maxSortNum + 1
	}

	image := &http_models.UserHomepageAlbumImage{
		AlbumId:    albumId,
		ImageUrl:   imageUrl,
		ImageDesc:  imageDesc,
		SortNum:    sortNum,
		UploadTime: time.Now(),
	}

	result := m.Orm.Create(image)
	if result.Error != nil {
		return errors.New("创建图片失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("创建图片失败：未成功写入数据库")
	}
	return nil
}

// DeleteImage 删除相册图片
func (m *UserHomepageAlbumImageDao) DeleteImage(id int64) error {
	if id == 0 {
		return errors.New("删除图片失败：图片ID不能为空")
	}

	result := m.Orm.Delete(&http_models.UserHomepageAlbumImage{}, id)
	if result.Error != nil {
		return errors.New("删除图片失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("删除图片失败：该图片不存在")
	}
	return nil
}

// GetImageByID 根据ID获取图片信息
func (m *UserHomepageAlbumImageDao) GetImageByID(id int64) (rt http_models.UserHomepageAlbumImage, returnError error) {
	if id == 0 {
		return rt, errors.New("获取图片失败：图片ID不能为空")
	}

	err := m.Orm.First(&rt, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rt, errors.New("图片不存在")
		}
		return rt, errors.New("获取图片失败：数据库操作异常，" + err.Error())
	}
	return rt, nil
}

// GetImagesByAlbumId 根据相册ID获取所有图片
func (m *UserHomepageAlbumImageDao) GetImagesByAlbumId(albumId int64) (rt []http_models.UserHomepageAlbumImage, returnError error) {
	if albumId == 0 {
		return rt, errors.New("获取图片列表失败：相册ID不能为空")
	}

	var images []http_models.UserHomepageAlbumImage
	err := m.Orm.Where("album_id = ?", albumId).Order("sort_num ASC").Find(&images).Error
	if err != nil {
		return rt, errors.New("获取图片列表失败：数据库操作异常，" + err.Error())
	}
	return images, nil
}

// UpdateImage 更新图片信息
func (m *UserHomepageAlbumImageDao) UpdateImage(image http_models.UserHomepageAlbumImage) error {
	if image.ID == 0 {
		return errors.New("更新图片失败：图片ID不能为空")
	}

	// 检查图片是否存在
	_, err := m.GetImageByID(image.ID)
	if err != nil {
		return err
	}

	// 使用map来更新，避免零值问题
	updateMap := make(map[string]interface{})
	if image.ImageUrl != "" {
		updateMap["image_url"] = image.ImageUrl
	}
	if image.ImageDesc != "" {
		updateMap["image_desc"] = image.ImageDesc
	}
	if image.SortNum != 0 {
		updateMap["sort_num"] = image.SortNum
	}

	if len(updateMap) == 0 {
		return errors.New("更新图片失败：没有需要更新的字段")
	}

	result := m.Orm.Model(&http_models.UserHomepageAlbumImage{}).Where("id = ?", image.ID).Updates(updateMap)
	if result.Error != nil {
		return errors.New("更新图片失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("更新图片失败：未更新任何数据")
	}
	return nil
}

// BatchDeleteImagesByAlbumId 批量删除相册下的所有图片
func (m *UserHomepageAlbumImageDao) BatchDeleteImagesByAlbumId(albumId int64) error {
	if albumId == 0 {
		return errors.New("批量删除图片失败：相册ID不能为空")
	}

	result := m.Orm.Where("album_id = ?", albumId).Delete(&http_models.UserHomepageAlbumImage{})
	if result.Error != nil {
		return errors.New("批量删除图片失败：数据库操作异常，" + result.Error.Error())
	}
	return nil
}

// UpdateImageSort 更新图片排序
func (m *UserHomepageAlbumImageDao) UpdateImageSort(id int64, sortNum int8) error {
	if id == 0 {
		return errors.New("更新图片排序失败：图片ID不能为空")
	}

	// 检查图片是否存在
	_, err := m.GetImageByID(id)
	if err != nil {
		return err
	}

	result := m.Orm.Model(&http_models.UserHomepageAlbumImage{}).Where("id = ?", id).Update("sort_num", sortNum)
	if result.Error != nil {
		return errors.New("更新图片排序失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("更新图片排序失败：未更新任何数据")
	}
	return nil
}

// GetImagesByAlbumIdPaginated 分页获取相册图片
func (m *UserHomepageAlbumImageDao) GetImagesByAlbumIdPaginated(albumId int64, page, pageSize int) (rt []http_models.UserHomepageAlbumImage, total int64, returnError error) {
	if albumId == 0 {
		return rt, 0, errors.New("获取图片列表失败：相册ID不能为空")
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var images []http_models.UserHomepageAlbumImage

	// 获取总数
	err := m.Orm.Model(&http_models.UserHomepageAlbumImage{}).Where("album_id = ?", albumId).Count(&total).Error
	if err != nil {
		return rt, 0, errors.New("获取图片总数失败：数据库操作异常，" + err.Error())
	}

	// 获取分页数据
	err = m.Orm.Where("album_id = ?", albumId).
		Order("sort_num ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&images).Error
	if err != nil {
		return rt, 0, errors.New("获取图片列表失败：数据库操作异常，" + err.Error())
	}

	return images, total, nil
}
