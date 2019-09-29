package scriptUtils

import (
	i118Utils "github.com/easysoft/zentaoatf/src/utils/i118"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
	scriptUtils "github.com/easysoft/zentaoatf/src/utils/script"
)

func Sort(cases []string) {
	for _, file := range cases {
		scriptUtils.SortFile(file)
	}

	logUtils.PrintToStdOut(i118Utils.I118Prt.Sprintf("success_sort_steps", len(cases)), -1)
}
