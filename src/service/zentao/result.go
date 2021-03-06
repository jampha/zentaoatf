package zentaoService

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/easysoft/zentaoatf/src/service/client"
	testingService "github.com/easysoft/zentaoatf/src/service/testing"
	configUtils "github.com/easysoft/zentaoatf/src/utils/config"
	constant "github.com/easysoft/zentaoatf/src/utils/const"
	i118Utils "github.com/easysoft/zentaoatf/src/utils/i118"
	"github.com/easysoft/zentaoatf/src/utils/log"
	stdinUtils "github.com/easysoft/zentaoatf/src/utils/stdin"
	stringUtils "github.com/easysoft/zentaoatf/src/utils/string"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/easysoft/zentaoatf/src/utils/zentao"
	"strconv"
)

func CommitResult(resultDir string) {
	conf := configUtils.ReadCurrConfig()
	Login(conf.Url, conf.Account, conf.Password)

	report := testingService.GetTestTestReportForSubmit(resultDir)

	for _, cs := range report.Cases {
		id := cs.Id
		title := cs.Title
		status := cs.Status

		confirm := true

		tips := fmt.Sprintf("%d. %s %s", id, title, stringUtils.Ucfirst(status))
		stdinUtils.InputForBool(&confirm, confirm, "confirm_commit_result", tips)

		requestObj := map[string]interface{}{"case": strconv.Itoa(id), "version": "0"}

		stepMap := map[string]string{}
		realMap := map[string]string{}
		for _, step := range cs.Steps {
			var stepStatus string
			if step.Status {
				stepStatus = constant.PASS.String()
			} else {
				stepStatus = constant.FAIL.String()
			}

			stepResults := ""
			for _, checkpoint := range step.CheckPoints {
				stepResults += checkpoint.Actual // strconv.FormatBool(checkpoint.Status) + ": " + checkpoint.Actual
			}
			stepMap[step.Id] = stepStatus
			realMap[step.Id] = stepResults

			requestObj["steps"] = stepMap
			requestObj["reals"] = realMap
		}

		uri := fmt.Sprintf("testtask-runCase-0-%d-1.json", id)
		url := conf.Url + uri

		_, ok := client.PostObject(url, requestObj)
		if ok {
			logUtils.PrintTo(i118Utils.I118Prt.Sprintf("success_to_commit_result", id))

			if vari.Verbose {
				resultId := GetLastResult(conf.Url, id)
				logUtils.PrintTo(fmt.Sprintf("returned result id = %d", resultId))
			}
		}
	}
}

func GetLastResult(baseUrl string, caseId int) int {
	params := fmt.Sprintf("0-%d-1.json", caseId)

	url := baseUrl + zentaoUtils.GenApiUri("testtask", "results", params)
	dataStr, ok := client.Get(url, nil)

	resultId := -1
	if ok {
		jsonData, err := simplejson.NewJson([]byte(dataStr))
		if err == nil {
			results, _ := jsonData.Get("results").Map()

			for key, _ := range results {
				numb, _ := strconv.Atoi(key)

				if resultId < numb {
					resultId = numb
				}
			}
		}
	}

	return resultId
}
