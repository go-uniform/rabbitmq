package _base

import (
	"fmt"
	"github.com/go-diary/diary"
	"service/service/info"
	"strings"
)

func InitializeDiary(test bool, level string, rate int) {
	handler := diary.HumanReadableHandler
	if test {
		// test mode is used for creating an automated testing environment
		fmt.Println("entering test mode")
		handler = nil
	}

	lvl := diary.ConvertFromTextLevel(level)
	if diary.IsValidLevel(lvl) {
		panic(fmt.Sprintf("level must be one of the following values: %s", strings.Join(diary.TextLevels, ", ")))
	}
	traceRate = rate

	d = diary.Dear(AppClient, AppProject, info.AppName, nil, info.AppRepository, info.AppCommit, []string{info.AppVersion}, nil, lvl, handler)
}