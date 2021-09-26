package vs

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

// BatchSetVsDomainConfigs invokes the vs.BatchSetVsDomainConfigs API synchronously
func (client *Client) BatchSetVsDomainConfigs(request *BatchSetVsDomainConfigsRequest) (response *BatchSetVsDomainConfigsResponse, err error) {
	response = CreateBatchSetVsDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchSetVsDomainConfigsWithChan invokes the vs.BatchSetVsDomainConfigs API asynchronously
func (client *Client) BatchSetVsDomainConfigsWithChan(request *BatchSetVsDomainConfigsRequest) (<-chan *BatchSetVsDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *BatchSetVsDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchSetVsDomainConfigs(request)
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

// BatchSetVsDomainConfigsWithCallback invokes the vs.BatchSetVsDomainConfigs API asynchronously
func (client *Client) BatchSetVsDomainConfigsWithCallback(request *BatchSetVsDomainConfigsRequest, callback func(response *BatchSetVsDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchSetVsDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.BatchSetVsDomainConfigs(request)
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

// BatchSetVsDomainConfigsRequest is the request struct for api BatchSetVsDomainConfigs
type BatchSetVsDomainConfigsRequest struct {
	*requests.RpcRequest
	Functions   string           `position:"Query" name:"Functions"`
	DomainNames string           `position:"Query" name:"DomainNames"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchSetVsDomainConfigsResponse is the response struct for api BatchSetVsDomainConfigs
type BatchSetVsDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchSetVsDomainConfigsRequest creates a request to invoke BatchSetVsDomainConfigs API
func CreateBatchSetVsDomainConfigsRequest() (request *BatchSetVsDomainConfigsRequest) {
	request = &BatchSetVsDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchSetVsDomainConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchSetVsDomainConfigsResponse creates a response to parse from BatchSetVsDomainConfigs response
func CreateBatchSetVsDomainConfigsResponse() (response *BatchSetVsDomainConfigsResponse) {
	response = &BatchSetVsDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
