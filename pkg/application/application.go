package application

import (
	"context"
	"github.com/hankeyyh/yrpc/pkg"
	"golang.org/x/sync/errgroup"
	"log"
)

type Application struct {
}

func (app *Application) Init() error {
	// 执行注册的初始化方法, 包括框架+用户自定义
	if err := pkg.DefaultRegister.Init(); err != nil {
		return err
	}
	return nil
}

func (app *Application) startListen() error {
	serverList := pkg.DefaultRegister.GetServerList()
	eg, ctx := errgroup.WithContext(context.Background())
	for _, server := range serverList {
		s := server
		eg.Go(func() error {
			return s.Listen(ctx)
		})
	}
	return eg.Wait()
}

func (app *Application) startServe() error {
	serverList := pkg.DefaultRegister.GetServerList()
	eg, ctx := errgroup.WithContext(context.Background())
	for _, server := range serverList {
		s := server
		eg.Go(func() error {
			return s.Serve(ctx)
		})
	}
	return eg.Wait()
}

func (app *Application) Start() error {
	//  注册的server（http，rpc）Listen
	if err := app.startListen(); err != nil {
		return err
	}
	// goroutine(startServers)
	if err := app.startServe(); err != nil {
		return err
	}
	// todo server.Serve 阻塞在acceptConn直到收到客户端请求
	// todo 每个收到的请求，分配一个task协程处理
	// todo 如果startServers出错，quit←err
	return nil
}

func (app *Application) Run() error {
	// todo app.Start
	// todo 注册退出信号SIGINT，SIGQUIT
	// todo 阻塞，等待退出信号←quit（异常，非异常），执行server.Stop
	// todo app.Destroy
	return nil
}

func (app *Application) Stop() error {
	// todo servers stop / gracefulStop
	// todo quit正常关闭
	return nil
}

func (app *Application) Destroy() {
	// 执行注册的Destroy方法（逆序执行），将返回的err合并打印
	if err := pkg.DefaultRegister.Destroy(); err != nil {
		log.Print(err)
	}
}
