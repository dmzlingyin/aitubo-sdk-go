package sdk

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Job interface {
	Create(req any) (*CreateJobResponse, error)
	Upscale(req *CreateUpscaleJobRequest) (*CreateJobResponse, error)
	Query()
}

type job struct {
	key string
}

func NewJob(key string) Job {
	return &job{
		key: key,
	}
}

func (j *job) Create(req any) (*CreateJobResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	r, err := http.NewRequest("POST", Domain+"/api/job/create", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+j.key)

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

func (j *job) Upscale(req *CreateUpscaleJobRequest) (*CreateJobResponse, error) {
	return nil, nil
}

func (j *job) Query() {}
