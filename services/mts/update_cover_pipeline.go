package mts

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) UpdateCoverPipeline(request *UpdateCoverPipelineRequest) (response *UpdateCoverPipelineResponse, err error) {
	response = CreateUpdateCoverPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateCoverPipelineWithChan(request *UpdateCoverPipelineRequest) (<-chan *UpdateCoverPipelineResponse, <-chan error) {
	responseChan := make(chan *UpdateCoverPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCoverPipeline(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}

	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) UpdateCoverPipelineWithCallback(request *UpdateCoverPipelineRequest, callback func(response *UpdateCoverPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCoverPipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdateCoverPipeline(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type UpdateCoverPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Role                 string           `position:"Query" name:"Role"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	State                string           `position:"Query" name:"State"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
}

type UpdateCoverPipelineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Pipeline  struct {
		Id           string `json:"Id" xml:"Id"`
		Name         string `json:"Name" xml:"Name"`
		State        string `json:"State" xml:"State"`
		Priority     int    `json:"Priority" xml:"Priority"`
		Role         string `json:"Role" xml:"Role"`
		NotifyConfig struct {
			Topic string `json:"Topic" xml:"Topic"`
			Queue string `json:"Queue" xml:"Queue"`
		} `json:"NotifyConfig" xml:"NotifyConfig"`
	} `json:"Pipeline" xml:"Pipeline"`
}

func CreateUpdateCoverPipelineRequest() (request *UpdateCoverPipelineRequest) {
	request = &UpdateCoverPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateCoverPipeline", "mts", "openAPI")
	return
}

func CreateUpdateCoverPipelineResponse() (response *UpdateCoverPipelineResponse) {
	response = &UpdateCoverPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
