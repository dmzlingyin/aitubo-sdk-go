package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Job interface {
	Create(ctx context.Context, req any, upscale bool) (*CreateJobResponse, error)
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

func (j *job) Create(ctx context.Context, req any, upscale bool) (*CreateJobResponse, error) {
	key, err := j.getKey(ctx)
	if err != nil {
		return nil, err
	}
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
	r.Header.Set("Authorization", "Bearer "+key)
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

func (j *job) getKey(ctx context.Context) (string, error) {
	key := j.key
	if v, ok := ctx.Value("key").(string); ok {
		key = v
	}
	if key == "" {
		return "", errors.New("key is empty")
	}
	return key, nil
}
