package main

import (
	"fmt"
	"time"
	
	"github.com/apex/log"
)

type Log struct {
	*log.Entry
}

const FORMAT = time.RFC3339

// Debugf level formatted message.
func (l *Log)  Debugf(msg string, v ...interface{}) {
	n := time.Now()
	l.Debug(fmt.Sprintf(n.Format(FORMAT) + ":  " + msg, v...))
}

// Infof level formatted message.
func (l *Log) Infof(msg string, v ...interface{}) {
	n := time.Now()
	l.Info(fmt.Sprintf(n.Format(FORMAT) + ":  " + msg, v...))
}

// Warnf level formatted message.
func (l *Log)  Warnf(msg string, v ...interface{}) {
	n := time.Now()
	l.Warn(fmt.Sprintf(n.Format(FORMAT) + ":  " + msg, v...))
}

// Errorf level formatted message.
func (l *Log)  Errorf(msg string, v ...interface{}) {
	n := time.Now()
	l.Error(fmt.Sprintf(n.Format(FORMAT) + ":  " + msg, v...))
}

// Fatalf level formatted message, followed by an exit.
func (l *Log)  Fatalf(msg string, v ...interface{}) {
	n := time.Now()
	l.Fatal(fmt.Sprintf(n.Format(FORMAT) + ":  " + msg, v...))
}