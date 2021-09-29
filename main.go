package main

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

func main() {
	svcConfig := &service.Config{
		Name:        "nws",
		DisplayName: "Nginx Windows Service",
		Description: "Nginx Windows Service.",
	}
	ns := NewNginxService()
	s, err := service.New(ns, svcConfig)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			s.Start()
			fmt.Println("服务安装成功")
			return
		}
		if os.Args[1] == "uninstall" {
			s.Stop()
			s.Uninstall()
			fmt.Println("服务卸载成功")
			return
		}
	}
	err = s.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
