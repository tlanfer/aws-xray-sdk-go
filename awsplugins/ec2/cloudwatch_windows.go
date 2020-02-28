package ec2

import "os"

var logConfigPath = os.Getenv("ProgramData") + "\\Amazon\\AmazonCloudWatchAgent\\log-config.json"