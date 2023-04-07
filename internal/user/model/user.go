package model

import (
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

	// Authority 中间表，运行时缓存到 redis
	// 多对多：一个人可以对应多个职位，多个人对应同一个部门
	Authority struct {
		gorm.Model
		UserID uint
		OrgID  uint
		// 组织树路径，如 "Factory1/Dept1/Group1"
		Path     string
		Position string
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
