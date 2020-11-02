package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"testing"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       int    `validate:"gte=0,lte=130"`
}

func Test_validator(t *testing.T) {
	//v:=validator.New()
	
	//内容支持翻译
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	v := validator.New()
	err := zh_translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		fmt.Println(err)
	}
	user := &User{
		FirstName: "Badger",
		//LastName:  "Smith",
		Age: 150,
	}
	err = v.Struct(user)
	if err != nil {
		fmt.Println(err)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			//return
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
			//fmt.Println(err.Namespace())
			//fmt.Println(err.Field())
			//fmt.Println(err.StructNamespace())
			//fmt.Println(err.StructField())
			//fmt.Println(err.Tag())
			//fmt.Println(err.ActualTag())
			//fmt.Println(err.Kind())
			//fmt.Println(err.Type())
			//fmt.Println(err.Value())
			//fmt.Println(err.Param())
			//fmt.Println()
		}
	}
	//myEmail := "joeybloggs.gmail.com"
	//err = v.Var(myEmail,"required,email")
	//if err != nil {
	//	fmt.Println(err) // output: Key: "" Error:Field validation for "" failed on the "email" tag
	//	return
	//}
}
