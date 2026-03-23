package core

import (
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/httpclient"
)

type ScanContext struct {
	Target    string
	Timeout   time.Duration
	UserAgent string
	Client    *httpclient.Client
}
