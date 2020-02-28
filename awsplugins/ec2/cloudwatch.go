package ec2

import (
	"encoding/json"
	"github.com/aws/aws-xray-sdk-go/internal/logger"
	"github.com/aws/aws-xray-sdk-go/internal/plugins"
	"os"
)

// {"version":"1","log_configs":[{"log_group_name":"foo"}],"region":"eu-central-1"}
type logConfig struct {
	LogConfig []struct{
		LogGroupName string `json:"log_group_name"`
	} `json:"log_config"`
}


func getLogGroupNames(filename string) ([]string, error) {
	f, e := os.Open(filename)

	if e != nil {
		return nil, e
	}

	defer f.Close()

	lc := logConfig{}
	e = json.NewDecoder(f).Decode(&lc)

	if e != nil {
		return nil, e
	}

	var names []string

	for _, c := range lc.LogConfig {
		names = append(names, c.LogGroupName)
	}

	return names, nil
}

func getLogReferences() []plugins.CloudwatchLogsMetadata {
	names, e := getLogGroupNames(logConfigPath)

	if e != nil {
		logger.Errorf("Unable to get log group names: %v", e)
		return nil
	}

	var logReferences []plugins.CloudwatchLogsMetadata
	for _, name := range names {
		logReferences = append(logReferences, plugins.CloudwatchLogsMetadata{name })
	}

	return logReferences
}
