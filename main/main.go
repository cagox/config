package main


import (
	"fmt"
	"github.com/cagox/config"
)


type MyConf  struct {
	config.ConfigurationStruct
	MyTestValue string
}

func main() {
	myConf := MyConf{}

	fmt.Println("Stuff")
	config.LoadConfigs(&myConf, "OCFCONF")

	fmt.Println("Host: " + myConf.Host)
	fmt.Println("Name Length: ", myConf.MinimumNameLength)
	fmt.Println("CookieName: " + myConf.CookieName)
	fmt.Println("TestValue: " + myConf.MyTestValue)
}