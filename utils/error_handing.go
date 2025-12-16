/*
此文件用于处理初始化错误
并且输出到日志
*/

package utils

import "errors" // 注意这里需要导入 errors 包

// AppendError 合并已存在的错误和新错误
// 当 existErr 为 nil 时，直接返回 newErr
// 否则将两个错误合并为一个
func AppendError(existErr, newErr error) error {
	// 使用 errors.Join 合并错误，当其中一个为 nil 时会自动忽略
	return errors.Join(existErr, newErr)
}
