package client

import (
	"encoding/json"
	"fmt"
	"github.com/ajg/form"
	"github.com/easysoft/zentaoatf/src/model"
	"github.com/easysoft/zentaoatf/src/utils/log"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Get(url string, params map[string]string) (string, bool) {
	if vari.Verbose {
		logUtils.PrintToCmd(url, -1)
	}
	client := &http.Client{}

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		if vari.Verbose {
			logUtils.PrintToCmd(reqErr.Error(), color.FgRed)
		}
		return "", false
	}

	req.Header.Set("cookie", vari.SessionVar+"="+vari.SessionId)

	q := req.URL.Query()
	q.Add(vari.SessionVar, vari.SessionId)
	if params != nil {
		for pkey, pval := range params {
			q.Add(pkey, pval)
		}

	}
	req.URL.RawQuery = q.Encode()

	resp, respErr := client.Do(req)
	if respErr != nil {
		if vari.Verbose {
			logUtils.PrintToCmd(respErr.Error(), color.FgRed)
		}
		return "", false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if vari.Verbose {
		logUtils.PrintUnicode(bodyStr)
	}

	var bodyJson model.ZentaoResponse
	jsonErr := json.Unmarshal(bodyStr, &bodyJson)
	if jsonErr != nil {
		if strings.Index(string(bodyStr), "<html>") > -1 {
			if vari.Verbose {
				logUtils.PrintToCmd("server return a html", -1)
			}
			return "", false
		} else {
			if vari.Verbose {
				logUtils.PrintToCmd(jsonErr.Error(), color.FgRed)
			}
			return "", false
		}
	}

	defer resp.Body.Close()

	status := bodyJson.Status
	if status == "" { // 非嵌套结构
		return string(bodyStr), true
	} else { // 嵌套结构
		dataStr := bodyJson.Data
		return dataStr, status == "success"
	}
}

func PostObject(url string, params interface{}) (string, bool) {
	if vari.Verbose {
		logUtils.PrintToCmd(url, -1)
		logUtils.PrintToCmd(fmt.Sprintf("%+v", params), -1)
	}

	client := &http.Client{}

	val, _ := form.EncodeToString(params)

	// convert data to post fomat
	re3, _ := regexp.Compile(`([^&]*?)=`)
	data := re3.ReplaceAllStringFunc(val, replacePostData)

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(data))
	if reqErr != nil {
		if vari.Verbose {
			logUtils.PrintToCmd(reqErr.Error(), color.FgRed)
		}
		return "", false
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cookie", vari.SessionVar+"="+vari.SessionId)

	resp, respErr := client.Do(req)
	if respErr != nil {
		if vari.Verbose {
			logUtils.PrintToCmd(respErr.Error(), color.FgRed)
		}
		return "", false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if vari.Verbose {
		logUtils.PrintUnicode(bodyStr)
	}

	var bodyJson model.ZentaoResponse
	jsonErr := json.Unmarshal(bodyStr, &bodyJson)
	if jsonErr != nil {
		if strings.Index(string(bodyStr), "<html>") > -1 { // some api return a html
			if vari.Verbose {
				logUtils.PrintToCmd("server return a html", -1)
			}
			return "", true
		} else {
			if vari.Verbose {
				logUtils.PrintToCmd(jsonErr.Error(), color.FgRed)
			}
			return "", false
		}
	}

	defer resp.Body.Close()

	status := bodyJson.Status
	if status == "" { // 非嵌套结构
		return string(bodyStr), true
	} else { // 嵌套结构
		dataStr := bodyJson.Data
		return dataStr, status == "success"
	}
}

func PostStr(url string, params map[string]string) (string, bool) {
	if vari.Verbose {
		logUtils.PrintToCmd(url, -1)
	}
	client := &http.Client{}

	paramStr := ""
	idx := 0
	for pkey, pval := range params {
		if idx > 0 {
			paramStr += "&"
		}
		paramStr = paramStr + pkey + "=" + pval
		idx++
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(paramStr))
	if reqErr != nil {
		if vari.Verbose {
			logUtils.PrintToCmd(reqErr.Error(), color.FgRed)
		}
		return "", false
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cookie", vari.SessionVar+"="+vari.SessionId)

	resp, respErr := client.Do(req)
	if respErr != nil {
		if vari.Verbose {
			logUtils.PrintToCmd(respErr.Error(), color.FgRed)
		}
		return "", false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if vari.Verbose {
		logUtils.PrintUnicode(bodyStr)
	}

	var bodyJson model.ZentaoResponse
	jsonErr := json.Unmarshal(bodyStr, &bodyJson)
	if jsonErr != nil && strings.Index(url, "login") == -1 { // ignore login request which return a html
		if vari.Verbose {
			logUtils.PrintToCmd(jsonErr.Error(), color.FgRed)
		}
		return "", false
	}

	defer resp.Body.Close()

	status := bodyJson.Status
	if status == "" { // 非嵌套结构
		return string(bodyStr), true
	} else { // 嵌套结构
		dataStr := bodyJson.Data
		return dataStr, status == "success"
	}
}

func replacePostData(str string) string {
	str = strings.ToLower(str[:1]) + str[1:]

	if strings.Index(str, ".") > -1 {
		str = strings.Replace(str, ".", "[", -1)
		str = strings.Replace(str, "=", "]=", -1)
	}
	return str
}
