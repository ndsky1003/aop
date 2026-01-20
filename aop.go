package aop

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"slices"

	"github.com/samber/lo"
)

type Flag = string

const (
	FlagDefer string = "defer"
	FlagInit  string = "init"
)

type IRunAble interface {
	Run() error
}

type run_func func() error

func (this run_func) Run() error {
	return this()
}

type item struct {
	f      IRunAble
	is_run bool
	opt    *Option
}

type aop struct {
	flag      Flag
	_points   []Point
	funcs_mgr map[Point][]*item
}

func New(flag Flag) *aop {
	return &aop{
		flag:      flag,
		funcs_mgr: make(map[Point][]*item),
	}
}

func (a *aop) AddFunc(p Point, f func() error, opts ...*Option) {
	opt := Options().SetName("").SetPriority(0).Merge(opts...)
	a.add(p, run_func(f), opt)
}

func (a *aop) Add(p Point, f IRunAble, opts ...*Option) {
	opt := Options().SetName("").SetPriority(0).Merge(opts...)
	a.add(p, f, opt)
}

func (a *aop) add(p Point, f IRunAble, opt *Option) {
	a.funcs_mgr[p] = append(a.funcs_mgr[p], &item{f: f, opt: opt})
	slices.SortStableFunc(a.funcs_mgr[p], func(a, b *item) int {
		return b.opt.getPriority() - a.opt.getPriority()
	})

	if !lo.Contains(a._points, p) {
		a._points = append(a._points, p)
		slices.SortStableFunc(a._points, func(a, b Point) int {
			return int(uint16(a) - uint16(b))
		})
	}
}

func (a *aop) RunPoint(p Point) error {
	for _, item := range a.funcs_mgr[p] {
		if item.is_run {
			continue
		}
		item.is_run = true
		name := item.opt.getName()
		slog.Info(fmt.Sprintf("[AOP-%5s] %10v 执行开始", a.flag, name))
		if err := item.f.Run(); err != nil {
			slog.Error(fmt.Sprintf("[AOP-%s] %10v 执行失败", a.flag, name), "err", err)
			debug.PrintStack()
			return err
		}
		slog.Info(fmt.Sprintf("[AOP-%5s] %10v 执行成功", a.flag, name))
	}
	return nil
}

func (a *aop) Run() error {
	for _, p := range a._points {
		if err := a.RunPoint(p); err != nil {
			return err
		}
	}
	return nil
}
