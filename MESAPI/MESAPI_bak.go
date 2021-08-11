package main

//#include <stdlib.h>
//#include <string.h>
import "C"
import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/tiaguinho/gosoap"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"unsafe"
)
type RXml_Report struct {
	Str []string `xml:"SaveSSN_NEWResult>anyType"`
}
type RXml_Check struct {
	Str []string `xml:"CheckSSN_NEWResult>anyType"`
}
type RXml_Model struct {
	Str []string `xml:"GetSSN_CertifiedModelResult>anyType"`
}

func init()  {
	logName:=fmt.Sprintf("log/MES_%s.log",time.Now().Format("2006_01_02"))
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		os.Mkdir("./log",0666)
		file, err = os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		}
	}
}
//export api_checkSN
func api_checkSN(sn *C.char,station *C.char,sret1 *C.char,sret2 *C.char) C.int {
	msg1,msg2:="",""
	ret:= apiChecksn(C.GoString((*C.char)(sn)), C.GoString((*C.char)(station)),&msg1,&msg2)
	log.Println(msg1,msg2)
	r:=C.CString(msg1)
	r2:=C.CString(msg2)
	C.strcpy(sret1, r)
	C.strcpy(sret2, r2)
	C.free(unsafe.Pointer(r))
	C.free(unsafe.Pointer(r2))
	return (C.int)(ret)
}
//export api_ReportResult
func api_ReportResult(sn *C.char,bSuc C._Bool,sMsg *C.char) C.int {
	msg1:=""
	ret:= apiReportsn(C.GoString((*C.char)(sn)),bool((C._Bool)(bSuc)),&msg1)
	log.Println(msg1)
	r:=C.CString(msg1)
	C.strcpy(sMsg, r)
	C.free(unsafe.Pointer(r))
	return (C.int)(ret)
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func tokenDecode(token string) (r interface{}){
	tokenBase:=strings.Split(token,".")[1]
	for ;len(tokenBase)%4>0;{
		tokenBase+="="
	}
	sDec, err := base64.StdEncoding.DecodeString(tokenBase)
	if err != nil {
		log.Printf("Decode token error: %s ", err.Error())
		return
	}
	err = json.Unmarshal(sDec,&r)
	return
}
func apiChecksn(sn,sntype string,resultStr,msgStr *string) int32 {
	url:=""
	apiStr:="/factory/mes/burnStatus/"
	stationStr:=""
	tokenFname := "token.txt"
	token := ""
	isOld:=false
	res:=false
	msg:=""
	if checkFileIsExist(tokenFname) {
		bytes, _ := ioutil.ReadFile(tokenFname)
		token = string(bytes)
		isOld=strings.Contains(token,"token")
	} else {
		wsdl,station:=readSoapCfg()
		if wsdl==""{
			log.Println("NotFound token.txt and SoapUrl Config")
			return -1
		}else {
			checkok:=checkSoapModel(wsdl,sn)
			if !checkok{
				return -1
			}
			res,msg=soapCheck(wsdl,sn,station)
			log.Println("[CheckSN]",res, msg)
			if res {
				*resultStr="PASS"
				*msgStr="Check Pass"
				return 0
			}else {
				*resultStr="FAIL"
				*msgStr="Check Fail"
				return -1
			}
		}
	}
	if isOld {
		var json_in map[string]interface{}
		err := json.Unmarshal([]byte(token), &json_in)
		if err != nil {
			fmt.Printf("%+v", json_in)
			log.Println("Unmarshal token Fail")
			return -1
		}
		tokenIn, ok := json_in["token"] /*如果确定是真实的,则存在,否则不存在 */
		if ok {
			token = fmt.Sprintf("%v", tokenIn)
		} else {
			fmt.Printf("%+v", json_in)
			log.Println("Unmarshal token Fail")
			return -1
		}
	}
	r:=tokenDecode(token)
	if r==nil{
		log.Println("token Decode Fail")
		return -1
	}
	m := r.(map[string]interface{})
	if isOld{
		host, ok := m["server_url_key_record"]
		if ok{
			host_str:=fmt.Sprintf("%s",host)
			url=fmt.Sprintf("%s/burnStatus/%s",host_str,sn)
			log.Println(url)
		}else {
			*resultStr="FAIL"
			*msgStr="Get token server_url_key_record error"
			log.Println("Get token server_url_key_record Fail")
			return -1
		}
		res, msg = mesOldGet(url, token)
	}else {
		host, ok := m["host"]
		if ok{
			checkModelApi:=fmt.Sprintf("%s/factory/mes/checkModel/%s/",host,sn)
			checkok:=checkMesModel(checkModelApi,token)
			if !checkok{
				return -1
			}
			url=fmt.Sprintf("%s%s%s%s",host,apiStr,sn,stationStr)
			log.Println(url)
		}else {
			*resultStr="FAIL"
			*msgStr="Get token host error"
			log.Println("Get token host Fail")
			return -1
		}
		res, msg = mesGet(url, token)
	}
	log.Println("[CheckSN]",res, msg)
	if res {
		*resultStr="PASS"
		*msgStr="Check Pass"
		return 0
	}else {
		*resultStr="FAIL"
		*msgStr="Check Fail"
		return -1
	}

}
func apiReportsn(sn string,result bool,msgStr *string) int {
	url:=""
	apiStr:="/factory/mes/firmwareStatus/"
	stationStr:=""
	if result{
		stationStr="/1"
	}else {
		stationStr="/2"
	}
	tokenFname := "token.txt"
	token := ""
	isOld:=false
	res:=false
	msg:=""
	if checkFileIsExist(tokenFname) {
		bytes, _ := ioutil.ReadFile(tokenFname)
		token = string(bytes)
		isOld=strings.Contains(token,"token")
	} else {
		wsdl,station:=readSoapCfg()
		if wsdl==""{
			log.Println("NotFound token.txt and SoapUrl Config")
			return -1
		}else {
			res,msg=soapReport(wsdl,sn,station,result,"","burnTool")
			log.Println("[ReportSN]",res, msg)
			if res {
				*msgStr="Upload Finish"
				return 0
			}else {
				*msgStr="Upload Fail"
				return -1
			}
		}
	}
	if isOld {
		var json_in map[string]interface{}
		err := json.Unmarshal([]byte(token), &json_in)
		if err != nil {
			fmt.Printf("%+v", json_in)
			log.Println("Unmarshal token Fail")
			return -1
		}
		tokenIn, ok := json_in["token"] /*如果确定是真实的,则存在,否则不存在 */
		if ok {
			token = fmt.Sprintf("%v", tokenIn)
		} else {
			fmt.Printf("%+v", json_in)
			log.Println("Unmarshal token Fail")
			return -1
		}
	}
	r:=tokenDecode(token)
	if r==nil{
		log.Println("token Decode Fail")
		return -1
	}
	m := r.(map[string]interface{})
	if isOld{
		host, ok := m["server_url_key_record"]
		if ok{
			host_str:=fmt.Sprintf("%s",host)
			url=fmt.Sprintf("%s/firmwareStatus/%s%s",host_str,sn, stationStr)
			log.Println(url)
		}else {
			*msgStr="Get token server_url_key_record error"
			log.Println("Get token server_url_key_record Fail")
			return -1
		}
		res, msg = mesOldPut(url, token)
	}else {
		host, ok := m["host"]
		if ok {
			url = fmt.Sprintf("%s%s%s%s", host, apiStr, sn, stationStr)
			log.Println(url)
		} else {
			*msgStr = "Get token host error"
			return -1
		}
		res, msg = mesPut(url, token, "") //未烧录/0;烧录成功/1;烧录失败/2
	}
	log.Println("[ReportSN]",res, msg)
	if res {
		*msgStr="Upload Finish"
		return 0
	}else {
		*msgStr="Upload Fail"
		return -1
	}

}
func mesGet(url string, token string) (judge bool, msg string) {
	judge = false
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg = "[Error]NewRequest"
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		msg = "[Error]Header"
		return
	}
	defer res.Body.Close()
	var r interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	m := r.(map[string]interface{})
	code, ok := m["code"] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		judge = fmt.Sprintf("%v", code) == "0" //map和interface都不能直接比较
		if judge {
			judge = reflect.ValueOf(m["body"].(map[string]interface{})["uploadStatus"]).Bool()
			if !judge {
				msg = "UploadStatus Fail"
			}
		} else {
			judge = false
			msg = "Error"
		}
	} else {
		msg3, ok3 := m["status"] /*如果确定是真实的,则存在,否则不存在 */
		if ok3 {
			msg = fmt.Sprintf("%v", msg3)
		}
	}
	return
}
func mesPut(url string, token string, jsondata string) (judge bool, msg string) {
	payload := strings.NewReader(jsondata)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		msg = "[Error]NewRequest"
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		msg = "[Error]Header"
		return
	}
	defer res.Body.Close()
	var r interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	m := r.(map[string]interface{})
	code, ok := m["code"] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		judge = fmt.Sprintf("%v", code) == "0" //map和interface都不能直接比较
		msg1, ok1 := m["message"]              /*如果确定是真实的,则存在,否则不存在 */
		if ok1 {
			msg = fmt.Sprintf("%s", msg1)
		} else {
			msg2, ok2 := m["msg"] /*如果确定是真实的,则存在,否则不存在 */
			if ok2 {
				msg = fmt.Sprintf("%s", msg2)
			}
		}
	} else {
		msg3, ok3 := m["status"] /*如果确定是真实的,则存在,否则不存在 */
		if ok3 {
			msg = fmt.Sprintf("%v", msg3)
		}
	}
	return
}
func mesOldGet(url string, token string) (judge bool, msg string) { //旧系统的烧录接口
	judge = false
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg = "[Error]NewRequest"
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		msg = "[Error]Header"
		return
	}
	defer res.Body.Close()
	var r interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	m := r.(map[string]interface{})
	//fmt.Printf("%+v",m)
	code, ok := m["err_code"] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		judge = fmt.Sprintf("%v", code) == "0" //map和interface都不能直接比较
		if judge {
			judge = reflect.ValueOf(m["data"].(map[string]interface{})["uploadStatus"]).Bool()
			if !judge {
				msg = "UploadStatus Fail"
			}
		} else {
			judge = false
			msg = "Error"
		}
	} else {
		msg3, ok3 := m["status"] /*如果确定是真实的,则存在,否则不存在 */
		if ok3 {
			msg = fmt.Sprintf("%v", msg3)
		}
	}
	return
}
func mesOldPut(url string, token string) (judge bool, msg string) {
	judge = false
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		msg = "[Error]NewRequest"
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		msg = "[Error]Header"
		return
	}
	defer res.Body.Close()
	var r interface{}
	err = json.NewDecoder(res.Body).Decode(&r)
	m := r.(map[string]interface{})
	//fmt.Printf("%+v",m)
	code, ok := m["err_code"] /*如果确定是真实的,则存在,否则不存在 */
	if ok {
		judge = fmt.Sprintf("%v", code) == "0" //map和interface都不能直接比较
		if !judge {
			msg = fmt.Sprintf("ErrorCode %+v", code)
		}
	} else {
		msg3, ok3 := m["status"] /*如果确定是真实的,则存在,否则不存在 */
		if ok3 {
			msg = fmt.Sprintf("%v", msg3)
		}
	}
	return
}
func soapModel(wsdl,sn string) (judge bool, msg string) {
	judge=false
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(wsdl, httpClient)
	if err != nil {
		msg = fmt.Sprintf("SoapClient error: %s", err)
		return
	}
	params := gosoap.Params{
		"strSSN": sn,
	}
	res, err := soap.Call("GetSSN_CertifiedModel", params)
	if err != nil {
		log.Printf("Call error: %s", err)
		return
	}
	r:=new(RXml_Model)
	err=xml.Unmarshal(res.Body,&r)
	if err != nil {
		log.Printf("RXml_Model.Unmarshal error: %s", err)
		return
	}
	judge=r.Str[0]=="PASS"
	msg=r.Str[1]
	log.Println("[CertifiedModel]", msg)
	return
}
func soapCheck(wsdl,sn,station string) (judge bool, msg string) {
	judge=false
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(wsdl, httpClient)
	if err != nil {
		msg = fmt.Sprintf("SoapClient error: %s", err)
		return
	}
	params := gosoap.Params{
		"strSN": sn,
		"station":station,
	}
	res, err := soap.Call("CheckSSN_NEW", params)
	if err != nil {
		log.Printf("[CheckSSN_NEW]Call error: %s", err)
		return
	}
	r:=new(RXml_Check)
	err=xml.Unmarshal(res.Body,&r)
	if err != nil {
		log.Printf("RXml_Check.Unmarshal error: %s", err)
		return
	}
	judge=r.Str[0]=="PASS"
	msg=r.Str[1]
	log.Println("[CheckSSN_NEW]", msg)
	return
}
func soapReport(wsdl,sn,station string,result bool,failcode,scan string) (judge bool, msg string) {
	judge=false
	strIsPass:="FAIL"
	if result{
		strIsPass="PASS"
	}
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(wsdl, httpClient)
	if err != nil {
		msg = fmt.Sprintf("SoapClient error: %s", err)
		return
	}
	params := gosoap.Params{
		"strSSN": sn,
		"strEventPoint":station,
		"strIspass":strIsPass,
		"strFailcode":failcode,
		"strScanner":scan,
	}
	res, err := soap.Call("SaveSSN_NEW", params)
	if err != nil {
		log.Printf("Call error: %s", err)
		return
	}
	r:= new(RXml_Report)
	err=xml.Unmarshal(res.Body,&r)
	if err != nil {
		log.Printf("RXml_Report.Unmarshal error: %s", err)
		return
	}
	judge=r.Str[0]=="PASS"
	msg=r.Str[1]
	log.Println("[SaveSSN_NEW]", msg)
	return
}
func readSoapCfg()(wsdl,station string){
	cfg, err := ini.Load("MesEnable.ini")
	if err != nil {
		log.Println("Read MesEnable.ini Fail:", err)
		wsdl,station="",""
		return
	}
	wsdl = cfg.Section("Setting").Key("SoapUrl").String()
	station = cfg.Section("Setting").Key("Station").String()
	if strings.Contains(wsdl,"asmx") {
		if !strings.Contains(wsdl,"?wsdl"){
			wsdl+="?wsdl"
		}
	}
	log.Println(wsdl,station)
	return
}
func checkSoapModel(wsdl,sn string)bool{
	cfg, err := ini.Load("BurnCfgUI.ini")
	if err != nil {
		log.Println("Read BurnCfgUI.ini Fail:", err)
		return false
	}
	models := cfg.Section("CheckImg").Key("model").Strings(".")
	if len(models)==0{
		return true
	}
	ok,tagModel:=soapModel(wsdl,sn)
	log.Println(tagModel)
	if ok{
		for _, m := range models {
			if m==tagModel{

				return true
			}
		}
	}
	log.Println("Check Model Fail")
	return false
}
func checkMesModel(urlapi,token string)bool{
	cfg, err := ini.Load("BurnCfgUI.ini")
	if err != nil {
		log.Println("Read BurnCfgUI.ini Fail:", err)
		return false
	}
	models := cfg.Section("CheckImg").Key("model").Strings(".")
	if len(models)==0{
		return true
	}
	for _, m := range models {
		url:=fmt.Sprintf("%s%s",urlapi,m)
		judge,msg:=mesGet(url,token)
		log.Println(msg)
		if judge{

			return true
		}
	}
	log.Println("Check Model Fail")
	return false
}
//以下为测试用函数
func cppTest()  {
	//tokenFileName := "token gz.txt"  //pms token
	tokenFileName := "token pms.txt"  //pms token
	//tokenFileName := "token old.txt"  //old token
	if checkFileIsExist(tokenFileName) {
		bytes, _ := ioutil.ReadFile(tokenFileName)
		err:=ioutil.WriteFile("token.txt",bytes,0644)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Input SN:")
	var sn,res string
	fmt.Scanln(&sn)
	//sn:="SEI0720213100012" // jiadun
	//sn:="AAA0000000000006"		//gz
	//sn:="SEI600TID0000020"  //pms
	//sn:="QASEI530AT00010"		//old
	msg1,msg2:="xx","yy"
	fmt.Printf("%+v\r\n", apiChecksn(sn,"usid",&msg1,&msg2))
	fmt.Println(msg1,msg2)
	fmt.Println("Report Result(Enter for Pass Or any for Fail):")
	fmt.Scanln(&res)
	if res==""{
		fmt.Printf("%+v\r\n", apiReportsn(sn,true,&msg1))
	}else {
		fmt.Printf("%+v\r\n", apiReportsn(sn,false,&msg1))
	}
	fmt.Println(msg1)
	fmt.Scanln("Enter for Exit")
}
func main() {
	//cppTest()   //测试CPP
	//fmt.Println(readSoapCfg())
}
