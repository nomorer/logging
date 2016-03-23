Package glogger provides a simple logger for Go, logs in file or console.
Example
=============
    package main
    
    import (
    	"github.com/justgolang/glogger"
    )
    
    func main() {
    	path := "/home/mckee/program/go/src/test/logs/test.log"
    	glogger.Setup(path)
    	defer glogger.Close()
    
    	glogger.Info("aaa")
    	glogger.Infof("aaa%d", 111)
    	glogger.Warn("bbb")
    	glogger.Warnf("bbb%d", 222)
    	glogger.Error("ccc")
    	glogger.Errorf("ccc%d", 333)
    	glogger.Debug("ddd")
    	glogger.Debugf("ddd%d", 444)
    	glogger.Fatal("eee")
    	glogger.Fatalf("eee%d", 555)
    }
and output like this:

    2016-03-23 17:40:18 ▶ INF aaa
    2016-03-23 17:40:18 ▶ INF aaa111
    2016-03-23 17:40:18 ▶ WAR bbb
    2016-03-23 17:40:18 ▶ WAR bbb222
    2016-03-23 17:40:18 ▶ ERR ccc
    2016-03-23 17:40:18 ▶ ERR ccc333
    2016-03-23 17:40:18 ▶ DEB ddd
    2016-03-23 17:40:18 ▶ DEB ddd444
    2016-03-23 17:40:18 ▶ FAT eee
    2016-03-23 17:40:18 ▶ FAT eee555
Install
=============
`go get github.com/justgolang/glogger`  
Use `go get -u` to update the package.  
