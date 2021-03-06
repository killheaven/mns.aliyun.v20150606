package mns

import (
	"net/http"
	"time"
)

type Config struct {
	AccessKeyId     string
	AccessKeySecret string

	// following is optional
	Timeout       time.Duration
	Base64Enabled bool
	HttpClient    *http.Client
}
