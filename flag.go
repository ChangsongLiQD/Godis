package main

import "flag"

func GetServerFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("godis", flag.ExitOnError)

	flagSet.Bool("version", false, "print version string")
	flagSet.Int("port", DefaultPort, "server port")

	return flagSet
}
