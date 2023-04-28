package model

import (
	"github.com/OpenIoT-Hub/openiot-server/internal/user/pack"
	"github.com/OpenIoT-Hub/openiot-server/pkg/consts"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
	"gorm.io/gorm"
)

type (
	// User 用户基础信息
	User struct {
		gorm.Model
		Name string
		// 保存实际用户密码的哈希值，防止数据库被攻击密码泄露时用户密码被盗用
		Password string
		// 为了方便正则检查和引入国际电话号，使用 string
		PhoneNum string
		Email    string
		Avatar   string
	}

	// Stuff 接口，封装部分用户功能，包括密码加密等。
	// TODO 完成 Stuff 接口的设计
	Stuff interface {
		GetHashPwdByID() // 获取数据库中的密码哈希值
	}

	// Authority 中间表，运行时缓存到 redis
	// 多对多：一个人可以对应多个职位，多个人对应同一个部门
	Authority struct {
		gorm.Model
		UserID uint
		OrgID  uint
		// 组织树路径，如 "Factory1/Dept1/Group1"
		Path     string
		Position []string
	}

	// Organization 组织结构，包含了所有的结构例如公司和部门
	// 支持自定义层级，保证对于访问的管理，这里用 level 表示
	// 例如，小组-部门-公司-集团，这是一个树结构
	Organization struct {
		gorm.Model
		// 为了组织树的正常解析，这里需要要求检查不允许使用'/'字符
		Name        string
		Level       string
		FatherOrgID uint
	}
)

// CheckUserByID 通过 UserID 查询用户是否存在
// 当用户表返回数据(count>0)时，返回 true
func CheckUserByID(userId uint) bool {
	var (
		count int64
	)

	db.Table(consts.UserTableName).Where("id = ?", userId).Count(&count)
	return count > 0
}

// CreateUser 创建新用户，input 接受 pack.User
// 先查询用户手机号码是否重复 (懒删除处理)
// 再进行用户数据的创建
// FIXME 懒删除判断未增加
// FIXME 引入 checkUserExist() 函数 ?
func CreateUser(info pack.UserBasic) int {
	var (
		count int64 = 0
	)

	if err := db.Table(consts.UserTableName).
		Where("phone_num = ?", info.PhoneNum).
		Find(&count).Error; err != nil {
		return errno.ErrorUserNotExist
	}

	if err := db.Table(consts.UserTableName).Create(&info).Error; err != nil {
		return errno.ErrorDatabaseOperate
	}

	return errno.Success
}

// RemoveUserByID 通过 ID 删除用户
// 首先查询用户是否存在，不存在则报错。
// 使用 gorm.Delete() 函数，本质上是懒删除，为 delete_at 字段添加时间戳
// FIXME 引入 checkUserExist() 函数 ?
func RemoveUserByID(userId uint) int {
	var (
		user User
	)

	if err := db.Table(consts.UserTableName).
		Where("id = ?", userId).
		Find(&user).Error; err != nil {
		return errno.ErrorUserNotExist
	}

	if err := db.Table(consts.UserTableName).
		Delete(&user).Error; err != nil {
		return errno.ErrorDatabaseOperate
	}

	return errno.Success
}

// UpdateUserByID 通过用户 ID 更新用户字段
// 先查询用户是否存在，再更新字段
func UpdateUserByID(info pack.UserBasic, userId uint) int {

	// 若 user_id 没有对应的用户，返回“用户不存在”的错误
	if !CheckUserByID(userId) {
		return errno.ErrorUserNotExist
	}

	// 更新对应的用户字段
	if err := db.Table(consts.UserTableName).
		Where("id = ?", userId).
		Updates(&info).Error; err != nil {
		return errno.ErrorDatabaseOperate
	}

	return errno.Success
}

// GetUserInfoByID 通过用户 ID 获取用户资料
// FIXME 更名函数为 GetUserBasicByID
func GetUserInfoByID(userId uint) (pack.UserBasic, int) {
	var (
		info pack.UserBasic
	)

	if err := db.Table(consts.UserTableName).
		Where("id = ?", userId).
		Find(&info).Error; err == gorm.ErrRecordNotFound {
		return info, errno.ErrorDatabaseRecordNotFound
	} else if err != nil {
		return info, errno.ErrorDatabaseOperate
	}

	return info, errno.Success
}

// ListUserInfo 获取用户列表
// FIXME 完成用户列表
func ListUserInfo(pageNum int) ([]pack.UserBasic, int) {
	var (
		info []pack.UserBasic
	)

	// FIXME how to write the offset?
	db.Table(consts.UserTableName).Limit(consts.PageSize).Offset(consts.PageNum)

	return info, errno.Success
}

// GetPositionsByID 通过用户 ID 获取用户职位
func GetPositionsByID(userId uint) ([]string, int) {
	var (
		positions []string
	)

	if err := db.Table(consts.AuthorityTableName).
		Where("user_id = ?", userId).
		Find(&positions).Error; err == gorm.ErrRecordNotFound {
		return positions, errno.ErrorDatabaseRecordNotFound
	} else if err != nil {
		return positions, errno.ErrorDatabaseOperate
	}

	return positions, errno.Success
}
