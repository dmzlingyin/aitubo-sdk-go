package sdk

import "testing"

func TestCreateText2ImageJob(t *testing.T) {
	client := NewClient("Your API Key")
	res, err := client.CreateText2ImageJob(&CreateText2ImageJobRequest{
		BaseRequest: BaseRequest{
			Prompt: "a girl",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
