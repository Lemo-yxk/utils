// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-22 17:44
**/

package utils

import (
	"os"
	"os/signal"
	"syscall"
)

type sig int

const Signal sig = iota

type done struct {
	fn func(func(sig os.Signal))
}

func (d *done) Done(fn func(sig os.Signal)) {
	d.fn(fn)
}

// listen all signal
func (s sig) ListenAll() *done {
	var signalList = []os.Signal{
		syscall.SIGABRT, syscall.SIGALRM, syscall.SIGBUS, syscall.SIGCHLD, syscall.SIGCONT,
		syscall.SIGFPE, syscall.SIGHUP, syscall.SIGILL, syscall.SIGINT, syscall.SIGIO,
		syscall.SIGIOT, syscall.SIGKILL, syscall.SIGPIPE, syscall.SIGPROF, syscall.SIGQUIT,
		syscall.SIGSEGV, syscall.SIGSTOP, syscall.SIGSYS, syscall.SIGTERM, syscall.SIGTRAP,
		syscall.SIGTSTP, syscall.SIGTTIN, syscall.SIGTTOU, syscall.SIGURG, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGVTALRM, syscall.SIGWINCH, syscall.SIGXCPU, syscall.SIGXFSZ,
	}
	// 创建信号
	signalChan := make(chan os.Signal, 1)
	// 通知
	signal.Notify(signalChan, signalList...)
	return &done{fn: func(f func(signal os.Signal)) {
		// 阻塞
		f(<-signalChan)
		// 停止
		signal.Stop(signalChan)
	}}
}

func (s sig) ListenKill() *done {
	var signalList = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	// 创建信号
	signalChan := make(chan os.Signal, 1)
	// 通知
	signal.Notify(signalChan, signalList...)
	return &done{fn: func(f func(signal os.Signal)) {
		// 阻塞
		f(<-signalChan)
		// 停止
		signal.Stop(signalChan)
	}}
}

func (s sig) Listen(sig ...os.Signal) *done {
	var signalList = sig
	// 创建信号
	signalChan := make(chan os.Signal, 1)
	// 通知
	signal.Notify(signalChan, signalList...)
	return &done{fn: func(f func(signal os.Signal)) {
		// 阻塞
		f(<-signalChan)
		// 停止
		signal.Stop(signalChan)
	}}
}

func (s sig) Signal(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}

func (s sig) KillGroup(pid int, sig syscall.Signal) error {
	return syscall.Kill(-pid, sig)
}
