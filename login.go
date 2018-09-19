package main

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var cookieMap = make(map[string]*http.Cookie)
var indexUrl = "http://www.12306.cn"
var hb = map[string]string{
	"scrAvailSize":      "TeRS",
	"os":                "hAqN",
	"cookieEnabled":     "VPIf",
	"javaEnabled":       "yD16",
	"hasLiedOs":         "ci5c",
	"srcScreenSize":     "tOHY",
	"scrDeviceXDPI":     "3jCe",
	"hasLiedResolution": "3neK",
	"appcodeName":       "qT7b",
	"timeZone":          "q5aJ",
	"userAgent":         "0aew",
	"mimeTypes":         "jp76",
	"userLanguage":      "hLzX",
	"plugins":           "ks0Q",
	"appMinorVersion":   "qBVW",
	"webSmartID":        "E3gR",
	"openDatabase":      "V8vl",
	"indexedDb":         "3sw-",
	"historyList":       "kU5z",
	"scrAvailHeight":    "88tV",
	"adblock":           "FMQw",
	"jsFonts":           "EOQP",
	"sessionStorage":    "HVia",
	"browserName":       "-UVA",
	"cookieCode":        "VySQ",
	"touchSupport":      "wNLf",
	"hasLiedBrowser":    "2xC5",
	"localCode":         "lEnu",
	"browserLanguage":   "q4f3",
	"browserVersion":    "d435",
	"hasLiedLanguages":  "j5po",
	"doNotTrack":        "VEek",
	"scrWidth":          "ssI5",
	"flashVersion":      "dzuS",
	"online":            "9vyE",
	"scrColorDepth":     "qmyu",
	"cpuClass":          "Md7A",
	"storeDb":           "Fvje",
	"systemLanguage":    "e6OK",
	"scrAvailWidth":     "E-lJ",
	"scrHeight":         "5Jwy",
	"localStorage":      "XM7l",
}

/*
url:  http://dynamic.12306.cn/otn/queryDishonest/index
Set-Cookie: route=c5c62a339e7744272a54643b3be5bf64; Path=/
Set-Cookie: JSESSIONID=C26159DAE07D25467696494FEE4A30D5; Path=/otn
Set-Cookie: BIGipServerotn=1943601418.50210.0000; path=/
*/
var query_dishonest_url = "http://dynamic.12306.cn/otn/queryDishonest/index"

func query_dishonest() error {
	fmt.Println("****************** query_dishonest *******************")
	defer func() {
		fmt.Println("--------------- query_dishonest ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "dynamic.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("Upgrade-Insecure-Requests", "1")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	header.Add("Referer", "http://www.12306.cn/mormhweb/")
	header.Add("Accept-Encoding", "gzip, deflate")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := doHttpRequest(query_dishonest_url, "GET", nil, nil, header)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	//b, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("response body: %s\n", string(b))
	return nil
}

/*
url: http://dynamic.12306.cn/otn/board/init
Set-Cookie: JSESSIONID=ADF300DE832E52C7D0D1D7844097D252; Path=/otn
Set-Cookie: route=9036359bb8a8a461c164a04f8f50b252; Path=/
Set-Cookie: BIGipServerotn=1189609738.64545.0000; path=/
*/
var board_init_url = "http://dynamic.12306.cn/otn/board/init"

func board_init() error {
	fmt.Println("****************** board_init *******************")
	defer func() {
		fmt.Println("--------------- board_init ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "dynamic.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("Upgrade-Insecure-Requests", "1")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	header.Add("Referer", "http://www.12306.cn/mormhweb/")
	header.Add("Accept-Encoding", "gzip, deflate")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := doHttpRequest(board_init_url, "GET", nil, nil, header)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(b))
	return nil
}

/*
url: https://kyfw.12306.cn/passport/web/auth/uamtk
Set-Cookie: _passport_session=cac5956d715e4e9dbf5e22db0dec46763616; Path=/passport
Set-Cookie: BIGipServerpool_passport=334299658.50215.0000; path=/
Set-Cookie: BIGipServerpassport=904397066.50215.0000; path=/
*/
var uamtk_url = "https://kyfw.12306.cn/passport/web/auth/uamtk"

func uamtk() error {
	fmt.Println("****************** uamtk *******************")
	defer func() {
		fmt.Println("--------------- uamtk ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Content-Length", "9")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	header.Add("Origin", "https://kyfw.12306.cn")
	header.Add("X-Requested-With", "XMLHttpRequest")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Add("Referer", "https://kyfw.12306.cn/otn/login/init")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//header.Add("Cookie","route=9036359bb8a8a461c164a04f8f50b252; BIGipServerotn=535822858.38945.0000")
	var cookies []*http.Cookie
	cookies = append(cookies,
		getCookie("route"),
		getCookie("BIGipServerotn"),
	)
	v := url.Values{
		"appId": {"otn"},
	}
	resp, err := doHttpRequest(uamtk_url, "POST", strings.NewReader(v.Encode()), cookies, header)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(b))
	return nil
}

/*
url: https://kyfw.12306.cn/otn/index/init
Set-Cookie: BIGipServerotn=2783183370.50210.0000; path=/
Set-Cookie: JSESSIONID=F630182DE6345A4025D6F15FBE6CDB2C; Path=/otn
Set-Cookie: route=9036359bb8a8a461c164a04f8f50b252; Path=/
*/
var otn_init_url = "https://kyfw.12306.cn/otn/index/init"

func otn_init() error {
	fmt.Println("****************** otn_init *******************")
	defer func() {
		fmt.Println("--------------- otn_init ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("Upgrade-Insecure-Requests", "1")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := doHttpRequest(otn_init_url, "GET", nil, nil, header)
	if err != nil {
		return err
	}
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(b))
	return nil
}

/*
url: https://kyfw.12306.cn/otn/HttpZF/GetJS

*/
var js_url = "https://kyfw.12306.cn/otn/HttpZF/GetJS"

func getJS() error {
	fmt.Println("****************** getJS *******************")
	defer func() {
		fmt.Println("--------------- getJS ---------------")
	}()
	header := http.Header{}
	header.Add("Accept", "*/*")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	header.Add("Cache-Control", "no-cache")
	header.Add("Connection", "keep-alive")
	//header.Add("Cookie", "JSESSIONID=E041BD155B2A8CF3BF9349512639F4CC; route=6f50b51faa11b987e576cdb301e545c4; BIGipServerotn=2597781770.64545.0000")
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Pragma", "no-cache")
	header.Add("Referer", "https://kyfw.12306.cn/otn/login/init")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	var cookies []*http.Cookie
	cookies = append(cookies,
		getCookie("JSESSIONID"),
		getCookie("route"),
		getCookie("BIGipServerotn"),
	)
	resp, err := doHttpRequest(js_url, "GET", nil, cookies, header)
	if err != nil {
		return err
	}
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	fmt.Println(resp.Status)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	writeFile(b, "/Users/huhai/test/golang/src/12306/1.js")
	//fmt.Printf("response body: %s\n", string(b))
	return nil
}

func writeFile(b []byte, path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0766);
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write(b)
}

/*
url: https://kyfw.12306.cn/otn/HttpZF/logdevice
Set-Cookie: RAIL_DEVICEID=GVNxk1ZeejVi8F-TVcBO8YC9WAoBAUiuxDqlTPcpL3pczzbo02VO1yNeytu1YjXD5xbdeFQL4ipALtmMJg7gnDr6DphsYPHUPECTX11tnXcdCRpMuByYWE_DKP05zn1Ez_puskB1BVhmJxWR1wyFx3hSUj7KaY43;
Set-Cookie: RAIL_EXPIRATION=1537564297071;
*/
var log_device_url = "https://kyfw.12306.cn/otn/HttpZF/logdevice"

func log_device() error {
	fmt.Println("****************** log_device *******************")
	defer func() {
		fmt.Println("--------------- log_device ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Accept", "*/*")
	header.Add("Referer", "https://kyfw.12306.cn/otn/index/init")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//header.Add("Cookie","JSESSIONID=F630182DE6345A4025D6F15FBE6CDB2C; route=9036359bb8a8a461c164a04f8f50b252; BIGipServerotn=2783183370.50210.0000")
	var cookies []*http.Cookie
	cookies = append(cookies,
		getCookie("JSESSIONID"),
		getCookie("route"),
		getCookie("BIGipServerotn"),
	)
	v := url.Values{
		"algID":     {"Pr04vtIurC"},
		"hashCode":  {"vSQdrUkejN70BnPT80xngycszeY6XZVMeurL7FzVYVg"},
		"FMQw":      {"0"},
		"q4f3":      {"zh-CN"},
		"VySQ":      {"FGHC_-7ZSBLRYq1x6f1pqsoBDgLRy_Fd"},
		"VPIf":      {"1"},
		"custID":    {"133"},
		"VEek":      {"unknown"},
		"dzuS":      {"0"},
		"yD16":      {"0"},
		"EOQP":      {"c227b88b01f5c513710d4b9f16a5ce52"},
		"lEnu":      {"3232241178"},
		"jp76":      {"52d67b2a5aa5e031084733d5006cc664"},
		"hAqN":      {"MacIntel"},
		"platform":  {"WEB"},
		"ks0Q":      {"d22ca0b81584fbea62237b14bd04c866"},
		"TeRS":      {"1057x1920"},
		"tOHY":      {"24xx1080x1920"},
		"Fvje":      {"i1l1o1s1"},
		"q5aJ":      {"-8"},
		"wNLf":      {"99115dfb07133750ba677d055874de87"},
		"0aew":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"},
		"E3gR":      {"ffec37f50882fef53f409c686ae28b23"},
		"timestamp": {"1537268155824"},
	}
	resp, err := doHttpRequest(log_device_url, "GET", strings.NewReader(v.Encode()), cookies, header)
	if err != nil {
		return err
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(b))
	ret := strings.TrimPrefix(string(b), "callbackFunction('")
	ret = strings.TrimSuffix(ret, "')")
	result := make(map[string]string)
	err = json.Unmarshal([]byte(ret), &result)
	if err != nil {
		return err
	}
	/*
			callbackFunction('
			{
		   		 "exp":"1537564297071",
			     "dfp":"GVNxk1ZeejVi8F-TVcBO8YC9WAoBAUiuxDqlTPcpL3pczzbo02VO1yNeytu1YjXD5xbdeFQL4ipALtmMJg7gnDr6DphsYPHUPECTX11tnXcdCRpMuByYWE_DKP05zn1Ez_puskB1BVhmJxWR1wyFx3hSUj7KaY43"
			}
			')
	*/
	fmt.Println(string(b), result)
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	setCookie("RAIL_EXPIRATION", &http.Cookie{Name: "RAIL_EXPIRATION", Value: result["exp"]})
	setCookie("RAIL_DEVICEID", &http.Cookie{Name: "RAIL_DEVICEID", Value: result["dfp"]})
	return nil
}

/*
url: https://kyfw.12306.cn/passport/captcha/captcha-image?login_site=E&module=login&rand=sjrand&0.4697784536353884
Set-Cookie: _passport_ct=8b882828417d4f268dcfdfc4db747fe8t2554; Path=/passport
返回验证码图片
*/
var get_image_url = "https://kyfw.12306.cn/passport/captcha/captcha-image"

func get_image() error {
	fmt.Println("****************** get_image *******************")
	defer func() {
		fmt.Println("--------------- get_image ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Accept", "image/webp,image/apng,image/*,*/*;q=0.8")
	header.Add("Referer", "https://kyfw.12306.cn/otn/login/init")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//header.Add("Cookie", "_passport_session=cac5956d715e4e9dbf5e22db0dec46763616; route=9036359bb8a8a461c164a04f8f50b252; BIGipServerotn=535822858.38945.0000; BIGipServerpool_passport=334299658.50215.0000; RAIL_EXPIRATION=1537603502166; RAIL_DEVICEID=dtglnHRMQpwHu9tWyKArDrvwzRnEwPLfGSFEdqbuYv6_MMGAjU9RE4iRBsb41ZndTpUUIcBU2mAvMBGfhUIZf7yorRqLwxG96Gyux95TTCzsm-rs-sXLPFTFAJDkhFeo5-MjY7wwR9xbvzpv2MrTsMDE_3bSWcGA")
	var cookies []*http.Cookie
	cookies = append(cookies,
		getCookie("_passport_session"),
		getCookie("route"),
		getCookie("BIGipServerotn"),
		getCookie("BIGipServerpool_passport"),
		getCookie("RAIL_EXPIRATION"),
		getCookie("RAIL_DEVICEID"),
	)
	v := url.Values{
		"login_site": {"E"},
		"module":     {"LOGIN"},
		"rand":       {"sjrand"},
	}
	resp, err := doHttpRequest(get_image_url, "GET", strings.NewReader(v.Encode()), cookies, header)
	if err != nil {
		return err
	}
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(b))
	return nil
}

/*
验证图片
url: https://kyfw.12306.cn/passport/captcha/captcha-check
*/
var check_image_url = "https://kyfw.12306.cn/passport/captcha/captcha-check"

func check_image() error {
	fmt.Println("****************** check_image *******************")
	defer func() {
		fmt.Println("--------------- check_image ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Content-Length", "50")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	header.Add("Origin", "https://kyfw.12306.cn")
	header.Add("X-Requested-With", "XMLHttpRequest")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Add("Referer", "https://kyfw.12306.cn/otn/login/init")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//header.Add("Cookie","_passport_session=861a5bfb5a174286a8e75a45c28e21997116; _passport_ct=376079a241d346339fda45622d3b9b93t1754; route=495c805987d0f5c8c84b14f60212447d; BIGipServerpool_passport=200081930.50215.0000; BIGipServerotn=451412234.24610.0000; RAIL_EXPIRATION=1537623474098; RAIL_DEVICEID=Tq9oEixU2wGwKT1LeRLU-cp_06j-Bd40duZPlPJtBKTsUYKxFHaB4Oc459MwJTzfP-wbMPIcDbzMqwzVRjvCXdyklLsdMY6cMnAVgnSLXDsjs1DqIRI_cikn4LINJjxxdCw4hu3gHdBdq0UnHdlATBuQU_KOJSoV; BIGipServerpassport=921174282.50215.0000
	var cookies []*http.Cookie
	cookies = append(cookies,
		getCookie("_passport_session"),
		getCookie("_passport_ct"),
		getCookie("route"),
		getCookie("BIGipServerotn"),
		getCookie("BIGipServerpool_passport"),
		getCookie("BIGipServerpassport"),
		getCookie("BIGipServerotn"),
		getCookie("RAIL_EXPIRATION"),
		getCookie("RAIL_DEVICEID"),
	)
	v := url.Values{
		"answer":     {""},
		"login_site": {"E"},
		"rand":       {"sjrand"},
	}
	resp, err := doHttpRequest(check_image_url, "POST", strings.NewReader(v.Encode()), cookies, header)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	result := make(map[string]string)
	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}
	fmt.Println(result)
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	fmt.Printf("response body: %s\n", string(b))
	return nil
}

/*
url: https://kyfw.12306.cn/passport/web/login
Set-Cookie: _passport_ct=; Max-Age=0; Expires=Thu, 01-Jan-1970 00:00:10 GMT; Path=/passport
*/
var login_url = "https://kyfw.12306.cn/passport/web/login"

func login() error {
	fmt.Println("****************** login *******************")
	defer func() {
		fmt.Println("--------------- login ---------------")
	}()
	header := http.Header{}
	header.Add("Host", "kyfw.12306.cn")
	header.Add("Connection", "keep-alive")
	header.Add("Content-Length", "58")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")
	header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	header.Add("Origin", "https://kyfw.12306.cn")
	header.Add("X-Requested-With", "XMLHttpRequest")
	header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Add("Referer", "https://kyfw.12306.cn/otn/login/init")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//header.Add("Cookie","_passport_session=574e712c7853479ca24d65be4de808b59357; _passport_ct=8e0ecfd1aa11484e919b7546a1e95685t9983; route=c5c62a339e7744272a54643b3be5bf64; BIGipServerotn=384303370.24610.0000; BIGipServerpassport=904397066.50215.0000; RAIL_EXPIRATION=1537664286529; RAIL_DEVICEID=OqdLk_2Elkol56e0ljjhPS-joMppe3Ex5EqneA0Azp6IURGK8-91taoADX5DxqtkVKEMUDyO1FZ6oIRJJ6ALcWhVGnN8dMNtmRAB6tJ2yHtRzcg0iSs6DmO7Rgk8qU3H2HoP_5q-cEejX3Has01TcXwCkNfrTUmz")
	var cookies []*http.Cookie
	cookies = append(cookies,
		getCookie("_passport_session"),
		getCookie("_passport_ct"),
		getCookie("route"),
		getCookie("BIGipServerotn"),
		getCookie("BIGipServerpassport"),
		getCookie("RAIL_EXPIRATION"),
		getCookie("RAIL_DEVICEID"),
	)
	v := url.Values{
		"username": {""},
		"password": {""},
		"appid":    {"otn"},
	}
	resp, err := doHttpRequest(login_url, "POST", strings.NewReader(v.Encode()), cookies, header)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	result := make(map[string]string)
	err = json.Unmarshal(b, &result)
	if err != nil {
		return err
	}
	fmt.Println(result)
	for _, ck := range resp.Cookies() {
		setCookie(ck.Name, ck)
	}
	fmt.Printf("response body: %s\n", string(b))
	return nil
}

func getCookie(key string) *http.Cookie {
	if ck, ok := cookieMap[key]; ok {
		return ck
	} else {
		fmt.Printf("the cookie %s is empty !\n", key)
		//panic("cookie empty !")
		return nil
	}
}

func setCookie(key string, cookie *http.Cookie) {
	if _, ok := cookieMap[key]; ok {
		fmt.Printf("the cookie %s update", key)
	}
	fmt.Printf("# Set-Cookie:  %s = %s\n", key, cookie.Value)
	cookieMap[key] = cookie
}

func loginStart() error {
	//checkError(otn_init())
	checkError(query_dishonest())
	checkError(getJS())
	checkError(uamtk())
	checkError(get_image())
	printCookie()
	return nil
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func printCookie() {
	for _, value := range cookieMap {
		fmt.Printf("--- %s = %s --- %s\n", value.Name, value.Value, value.Expires.Format("2006-01-02 15:04:05"))
	}
}
func convertGBK2Utf8(src string) []byte {
	dc := mahonia.NewDecoder("gbk")
	dst := dc.ConvertString(src)
	return []byte(dst)
}
func convertUtf82GBK(src string) string {
	ec := mahonia.NewEncoder("gbk")
	dst := ec.ConvertString(src)
	return dst
}
