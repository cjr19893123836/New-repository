package dao

import (
	"Supply_and_Demand/global"
	"Supply_and_Demand/http_models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserHomepageTopDao struct {
	Orm *gorm.DB
}

func NewUserHomepageTopDao() *UserHomepageTopDao {
	return &UserHomepageTopDao{
		Orm: global.DB,
	}
}

// CreateTop 创建置顶记录
func (m *UserHomepageTopDao) CreateTop(homepageId int64, contentType int8, contentId int64, topSort int8, topStartTime, topEndTime time.Time) error {
	if homepageId == 0 {
		return errors.New("创建置顶记录失败：主页ID不能为空")
	}
	if contentId == 0 {
		return errors.New("创建置顶记录失败：内容ID不能为空")
	}
	if topStartTime.After(topEndTime) {
		return errors.New("创建置顶记录失败：开始时间不能晚于结束时间")
	}

	// 使用map来创建记录，避免字段名不匹配问题
	topData := map[string]interface{}{
		"homepage_id":    homepageId,
		"content_type":   contentType,
		"content_id":     contentId,
		"top_sort":       topSort,
		"top_start_time": topStartTime,
		"top_end_time":   topEndTime,
		"create_time":    time.Now(),
	}

	result := m.Orm.Table("user_homepage_top").Create(topData)
	if result.Error != nil {
		return errors.New("创建置顶记录失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("创建置顶记录失败：未成功写入数据库")
	}
	return nil
}

// DeleteTop 删除置顶记录
func (m *UserHomepageTopDao) DeleteTop(id int64) error {
	if id == 0 {
		return errors.New("删除置顶记录失败：置顶记录ID不能为空")
	}

	result := m.Orm.Table("user_homepage_top").Where("id = ?", id).Delete(&http_models.UserHomepageTop{})
	if result.Error != nil {
		return errors.New("删除置顶记录失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("删除置顶记录失败：该置顶记录不存在")
	}
	return nil
}

// GetTopByID 根据ID获取置顶记录
func (m *UserHomepageTopDao) GetTopByID(id int64) (rt http_models.UserHomepageTop, returnError error) {
	if id == 0 {
		return rt, errors.New("获取置顶记录失败：置顶记录ID不能为空")
	}

	err := m.Orm.Table("user_homepage_top").Where("id = ?", id).First(&rt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rt, errors.New("置顶记录不存在")
		}
		return rt, errors.New("获取置顶记录失败：数据库操作异常，" + err.Error())
	}
	return rt, nil
}

// GetTopByHomepageId 根据主页ID获取置顶记录列表
func (m *UserHomepageTopDao) GetTopByHomepageId(homepageId int64) (rt []http_models.UserHomepageTop, returnError error) {
	if homepageId == 0 {
		return rt, errors.New("获取置顶记录失败：主页ID不能为空")
	}

	err := m.Orm.Table("user_homepage_top").Where("homepage_id = ?", homepageId).Order("top_sort ASC").Find(&rt).Error
	if err != nil {
		return rt, errors.New("获取置顶记录失败：数据库操作异常，" + err.Error())
	}
	return rt, nil
}

// GetTopByContent 根据内容类型和内容ID获取置顶记录
func (m *UserHomepageTopDao) GetTopByContent(contentType int8, contentId int64) (rt http_models.UserHomepageTop, returnError error) {
	if contentId == 0 {
		return rt, errors.New("获取置顶记录失败：内容ID不能为空")
	}

	err := m.Orm.Table("user_homepage_top").Where("content_type = ? AND content_id = ?", contentType, contentId).First(&rt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rt, errors.New("未找到该内容的置顶记录")
		}
		return rt, errors.New("获取置顶记录失败：数据库操作异常，" + err.Error())
	}
	return rt, nil
}

// UpdateTop 更新置顶记录信息
func (m *UserHomepageTopDao) UpdateTop(top http_models.UserHomepageTop) error {
	if top.ID == 0 {
		return errors.New("更新置顶记录失败：置顶记录ID不能为空")
	}

	// 检查记录是否存在
	_, err := m.GetTopByID(top.ID)
	if err != nil {
		return err
	}

	// 使用map来更新，避免字段名问题
	updateMap := make(map[string]interface{})

	if top.ContentType != 0 {
		updateMap["content_type"] = top.ContentType
	}
	if top.ContentId != 0 {
		updateMap["content_id"] = top.ContentId
	}
	if top.TopSort != 0 {
		updateMap["top_sort"] = top.TopSort
	}

	if len(updateMap) == 0 {
		return errors.New("更新置顶记录失败：没有需要更新的字段")
	}

	result := m.Orm.Table("user_homepage_top").Where("id = ?", top.ID).Updates(updateMap)
	if result.Error != nil {
		return errors.New("更新置顶记录失败：数据库操作异常，" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return errors.New("更新置顶记录失败：未更新任何数据")
	}
	return nil
}

// GetValidTopsByHomepageId 获取当前有效的置顶记录（未过期的）
func (m *UserHomepageTopDao) GetValidTopsByHomepageId(homepageId int64) (rt []http_models.UserHomepageTop, returnError error) {
	if homepageId == 0 {
		return rt, errors.New("获取有效置顶记录失败：主页ID不能为空")
	}

	currentTime := time.Now()
	err := m.Orm.Table("user_homepage_top").
		Where("homepage_id = ? AND top_start_time <= ? AND (top_end_time IS NULL OR top_end_time >= ?)",
			homepageId, currentTime, currentTime).
		Order("top_sort ASC").
		Find(&rt).Error
	if err != nil {
		return rt, errors.New("获取有效置顶记录失败：数据库操作异常，" + err.Error())
	}
	return rt, nil
}

// BatchDeleteTopByHomepageId 批量删除指定主页的所有置顶记录
func (m *UserHomepageTopDao) BatchDeleteTopByHomepageId(homepageId int64) error {
	if homepageId == 0 {
		return errors.New("批量删除置顶记录失败：主页ID不能为空")
	}

	result := m.Orm.Table("user_homepage_top").Where("homepage_id = ?", homepageId).Delete(&http_models.UserHomepageTop{})
	if result.Error != nil {
		return errors.New("批量删除置顶记录失败：数据库操作异常，" + result.Error.Error())
	}
	return nil
}
