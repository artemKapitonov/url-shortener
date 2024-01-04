package convertor

type EntityConvetor struct{}

func New() *EntityConvetor {
	return &EntityConvetor{}
}

func (c *EntityConvetor) Convert() {
	panic("implement me")
}
