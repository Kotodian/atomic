package fileout

import (
	"bytes"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"time"
)

type WriterSyncer interface {
	io.Writer
	Sync() error
}

func NewRollingFile(path, srvName string, maxSize, MaxAge int) WriterSyncer {
	if err := os.MkdirAll(path, 0766); err != nil {
		panic(err)
	}
	return newLumberjackWriteSyncer(&lumberjack.Logger{
		Filename:  filepath.Join(path, srvName+".log"),
		MaxSize:   maxSize,
		MaxAge:    MaxAge,
		LocalTime: true,
		Compress:  false,
	})
}

type lumberjackWriteSyncer struct {
	*lumberjack.Logger
	buf       *bytes.Buffer
	logChan   chan []byte
	closeChan chan interface{}
	maxSize   int
}

func (l *lumberjackWriteSyncer) Sync() error {
	return nil
}

func newLumberjackWriteSyncer(l *lumberjack.Logger) *lumberjackWriteSyncer {
	ws := &lumberjackWriteSyncer{
		Logger:    l,
		buf:       bytes.NewBuffer([]byte{}),
		logChan:   make(chan []byte, 5000),
		closeChan: make(chan interface{}),
		maxSize:   1024,
	}
	go ws.run()
	return ws
}

func (l *lumberjackWriteSyncer) run() {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			if l.buf.Len() > 0 {
				l.sync()
			}
		case bs := <-l.logChan:
			_, err := l.buf.Write(bs)
			if err != nil {
				continue
			}

			if l.buf.Len() > l.maxSize {
				l.sync()
			}
		case <-l.closeChan:
			l.sync()
			return
		}
	}
}

func (l *lumberjackWriteSyncer) sync() error {
	defer l.buf.Reset()
	_, err := l.Logger.Write(l.buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (l *lumberjackWriteSyncer) Stop() {
	close(l.closeChan)
}
