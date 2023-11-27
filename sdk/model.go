package sdk

type Model interface {
	List()
	Query()
}

type model struct {
}

func NewModel() Model {
	return &model{}
}

func (m *model) List() {
}

func (m *model) Query() {
}
