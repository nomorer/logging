Package glogger provides a simple logger with level and log rotation for Go.
Example
=============
    package main
    
    import (
    	"fmt"
    	log "github.com/justgolang/glogger"
    )
    
    func main() {
    	path := "/home/mckee/program/go/src/test/logs/test.log"
    	if err := log.Setup(path, log.LevelDebug, log.HourlyRotate); err != nil {
    		fmt.Print(err)
    		return
    	}
    	defer func() {
    		log.Close()
    	}()
    
    	log.Debug("ddd")
    	log.Debugf("ddd%d", 444)
    	log.Info("aaa")
    	log.Infof("aaa%d", 111)
    
    	//will no show
    	log.SetLevel(log.LevelError)
    	log.Warn("bbb")
    	log.Warnf("bbb%d", 222)
    
    	//will show
    	log.Error("ccc")
    	log.Errorf("ccc%d", 333)
    	log.Fatal("eee")
    
    	//will not show
    	log.Fatalf("eee%d", 555)
    }


and output like this:

    2016-04-05 19:42:35 ▶ DEB ddd
    2016-04-05 19:42:35 ▶ DEB ddd444
    2016-04-05 19:42:35 ▶ INF aaa
    2016-04-05 19:42:35 ▶ INF aaa111
    2016-04-05 19:42:35 ▶ ERR ccc
    2016-04-05 19:42:35 ▶ ERR ccc333
    2016-04-05 19:42:35 ▶ FAT eee
Install
=============
`go get github.com/justgolang/glogger`  
use `go get -u` to update the package.  
