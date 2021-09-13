package validation

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"sync"
)

// Result 表单验证后返回的数据结构.
type Result map[string][]string

type singleton struct{}

var instance *singleton

var mu sync.Mutex

// GetInstance 一个带安全锁的单例模式.
func GetInstance() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// SetupValidator 初始化表单验证模块.
func (singleton) SetupValidator() (ut.Translator, *validator.Validate, error) {
	// 验证器初始化
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()
	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := f.Tag.Get("label") + " "
		if name == "-" {
			return ""
		}
		return name
	})

	// 注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, nil, err
	}
	return trans, validate, nil
}
