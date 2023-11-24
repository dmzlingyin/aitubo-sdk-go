package sdk

type Model interface {
	ListModels()
	QueryModel()
}

type model struct {
}

func NewModel() Model {
	return &model{}
}

func (m *model) ListModels() {
}

func (m *model) QueryModel() {
}
