package logrus_hook

import (
	"testing"
	"time"
)

func TestRotateFile_Filter(t *testing.T) {
	type Filter struct {
		Files    []string
		BackTime time.Duration
		CurFn    string
		Expected []string
	}

	now := time.Now()
	filters := []Filter{
		{
			Files:    []string{"info.log", genFileName("info.log", now.Add(-1*time.Minute)), genFileName("info.log", now.Add(-1*time.Hour))},
			BackTime: time.Minute * 10,
			CurFn:    "info.log",
			Expected: []string{genFileName("info.log", now.Add(-1*time.Hour))},
		},
		{
			Files:    []string{"./info.log", genFileName("./info.log", now.Add(-1*time.Minute)), genFileName("./info.log", now.Add(-1*time.Hour))},
			BackTime: time.Minute * 10,
			CurFn:    "./info.log",
			Expected: []string{genFileName("./info.log", now.Add(-1*time.Hour))},
		},
	}

	for _, f := range filters {
		if !equalArray(filterBackFiles(f.Files, f.CurFn, f.BackTime), f.Expected) {
			t.Errorf("filter error expected %v ret %v", f.Expected, filterBackFiles(f.Files, f.CurFn, f.BackTime))
		}
	}
}

func equalArray(l []string, r []string) bool {
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}

	return len(l) == len(r)
}
