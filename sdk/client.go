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

var BaseURL = "https://creator.aitubo.ai/api"

type Client struct {
	apiKey  string
	baseURL string
	job     Job
	model   Model
}

func NewClient(key string) *Client {
	return &Client{
		apiKey:  key,
		baseURL: BaseURL,
		job:     NewJob(),
		model:   NewModel(),
	}
}

func (c *Client) CreateText2ImageJob(req *CreateText2ImageJobRequest) (*CreateJobResponse, error) {
	return c.job.CreateText2ImageJob(req)
}

func (c *Client) CreateImage2ImageJob(req *CreateImage2ImageJobRequest) (*CreateJobResponse, error) {
	return c.job.CreateImage2ImageJob(req)
}

func (c *Client) CreateControlNetJob(req *CreateControlNetJobRequest) (*CreateJobResponse, error) {
	return c.job.CreateControlNetJob(req)
}

func (c *Client) CreateUpscaleJob(req *CreateUpscaleJobRequest) (*CreateJobResponse, error) {
	return c.job.CreateUpscaleJob(req)
}

func (c *Client) QueryJob() {
	c.job.QueryJob()
}

func (c *Client) ListModels() {
	c.model.ListModels()
}

func (c *Client) QueryModel() {
	c.model.QueryModel()
}
