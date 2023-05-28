package register

import (
	"github.com/hankeyyh/yrpc/internal/server"
	"go.uber.org/multierr"
	"sync"
)

type Func func() error

type Register struct {
	initFuncList    []Func // 框架和用户的初始化函数
	destroyFuncList []Func // 框架和用户的destroy函数
	serverList      []server.Server
	mu              sync.Mutex
}

func (r *Register) RegisterInitFunc(f Func) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.initFuncList = append(r.initFuncList, f)
}

func (r *Register) GetInitFuncList() []Func {
	return r.initFuncList
}

// Init 顺序执行注册的initFunc，遇到错误终止
func (r *Register) Init() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, f := range r.initFuncList {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Register) RegisterDestroyFunc(f Func) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.destroyFuncList = append(r.destroyFuncList, f)
}

func (r *Register) GetDestroyFuncList() []Func {
	return r.destroyFuncList
}

// Destroy 逆序执行destroy函数
func (r *Register) Destroy() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var errs error
	for i := len(r.destroyFuncList) - 1; i >= 0; i-- {
		f := r.destroyFuncList[i]
		if err := f(); err != nil {
			errs = multierr.Append(errs, err)
		}
	}
	return errs
}

func (r *Register) RegisterServer(s server.Server) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.serverList = append(r.serverList, s)
}

func (r *Register) GetServerList() []server.Server {
	return r.serverList
}
