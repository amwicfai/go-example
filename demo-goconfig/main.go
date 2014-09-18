package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		fmt.Println("读取配置文件失败[config.ini]")
		return
	}

	v, _ := cfg.GetValue("Demo", "key1")
	fmt.Println("GetValue: ", v)

	age := cfg.MustInt("parent", "age")
	fmt.Println("MustInt: ", age)

	//借助 Go 语言变参的功能，当 Must 系列方法拥有三个参数时，则第三个参数即为获取失败时的默认值
	age = cfg.MustInt("parent", "age", 100)
	fmt.Println("MustInt and default: ", age)

	money := cfg.MustFloat64("parent", "money", 200)
	fmt.Println("MustFloat64 and default: ", money)

	google, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "google")
	fmt.Println("default section: ", google)

	vInt, err := cfg.Int("parent", "age")
	if err != nil {
		fmt.Println("Int: ", err)
		return
	}
	fmt.Println("Int: ", vInt)
}
