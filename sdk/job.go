package sdk

type Job interface {
	CreateText2ImageJob(req *CreateText2ImageJobRequest) (*CreateJobResponse, error)
	CreateImage2ImageJob(req *CreateImage2ImageJobRequest) (*CreateJobResponse, error)
	CreateControlNetJob(req *CreateControlNetJobRequest) (*CreateJobResponse, error)
	CreateUpscaleJob(req *CreateUpscaleJobRequest) (*CreateJobResponse, error)
	QueryJob()
}

type job struct {
}

func NewJob() Job {
	return &job{}
}

func (j *job) CreateText2ImageJob(req *CreateText2ImageJobRequest) (*CreateJobResponse, error) {
	return nil, nil
}

func (j *job) CreateImage2ImageJob(req *CreateImage2ImageJobRequest) (*CreateJobResponse, error) {
	return nil, nil
}

func (j *job) CreateControlNetJob(req *CreateControlNetJobRequest) (*CreateJobResponse, error) {
	return nil, nil
}

func (j *job) CreateUpscaleJob(req *CreateUpscaleJobRequest) (*CreateJobResponse, error) {
	return nil, nil
}

func (j *job) QueryJob() {
}
