/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sdk

import "context"

var Domain = "https://app.aitubo.ai"

type Client struct {
	apiKey string
	job    Job
	model  Model
}

func NewClient(key string) *Client {
	return &Client{
		apiKey: key,
		job:    NewJob(key),
		model:  NewModel(),
	}
}

func (c *Client) CreateText2ImageJob(ctx context.Context, req *CreateText2ImageJobRequest) (*CreateJobResponse, error) {
	return c.job.Create(ctx, req, false)
}

func (c *Client) CreateImage2ImageJob(ctx context.Context, req *CreateImage2ImageJobRequest) (*CreateJobResponse, error) {
	return c.job.Create(ctx, req, false)
}

func (c *Client) CreateControlNetJob(ctx context.Context, req *CreateControlNetJobRequest) (*CreateJobResponse, error) {
	return c.job.Create(ctx, req, false)
}

func (c *Client) CreateUpscaleJob(ctx context.Context, req *CreateUpscaleJobRequest) (*CreateJobResponse, error) {
	return c.job.Create(ctx, req, true)
}

func (c *Client) QueryJob(id string) (*QueryJobResponse, error) {
	return c.job.Query(id)
}

func (c *Client) ListModels(req *ListModelRequest) (*ListModelResponse, error) {
	return c.model.List(req)
}

func (c *Client) QueryModel() {
	c.model.Query()
}
