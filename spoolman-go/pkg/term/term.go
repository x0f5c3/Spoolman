package term

import (
	"github.com/pterm/pterm"
	"os"
	"sync"
)

type Level byte

const (
	Debug Level = iota
	Description
	Info
	Success
	Warning
	Error
	Fatal
)

var levelToPrinter = map[Level]pterm.PrefixPrinter{
	Debug:       pterm.Debug,
	Description: pterm.Description,
	Info:        pterm.Info,
	Success:     pterm.Success,
	Warning:     pterm.Warning,
	Error:       pterm.Error,
	Fatal:       pterm.Fatal,
}

var Verbose bool

var initOnce = sync.OnceFunc(initPterm)

func initPterm() {
	if Verbose {
		pterm.EnableOutput()
		pterm.EnableColor()
		pterm.EnableStyling()
		pterm.EnableDebugMessages()
		pterm.SetDefaultOutput(os.Stdout)
	}
}

func Printfln(lvl Level, msg string, args ...any) {
	initOnce()
	if !Verbose {
		return
	}
	printer, ok := levelToPrinter[lvl]
	if !ok {
		return
	}
	printer.Printfln(msg, args...)
}

func GetPrinter(lvl Level) pterm.PrefixPrinter {
	initOnce()
	return levelToPrinter[lvl]
}
