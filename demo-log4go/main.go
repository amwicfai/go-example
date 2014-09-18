package main

import (
	l4g "code.google.com/p/log4go"
	"time"
)

func main() {
	log := l4g.NewDefaultLogger(l4g.DEBUG)

	//输出到控制台,级别为DEBUG
	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())

	//logs文件夹需要运行前创建，否则异常
	//输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
	log.AddFilter("log", l4g.DEBUG, l4g.NewFileLogWriter("logs/test.log", false))

	//使用加载配置文件,类似与java的log4j.propertites
	//l4g.LoadConfiguration("log4go.xml")

	//l4g.Info("xxx")
	//log.Error("yyy")
	log.Debug("the time is now :%s -- %s", "213", "sad")

	//程序退出太快是来不及写log的
	time.Sleep(1000 * time.Millisecond)
	//log.Write2Exit()

	//注:如果不是一直运行的程序,请加上这句话,否则主线程结束后,也不会输出和log到日志文件
	defer log.Close()
}
