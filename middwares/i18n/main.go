package i18n

import (
	"github.com/iris-contrib/middleware/i18n"
	"github.com/kataras/iris"
)

// i18n define
type lang struct {
	Key string
	Name string
	File string
}

type langs map[string]lang

type li18n struct {
	LOCATION string
	DEFAULT_LANG string
	LANGS []lang
}


var langData = I18N.LANGS
var langMap = map[string]string{}
var defaultLang = I18N.DEFAULT_LANG
var Handler iris.HandlerFunc = nil

func init(){
	for _, lang := range langData {
		langMap[lang.Key] = I18N.LOCATION + lang.File
	}
	
	Handler = i18n.New(i18n.Config{Default: defaultLang, Languages: langMap})
}
