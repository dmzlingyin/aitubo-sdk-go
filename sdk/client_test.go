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

func TestQueryJob(t *testing.T) {
	client := NewClient("Your API Key")
	res, err := client.QueryJob("656d8f7077ce99faf53786c1")
	if err != nil {
		t.Fatal(err)
	}
	if res.Code == 0 {
		if res.Data.Status == 2 {
			t.Log(res.Data.Result.Data)
		} else {
			t.Log(res.Data.Status)
		}
	} else {
		t.Log(res.Code)
	}
}
