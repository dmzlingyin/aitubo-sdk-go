package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Job interface {
	Create(req any, upscale bool) (*CreateJobResponse, error)
	Query(id string) (*QueryJobResponse, error)
}

func NewJob(key string) Job {
	return &job{
		key: key,
	}
}

type job struct {
	key string
}

func (j *job) Create(req any, upscale bool) (*CreateJobResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	path := "/api/job/create"
	if upscale {
		path = "/api/job/upscale"
	}
	r, err := http.NewRequest("POST", Domain+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+j.key)
	r.Header.Set("X-App-Name", "aitubo-api")
	r.Header.Set("X-App-Version", Version)

	client := http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res CreateJobResponse
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (j *job) Query(id string) (*QueryJobResponse, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	resp, err := http.Get(Domain + "/api/job/get?id=" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res QueryJobResponse
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
