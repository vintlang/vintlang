package bundler
// package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/vintlang/vintlang/repl"
// )

// var BundlerVersion = "{{.BundlerVersion}}"
// var BuildTime = "{{.BuildTime}}"

// func main() {
// 	bundledDetails := flag.Bool("-i", false, "Show the bundle details of the app")
// 	if *bundledDetails {
// 		fmt.Printf("[Bundler Version: %s | Build Time: %s]\n", BundlerVersion, BuildTime)
// 	}
// 	code := ` + "`{{.Code}}`" + `
// 	repl.Read(code)
// }
