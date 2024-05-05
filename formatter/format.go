package formatter

import (
	"github.com/mylxsw/asteria/event"
)

// Formatter 日志格式化接口
type Formatter interface {
	// Format 日志格式化
	Format(f event.Event) string
	//FormatI(f event.Event) string
	FormatIndent(f event.Event, prefix, indent string) string
}
