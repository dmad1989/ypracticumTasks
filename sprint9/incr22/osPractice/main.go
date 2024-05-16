package main

import (
	"fmt"
	"os"
)

func main() {
	ex, _ := os.Executable()
	host, _ := os.Hostname()
	cache, _ := os.UserCacheDir()
	config, _ := os.UserConfigDir()
	home, _ := os.UserHomeDir()

	fmt.Printf(`
Args = %v
Executable = %s
Hostname = %s
TempDir = %s
CacheDir = %s
ConfigDir = %s
HomeDir = %s
`, os.Args, ex, host, os.TempDir(), cache, config, home)
}
