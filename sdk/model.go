package sdk

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type Model interface {
	List(req *ListModelRequest) (*ListModelResponse, error)
	Query()
}

type model struct {
}

func NewModel() Model {
	return &model{}
}

func (m *model) List(req *ListModelRequest) (*ListModelResponse, error) {
	resp, err := http.Get(Domain + "/api/model/list?" + m.getParams(req).Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res ListModelResponse
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (m *model) getParams(req *ListModelRequest) url.Values {
	params := url.Values{}
	params.Add("pageIndex", strconv.Itoa(req.PageIndex))
	params.Add("pageSize", strconv.Itoa(req.PageSize))
	params.Add("type", req.Type)
	params.Add("kind", req.Kind)
	params.Add("name", req.Name)
	params.Add("baseModel", req.BaseModel)
	return params
}

func (m *model) Query() {
}
