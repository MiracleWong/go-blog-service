package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// locales 多语言包
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		// 获取header 参数locale
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)

		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			// 判断 locale语言类别
			switch locale {
			case "zh":
				// RegisterDefaultTranslations 将验证器和对应的语言类型进行注册
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 将Translator 存储到全局
			c.Set("trans", trans)
		}
		c.Next()
	}
}