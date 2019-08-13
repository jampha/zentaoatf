package testingService

import (
	"fmt"
	"github.com/easysoft/zentaoatf/src/model"
	"github.com/easysoft/zentaoatf/src/utils/common"
	constant "github.com/easysoft/zentaoatf/src/utils/const"
	"github.com/easysoft/zentaoatf/src/utils/file"
	i118Utils "github.com/easysoft/zentaoatf/src/utils/i118"
	"github.com/easysoft/zentaoatf/src/utils/log"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/fatih/color"
	"strings"
	"time"
)

func Print(report model.TestReport, workDir string) {
	startSec := time.Unix(report.StartTime, 0)
	endSec := time.Unix(report.EndTime, 0)

	logs := make([]string, 0)

	logUtils.PrintAndLog(&logs, i118Utils.I118Prt.Sprintf("run_scripts", report.Path, report.Env))

	logUtils.PrintAndLog(&logs, i118Utils.I118Prt.Sprintf("time_from_to",
		startSec.Format("2006-01-02 15:04:05"), endSec.Format("2006-01-02 15:04:05"), report.Duration))

	logUtils.PrintAndLog(&logs, fmt.Sprintf("%s: %d", i118Utils.I118Prt.Sprintf("total"), report.Total))
	logUtils.PrintAndLogColorLn(&logs, fmt.Sprintf("  %s: %d", i118Utils.I118Prt.Sprintf("pass"), report.Pass), color.FgGreen)
	logUtils.PrintAndLogColorLn(&logs, fmt.Sprintf("  %s: %d", i118Utils.I118Prt.Sprintf("fail"), report.Fail), color.FgRed)
	logUtils.PrintAndLogColorLn(&logs, fmt.Sprintf("  %s: %d", i118Utils.I118Prt.Sprintf("skip"), report.Skip), color.FgYellow)

	for _, cs := range report.Cases {
		str := "\n %s %s"
		status := cs.Status.String()
		statusColor := logUtils.ColoredStatus(status)

		logs = append(logs, fmt.Sprintf(str, status, cs.Path))
		logUtils.Printt(fmt.Sprintf(str+"\n", statusColor, cs.Path))

		if len(cs.Steps) > 0 {
			count := 0
			for _, step := range cs.Steps {
				if count > 0 { // 空行
					logUtils.PrintAndLog(&logs, "")
				}

				str := "  %s%d: %s   %s"
				status := commonUtils.BoolToPass(step.Status)
				statusColor := logUtils.ColoredStatus(status)

				logs = append(logs, fmt.Sprintf(str, i118Utils.I118Prt.Sprintf("step"), step.Numb, status, step.Name))
				logUtils.Printt(fmt.Sprintf(str, i118Utils.I118Prt.Sprintf("step"), step.Numb, statusColor, step.Name+"\n"))

				count1 := 0
				for _, cp := range step.CheckPoints {
					if count1 > 0 { // 空行
						logUtils.PrintAndLog(&logs, "")
					}

					cpStatus := commonUtils.BoolToPass(step.Status)
					cpStatusColored := logUtils.ColoredStatus(cpStatus)
					logs = append(logs, fmt.Sprintf("    %s%d: %s", i118Utils.I118Prt.Sprintf("checkpoint"), cp.Numb,
						commonUtils.BoolToPass(cp.Status)))
					logUtils.Printt(fmt.Sprintf("    %s%d: %s\n", i118Utils.I118Prt.Sprintf("checkpoint"), cp.Numb, cpStatusColored))

					logUtils.PrintAndLog(&logs, fmt.Sprintf("      %s %s", i118Utils.I118Prt.Sprintf("expect_result"), cp.Expect))
					logUtils.PrintAndLog(&logs, fmt.Sprintf("      %s %s", i118Utils.I118Prt.Sprintf("actual_result"), cp.Actual))

					count1++
				}

				count++
			}
		} else {
			logUtils.PrintAndLog(&logs, "   "+i118Utils.I118Prt.Sprintf("no_checkpoints"))
		}
	}

	fileUtils.WriteFile(workDir+constant.LogDir+vari.RunDir+"result.txt", strings.Join(logs, "\n"))
}
