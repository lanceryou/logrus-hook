package logrus_hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type StackHook struct {
	skip int
	mu   sync.Mutex
}

func NewStackHook(skip int) *StackHook {
	return &StackHook{skip: skip}
}

func (s *StackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (s *StackHook) Fire(entry *logrus.Entry) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var funcName string
	pc, filename, line, ok := runtime.Caller(s.skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcName = filepath.Ext(funcName)            // .foo
		funcName = strings.TrimPrefix(funcName, ".") // foo
		filename = filepath.Base(filename)
	}

	entry.Data["caller"] = fmt.Sprintf("%v:%v:%v ", filename, funcName, line)
	return nil
}
