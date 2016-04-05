Package glogger provides a simple logger for Go, logs in file or console.
Example
=============
    package main
    
    import (
    	"github.com/justgolang/glogger"
    	"fmt"
    )
    
    func main() {
    	path := "/home/mckee/program/go/src/test/logs/test.log"
    	if err := glogger.Setup(path, glogger.LevelDebug); err != nil {
    		fmt.Print(err)
    		return
    	}
    	defer glogger.Close()
    
    	glogger.Debug("ddd")
    	glogger.Debugf("ddd%d", 444)
    	glogger.Info("aaa")
    	glogger.Infof("aaa%d", 111)
    
    	//will no show
    	glogger.SetLevel(glogger.LevelError)
    	glogger.Warn("bbb")
    	glogger.Warnf("bbb%d", 222)
    
    	//will show
    	glogger.Error("ccc")
    	glogger.Errorf("ccc%d", 333)
    	glogger.Fatal("eee")
    	glogger.Fatalf("eee%d", 555)
    }

and output like this:

    2016-04-05 19:42:35 ▶ DEB ddd
    2016-04-05 19:42:35 ▶ DEB ddd444
    2016-04-05 19:42:35 ▶ INF aaa
    2016-04-05 19:42:35 ▶ INF aaa111
    2016-04-05 19:42:35 ▶ ERR ccc
    2016-04-05 19:42:35 ▶ ERR ccc333
    2016-04-05 19:42:35 ▶ FAT eee
    2016-04-05 19:42:35 ▶ FAT eee555
Install
=============
`go get github.com/justgolang/glogger`  
use `go get -u` to update the package.  
