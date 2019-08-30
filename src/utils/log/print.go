package logUtils

import (
	"encoding/json"
	"fmt"
	constant "github.com/easysoft/zentaoatf/src/utils/const"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/fatih/color"
	"io"
	"strings"
)

func PrintUsage() {
	PrintToStdOut("\nUsage: ", color.FgCyan)
	fmt.Printf("%s\n", constant.Usage)

	PrintToStdOut("Example: ", color.FgCyan)
	fmt.Printf("%s\n", constant.Example)
}

func PrintToStdOut(msg string, attr color.Attribute) {
	output := color.Output

	if attr == -1 {
		fmt.Fprintf(output, msg)
	} else {
		color.New(attr).Fprintf(output, msg)
	}
}

func PrintToCmd(msg string, attr color.Attribute) {
	var output io.Writer
	if vari.RunFromCui {
		output, _ = vari.Cui.View(constant.CuiRunOutputView)
	} else {
		output = color.Output
	}

	if attr == -1 {
		fmt.Fprintf(output, msg)
	} else {
		clr := color.New(attr)
		clr.Fprintf(output, msg)
	}
}

func PrintUnicode(str []byte) {
	var a interface{}

	temp := strings.Replace(string(str), "\\\\", "\\", -1)

	err := json.Unmarshal([]byte(temp), &a)

	var msg string
	if err == nil {
		msg = fmt.Sprint(a)
	} else {
		msg = temp
	}

	PrintToCmd(msg, -1)
}
