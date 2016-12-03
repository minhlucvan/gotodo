package i18n

import "github.com/minhlucvan/gotodo/config"

// I18n setup
var I18N = &li18n{
	LOCATION: config.APPPATH + "/languages/",
	DEFAULT_LANG: "en_us",
	LANGS: []lang{
		{
			Key: "en_us",
			Name: "English",
			File: "locale_en-US.ini",
		},
		{
			Key: "vi_vn",
			Name: "vietnamese",
			File: "locale_vi-VN.ini",
		},
	},
}