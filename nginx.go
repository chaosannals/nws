package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cihub/seelog"
	"github.com/kardianos/service"
)

//NginxService 服务
type NginxService struct {
	root   string
	logger seelog.LoggerInterface
}

//NewNginxService 创建
func NewNginxService() *NginxService {
	// 获取程序所在路径
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("当前路径: %v\n", root)
	if err != nil {
		return nil
	}
	// 设置当前工作路径为 Nginx 程序所在路径
	err = os.Chdir(root)
	if err != nil {
		return nil
	}
	// 加载设置。
	path := filepath.Join(root, "seelog.xml")
	logger, err := seelog.LoggerFromConfigAsFile(path)
	if err != nil {
		fmt.Printf("seelog.xml 失败: %v\n", err)
		return nil
	}
	seelog.ReplaceLogger(logger)
	return &NginxService{
		root:   root,
		logger: logger,
	}
}

//Start 开始
func (p *NginxService) Start(s service.Service) error {
	p.logger.Info("服务启动")

	// 设置工作目录为 nginx 目录
	wkdir := filepath.Join(p.root, "nginx")
	e, err := IsExists(wkdir)
	if err != nil {
		return err
	}
	if e {
		err := os.Chdir(wkdir)
		if err != nil {
			return err
		}
	}

	// 启动 Nginx
	cmd := exec.Command("nginx")
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

//Stop 停止
func (p *NginxService) Stop(s service.Service) error {
	p.logger.Info("服务关闭")
	// 设置工作目录为 nginx 目录
	wkdir := filepath.Join(p.root, "nginx")
	e, err := IsExists(wkdir)
	if err != nil {
		return err
	}
	if e {
		err := os.Chdir(wkdir)
		if err != nil {
			return err
		}
	}
	// 关闭 Nginx 命令
	cmd := exec.Command("nginx", "-s", "stop")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
