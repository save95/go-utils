package locker

import (
	"errors"
	"time"
)

const (
	defaultTimeout = 5 * 60 * time.Second // 默认超时时间，5分钟有效
)

var (
	ErrorLockOccupied = errors.New("lock is occupied")
)
