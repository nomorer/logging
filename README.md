package logging provides a simple logger with level and log rotation for Go.
Getting Started
=============
    package main
    
    import (
    	"fmt"
    	log "github.com/justgolang/logging"
    	"time"
    )
    
    func main() {
    	path := "/home/mckee/program/go/src/test/logs/access.log"
    	if err := log.Setup(path, log.LevelDebug, log.MinutelyRotate); err != nil {
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
    
    	for i := 0; i < 100; i++ {
    		log.Infof("log at: %d", i)
    		time.Sleep(1000 * 1000 * 1000)
    	}
    
    	//will no show
    	log.SetLevel(log.LevelError)
    	log.Warn("bbb")
    	log.Warnf("bbb%d", 222)
    
    	//will show
    	log.Error("ccc")
    	log.Errorf("ccc%d", 333)
    
    	//will exit
    	log.Fatal("eee")
    
    	//will not show
    	log.Fatalf("eee%d", 555)
    }

By executing the code above, that is rotate minutely, you will get output like this:

    2016-04-10 22:41:31 ▶ DEB ddd
    2016-04-10 22:41:31 ▶ DEB ddd444
    2016-04-10 22:41:31 ▶ INF aaa
    2016-04-10 22:41:31 ▶ INF aaa111
    ......
    2016-04-10 22:43:10 ▶ INF log at: 99
    2016-04-10 22:43:11 ▶ ERR ccc
    2016-04-10 22:43:11 ▶ ERR ccc333
    2016-04-10 22:43:11 ▶ FAT eee
and will generate at least two files: `access.log` as current log file, `access.log-2006-01-02_15-04` as backup log file.
Features
=============
* Logging in File with rotation
* Logging with levels
* Enable/disable Logger
* Automatic word wrapping

Levels
=============
*   Debug
*   Info
*   Warn
*   Error
*   Fatal (in this case program will exit)

RotateRules
=============
*	MonthlyRotate (rotate the logs monthly)
*	DailyRotate (rotate the logs daily)
*	HourlyRotate (rotate the logs hourly)
*	MinutelyRotate (rotate the logs minutely)

Installation
=============
`go get github.com/justgolang/logging`  
use `go get -u` to update the package.  