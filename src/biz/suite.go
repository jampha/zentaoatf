package biz

import (
	"github.com/easysoft/zentaoatf/src/utils"
	"strings"
)

func GenSuite(cases []string) {
	str := strings.Join(cases, "\n")

	utils.WriteFile(utils.Prefer.WorkDir+utils.ScriptDir+"all."+utils.SuiteExt, str)
}