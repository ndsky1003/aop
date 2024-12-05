package aop

type Option struct {
	name     *string
	priority *int
}

func Options() *Option {
	return &Option{}
}

func (this *Option) SetPriority(p int) *Option {
	this.priority = &p
	return this
}
func (this *Option) getPriority() int {
	if this == nil || this.priority == nil {
		return 0
	}
	return *this.priority
}

func (this *Option) SetName(n string) *Option {
	this.name = &n
	return this
}

func (this *Option) getName() string {
	if this == nil || this.name == nil {
		return ""
	}
	return *this.name
}

func (this *Option) merge(opt *Option) {
	if opt == nil {
		return
	}
	if opt.priority != nil {
		this.priority = opt.priority
	}
	if opt.name != nil {
		this.name = opt.name
	}
}

func (this *Option) Merge(opts ...*Option) *Option {
	for _, opt := range opts {
		this.merge(opt)
	}
	return this
}
