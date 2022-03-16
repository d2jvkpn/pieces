package misc

import (
	"runtime/debug"
	"strings"
	"time"
)

func GetBuildInfo(buildVersion string) (info map[string]string) {
	info = make(map[string]string, 10)
	buildInfo, _ := debug.ReadBuildInfo()

	info["goVersion"] = buildInfo.GoVersion
	info["startTime"] = time.Now().Format(time.RFC3339)

	if buildVersion != "" {
		info["buildVersion"] = buildVersion
	}

	parseFlags := func(str string) {
		for _, v := range strings.Fields(str) {
			k, v, _ := strings.Cut(v, "=")
			if strings.HasPrefix(k, "main.") && v != "" {
				info[k[5:]] = v
			}
		}
	}

	for _, v := range buildInfo.Settings {
		if v.Key == "-ldflags" || v.Key == "--ldflags" {
			parseFlags(v.Value)
			break
		}
	}

	return info
}
