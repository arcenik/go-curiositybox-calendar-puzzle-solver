package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

const release = "0.5.0"

type DB interface {
	Close() error
}

type Globals struct {
	Debug   bool        `help:"Enable the debug mode" default:"false"`
	Version VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	Globals

	Print    PrintCmd    `cmd:"" help:"Print the solution from the database"`
	SolveAll SolveAllCmd `cmd:"" help:"Start the solver mode"`
}

func main() {
	cli := CLI{
		Globals: Globals{
			Version: VersionFlag(release),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("cbox-calendar-puzzle"),
		kong.Description("A tool for the calendar puzzle from CuriosityBox (tm)"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": release,
		})
	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
