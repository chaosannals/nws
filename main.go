package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/kardianos/service"
)

func IsExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// 初始化 Nginx
func InitNginx() error {
	p := "nginx-1.21.3.zip"
	name := p[:len(p)-4]

	// 判断是否有压缩包
	e, err := IsExists(p)
	if err != nil {
		return err
	}
	if !e {
		return nil
	}

	// 判断目录是否存在
	folder := "nginx"
	e, err = IsExists(folder)
	if err != nil {
		return err
	}
	if e {
		return nil
	}

	// 解压
	err = UnzipNginx(p)
	if err != nil {
		return err
	}
	err = os.Rename(name, folder)
	if err != nil {
		return err
	}
	return nil
}

// 解压 Nginx
func UnzipNginx(p string) error {
	source, err := os.Open(p)
	if err != nil {
		return err
	}
	defer source.Close()
	zipFile, err := zip.OpenReader(source.Name())
	if err != nil {
		return err
	}
	defer zipFile.Close()
	for _, innerFile := range zipFile.File {
		info := innerFile.FileInfo()
		if info.IsDir() {
			err = os.MkdirAll(innerFile.Name, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			sf, err := innerFile.Open()
			if err != nil {
				return err
			}
			defer sf.Close()
			nf, err := os.Create(innerFile.Name)
			if err != nil {
				return err
			}
			defer nf.Close()
			io.Copy(nf, sf)
		}
	}
	return nil
}

// 主入口
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
		return
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err = InitNginx()
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
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
