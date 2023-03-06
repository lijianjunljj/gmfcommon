package loger

import (
	"log"
	"os"
	"path"
)

type Loger struct {
	path         string
	flag         int
	filenameFunc func() string
}

func NewLoger(path string, flag int, filenameFunc func() string) *Loger {
	return &Loger{
		path:         path,
		flag:         flag,
		filenameFunc: filenameFunc,
	}
}
func (l *Loger) Init() (*log.Logger, error) {
	var baseLogger *log.Logger
	if l.path != "" {
		filename := l.filenameFunc()
		file, err := os.Create(path.Join(l.path, filename))
		if err != nil {
			return nil, err
		}

		baseLogger = log.New(file, "", l.flag)
	} else {
		baseLogger = log.New(os.Stdout, "", l.flag)
	}
	return baseLogger, nil
}
