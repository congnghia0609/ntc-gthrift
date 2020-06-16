/**
 *
 * @author nghiatc
 * @since Jun 16, 2020
 */

package main

import (
	"fmt"
	"github.com/congnghia0609/ntc-gconf/nconf"
	"path/filepath"
	"runtime"
)

func InitNConf2() {
	_, b, _, _ := runtime.Caller(0)
	wdir := filepath.Dir(b)
	fmt.Println("wdir:", wdir)
	nconf.Init(wdir)
}

func main() {
	// Init NConf
	InitNConf2()


}
