package aop

var d = New(FlagInit)

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

// 凸显一个切入口
func InjectFunc(p Point, f func() error, opts ...*Option) {
	d.AddFunc(p, f, opts...)
}

func Inject(p Point, f IRunAble, opts ...*Option) {
	d.Add(p, f, opts...)
}
