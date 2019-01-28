package coat

import (
	"log"
	"os"
	"time"

	"github.com/bclicn/color"
)

// Console : interface to formated terminal output
type Console struct {
	Log func(v ...interface{})
	Err func(v ...interface{})
}

// NewConsole returns a formated output console
func NewConsole(name string, silence bool) *Console {
	_err := log.New(os.Stderr, "", 0)
	_log := log.New(os.Stdout, "", 0)
	if silence {
		return &Console{
			Log: func(v ...interface{}) {},
			Err: func(v ...interface{}) {}}
	}

	return &Console{
		Log: func(v ...interface{}) {
			_log.SetPrefix(color.BBlue("["+name+"]~[") +
				color.BPurple(time.Now().Format("2006-01-02 15:04:05.000000")) +
				color.BBlue("]~"))
			_log.Println(v)
		},
		Err: func(v ...interface{}) {
			_err.SetPrefix(color.BRed("["+name+"]~[") +
				color.BPurple(time.Now().Format("2006-01-02 15:04:05.000000")) +
				color.BRed("]~"))
			_err.Println(v)
		}}
}
