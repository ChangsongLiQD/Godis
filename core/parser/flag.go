package parser

import "flag"

const (
	DefaultPort = 6666
)

func GetFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("godis", flag.ExitOnError)

	flagSet.Bool("version", false, "print version string")
	flagSet.Int("port", DefaultPort, "server port")

	return flagSet
}
