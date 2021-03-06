package main

import (
	"flag"
	"github.com/easysoft/zentaoatf/src/action"
	configUtils "github.com/easysoft/zentaoatf/src/utils/config"
	fileUtils "github.com/easysoft/zentaoatf/src/utils/file"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"
)

var (
	language        string
	independentFile bool
	keywords        string

	productId string
	moduleId  string
	taskId    string
	suiteId   string

	flagSet *flag.FlagSet
)

func main() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		cleanup()
		os.Exit(0)
	}()

	flagSet = flag.NewFlagSet("atf", flag.ContinueOnError)

	flagSet.StringVar(&productId, "p", "", "")
	flagSet.StringVar(&productId, "product", "", "")

	flagSet.StringVar(&moduleId, "m", "", "")
	flagSet.StringVar(&moduleId, "module", "", "")

	flagSet.StringVar(&suiteId, "s", "", "")
	flagSet.StringVar(&suiteId, "suiteId", "", "")

	flagSet.StringVar(&taskId, "t", "", "")
	flagSet.StringVar(&taskId, "taskId", "", "")

	flagSet.StringVar(&language, "l", "", "")
	flagSet.StringVar(&language, "language", "", "")

	flagSet.BoolVar(&independentFile, "i", false, "")
	flagSet.BoolVar(&independentFile, "independent", false, "")

	flagSet.StringVar(&keywords, "k", "", "")
	flagSet.StringVar(&keywords, "keywords", "", "")

	flagSet.BoolVar(&vari.Verbose, "verbose", false, "")

	var placeholder string
	flagSet.StringVar(&placeholder, "h", "", "")
	flagSet.StringVar(&placeholder, "r", "", "")
	flagSet.StringVar(&placeholder, "v", "", "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "run", ".")
	}

	switch os.Args[1] {
	case "run", "-r":
		run(os.Args)

	case "checkout", "co":
		if err := flagSet.Parse(os.Args[2:]); err == nil {
			action.Generate(productId, moduleId, suiteId, taskId, independentFile, language)
		}

	case "update", "up":
		if err := flagSet.Parse(os.Args[2:]); err == nil {
			action.Generate(productId, moduleId, suiteId, taskId, independentFile, language)
		}

	case "ci":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.CommitCases(files)
		}

	case "cr":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.CommitResult(files)
		}

	case "cb":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.CommitBug(files)
		}

	case "list", "ls", "-l":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			if len(files) == 0 {
				files = append(files, ".")
			}

			action.List(files, keywords)
		}

	case "view", "-v":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.View(files, keywords)
		}

	case "set", "-set":
		action.Set()

	case "sort", "-sort":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.Sort(files)
		}

	case "clean", "-clean", "-c":
		action.Clean()

	case "help", "-h":
		logUtils.PrintUsage()

	default: // run
		if len(os.Args) > 1 {
			args := []string{os.Args[0], "run"}
			args = append(args, os.Args[1:]...)

			run(args)
		} else {
			logUtils.PrintUsage()
		}
	}
}

func run(args []string) {
	files := fileUtils.GetFilesFromParams(args[2:])

	err := flagSet.Parse(args[len(files)+2:])
	if err == nil {
		if len(files) == 0 {
			files = append(files, ".")
		}

		action.Run(files, suiteId, taskId)
	} else {
		logUtils.PrintUsage()
	}
}

func init() {
	cleanup()

	vari.RunFromCui = false
	configUtils.InitConfig()
}

func cleanup() {
	color.Unset()
}
