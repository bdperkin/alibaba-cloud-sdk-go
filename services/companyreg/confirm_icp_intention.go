package companyreg

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

// ConfirmIcpIntention invokes the companyreg.ConfirmIcpIntention API synchronously
func (client *Client) ConfirmIcpIntention(request *ConfirmIcpIntentionRequest) (response *ConfirmIcpIntentionResponse, err error) {
	response = CreateConfirmIcpIntentionResponse()
	err = client.DoAction(request, response)
	return
}

// ConfirmIcpIntentionWithChan invokes the companyreg.ConfirmIcpIntention API asynchronously
func (client *Client) ConfirmIcpIntentionWithChan(request *ConfirmIcpIntentionRequest) (<-chan *ConfirmIcpIntentionResponse, <-chan error) {
	responseChan := make(chan *ConfirmIcpIntentionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfirmIcpIntention(request)
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

// ConfirmIcpIntentionWithCallback invokes the companyreg.ConfirmIcpIntention API asynchronously
func (client *Client) ConfirmIcpIntentionWithCallback(request *ConfirmIcpIntentionRequest, callback func(response *ConfirmIcpIntentionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfirmIcpIntentionResponse
		var err error
		defer close(result)
		response, err = client.ConfirmIcpIntention(request)
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

// ConfirmIcpIntentionRequest is the request struct for api ConfirmIcpIntention
type ConfirmIcpIntentionRequest struct {
	*requests.RpcRequest
	BizIds string `position:"Query" name:"BizIds"`
}

// ConfirmIcpIntentionResponse is the response struct for api ConfirmIcpIntention
type ConfirmIcpIntentionResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	ConfirmUrl string `json:"ConfirmUrl" xml:"ConfirmUrl"`
}

// CreateConfirmIcpIntentionRequest creates a request to invoke ConfirmIcpIntention API
func CreateConfirmIcpIntentionRequest() (request *ConfirmIcpIntentionRequest) {
	request = &ConfirmIcpIntentionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ConfirmIcpIntention", "", "")
	request.Method = requests.POST
	return
}

// CreateConfirmIcpIntentionResponse creates a response to parse from ConfirmIcpIntention response
func CreateConfirmIcpIntentionResponse() (response *ConfirmIcpIntentionResponse) {
	response = &ConfirmIcpIntentionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
