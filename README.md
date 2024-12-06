#### USAGE
1. Point 越小越优先执行,遵行iota自增的特性来定义Point
2. Priority 越大越优先执行,代表优先级
```golang
func main() {
	aop.RunPoint(17)
	aop.RunPoint(15) //强制执行切入点
	aop.Run() //执行剩下的函数
}

func init() {
	aop.AddFunc(15, func() error {
		fmt.Println("init 15")
		return nil
	}, aop.Options().SetName("15"))

}

func init() {
	aop.Add(17, &obj{}, aop.Options().SetName("17").SetPriority(19))
	aop.AddFunc(17, func() error {
		fmt.Println("init 17 run func")
		return errors.New("ddddd")
	}, aop.Options().SetName("17 func ").SetPriority(18))
}

func init() {
	aop.AddFunc(18, func() error {
		fmt.Println("init 18")
		return nil
	}, aop.Options().SetName("18"))
}

type obj struct {
}

func (this *obj) Run() error {
	fmt.Println("obj run")
	return nil
}

