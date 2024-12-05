package aop

import (
	"fmt"
	"sort"
)

type IRunAble interface {
	Run() error
}

type run_func func() error

func (this run_func) Run() error {
	return this()
}

type point uint8

type item struct {
	f      IRunAble
	is_run bool
	opt    *Option
}

type aop struct {
	_points   []point
	funcs_mgr map[point][]*item
}

func NewAop() *aop {
	return &aop{
		funcs_mgr: make(map[point][]*item),
	}
}

func (this *aop) AddFunc(p point, f func() error, opts ...*Option) {
	opt := Options().SetName("").SetPriority(0).Merge(opts...)
	this.add(p, run_func(f), opt)
}

func (this *aop) Add(p point, f IRunAble, opts ...*Option) {
	opt := Options().SetName("").SetPriority(0).Merge(opts...)
	this.add(p, f, opt)
}

func (this *aop) add(p point, f IRunAble, opt *Option) {
	this.funcs_mgr[p] = append(this.funcs_mgr[p], &item{f: f, opt: opt})
	sort.SliceStable(this.funcs_mgr[p], func(i, j int) bool {
		return this.funcs_mgr[p][i].opt.getPriority() > this.funcs_mgr[p][j].opt.getPriority()
	})
	if !contains(this._points, p) {
		this._points = append(this._points, p)
		sort.SliceStable(this._points, func(i, j int) bool {
			return this._points[i] > this._points[j]
		})
	}
}

func contains(points []point, p point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}

func (this *aop) RunPoint(p point) error {
	for _, item := range this.funcs_mgr[p] {
		if item.is_run {
			continue
		}
		item.is_run = true
		name := item.opt.getName()
		fmt.Printf("[AOP] ################ %v 执行开始 ##############\n", name)
		if err := item.f.Run(); err != nil {
			fmt.Printf("[AOP] ################ %v 执行失败 ##############%v\n", name, err)
			return err
		}
		fmt.Printf("[AOP] ################ %v 执行成功 ##############\n", name)
	}
	return nil
}

func (this *aop) Run() error {
	for _, p := range this._points {
		for _, item := range this.funcs_mgr[p] {
			if item.is_run {
				continue
			}
			item.is_run = true
			name := item.opt.getName()
			fmt.Printf("[AOP] ################ %v 执行开始 ##############\n", name)
			if err := item.f.Run(); err != nil {
				fmt.Printf("[AOP] ################ %v 执行失败 ##############%v\n", name, err)
				return err
			}
			fmt.Printf("[AOP] ################ %v 执行成功 ##############\n", name)
		}
	}
	return nil
}
