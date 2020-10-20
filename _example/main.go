package main

import (
	"fmt"
	"github.com/hnmaonanbei/go-baidu-fanyi"
)

func main() {

	/*
		GU    = "gu"    // 古吉拉特语
		ROM   = "rom"   // 罗马尼亚语
		GA    = "ga"    // 爱尔兰语
		GL    = "gl"    // 加利西亚语
		LB    = "lb"    // 卢森堡语
		TT    = "tt"    // 塔塔尔语
		TR    = "tr"    // 土耳其语
		TN    = "tn"    // 塞茨瓦纳语
		LT    = "lt"    // 立陶宛语
		TH    = "th"    // 泰语
		TE    = "te"    // 泰卢固语
		FIN   = "fin"   // 芬兰语
		TA    = "ta"    // 泰米尔语
		YO    = "yo"    // 约鲁巴语
		DAN   = "dan"   // 丹麦语
		DE    = "de"    // 德语
		QU    = "qu"    // 凯楚亚语
		EL    = "el"    // 希腊语
		EO    = "eo"    // 世界语
		EN    = "en"    // 英语
		ZH    = "zh"    // 中文
		ARA   = "ara"   // 阿拉伯语
		EU    = "eu"    // 巴斯克语
		ZU    = "zu"    // 祖鲁语
		RU    = "ru"    // 俄语
		EST   = "est"   // 爱沙尼亚语
		JPKA  = "jpka"  // 日语假名
		BE    = "be"    // 白俄罗斯语
		MS    = "may"    // 马来语
		BN    = "bn"    // 孟加拉语
		JP    = "jp"    // 日语
		BS    = "bs"    // 波斯尼亚语
		YUE   = "yue"   // 粤语
		OR    = "or"    // 奥利亚语
		VIE   = "vie"   // 越南语
		CA    = "ca"    // 加泰罗尼亚语
		CY    = "cy"    // 威尔士语
		CS    = "cs"    // 捷克语
		LV    = "lv"    // 拉脱维亚语
		FRA   = "fra"   // 法语
		PT    = "pt"    // 葡萄牙语
		SWE   = "swe"   // 瑞典语
		TL    = "tl"    // 菲律宾语
		PA    = "pa"    // 旁遮普语
		CHT   = "cht"   // 中文繁体
		KOR   = "kor"   // 韩语
		PL    = "pl"    // 波兰语
		HY    = "hy"    // 亚美尼亚语
		HR    = "hr"    // 克罗地亚语
		IU    = "iu"    // 因纽特语
		HU    = "hu"    // 匈牙利语
		HI    = "hi"    // 印地语
		BUL   = "bul"   // 保加利亚语
		HA    = "ha"    // 豪萨语
		UZ    = "uz"    // 乌兹别克语
		MI    = "mi"    // 毛利语
		MK    = "mk"    // 马其顿语
		UR    = "ur"    // 乌尔都语
		MT    = "mt"    // 马耳他语
		SLO   = "slo"   // 斯洛文尼亚语
		UK    = "uk"    // 乌克兰语
		MR    = "mr"    // 马拉提语
		AF    = "af"    // 南非语
		IS    = "is"    // 冰岛语
		IR    = "ir"    // 伊朗语
		AM    = "am"    // 阿姆哈拉语
		IT    = "it"    // 意大利语
		IW    = "iw"    // 希伯来语
		AS    = "as"    // 阿萨姆语
		SPA   = "spa"   // 西班牙语
		AZ    = "az"    // 阿塞拜疆语
		ID    = "id"    // 印尼语
		IG    = "ig"    // 伊博语
		NL    = "nl"    // 荷兰语
		PT_BR = "pt_BR" // 巴西语
		NO    = "no"    // 挪威语
		NE    = "ne"    // 尼泊尔语
		FA    = "fa"    // 波斯语
		WYW   = "wyw"   // 文言文
		KA    = "ka"    // 格鲁吉亚语
		KK    = "kk"    // 哈萨克语
		SR    = "sr"    // 塞尔维亚语
		SQ    = "sq"    // 阿尔巴尼亚语
		SW    = "sw"    // 斯瓦希里语
		KN    = "kn"    // 卡纳达语
		SK    = "sk"    // 斯洛伐克语
		SI    = "si"    // 僧加罗语
		KY    = "ky"    // 吉尔吉斯语
	*/

	tk, err := bdfanyi.Gtk()
	if err != nil {
		fmt.Println(err)
		return
	}

	options := bdfanyi.NewOptions(bdfanyi.ZH, bdfanyi.EN, tk, "")
	r, err := bdfanyi.Do("你好", options)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s -> %s \n", r.TransResult.Data[0].Src, r.TransResult.Data[0].Dst)

	options.To = bdfanyi.JP
	r, err = bdfanyi.Do("你好", options)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s -> %s \n", r.TransResult.Data[0].Src, r.TransResult.Data[0].Dst)

	options.To = bdfanyi.KOR
	r, err = bdfanyi.Do("你好", options)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s -> %s \n", r.TransResult.Data[0].Src, r.TransResult.Data[0].Dst)

	options.To = bdfanyi.MS
	r, err = bdfanyi.Do("你好", options)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s -> %s \n", r.TransResult.Data[0].Src, r.TransResult.Data[0].Dst)

}
