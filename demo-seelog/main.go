package main

import common "newegg.com/seelog/common"

func main() {
	log := common.Logger
	defer log.Flush()
	for i := 0; i < 100000; i++ {
		log.Info("Hello from Seelog!")
		log.Error("this is error")
	}

}
