package aop

var d = NewAop()

func AddFunc(p Point, f func() error, opts ...*Option) {
	d.AddFunc(p, f, opts...)
}

func Add(p Point, f IRunAble, opts ...*Option) {
	d.Add(p, f, opts...)
}

func RunPoint(p Point) error {
	return d.RunPoint(p)
}

func Run() error {
	return d.Run()
}
