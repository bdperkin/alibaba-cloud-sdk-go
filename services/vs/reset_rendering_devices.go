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

// ResetRenderingDevices invokes the vs.ResetRenderingDevices API synchronously
func (client *Client) ResetRenderingDevices(request *ResetRenderingDevicesRequest) (response *ResetRenderingDevicesResponse, err error) {
	response = CreateResetRenderingDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// ResetRenderingDevicesWithChan invokes the vs.ResetRenderingDevices API asynchronously
func (client *Client) ResetRenderingDevicesWithChan(request *ResetRenderingDevicesRequest) (<-chan *ResetRenderingDevicesResponse, <-chan error) {
	responseChan := make(chan *ResetRenderingDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetRenderingDevices(request)
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

// ResetRenderingDevicesWithCallback invokes the vs.ResetRenderingDevices API asynchronously
func (client *Client) ResetRenderingDevicesWithCallback(request *ResetRenderingDevicesRequest, callback func(response *ResetRenderingDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetRenderingDevicesResponse
		var err error
		defer close(result)
		response, err = client.ResetRenderingDevices(request)
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

// ResetRenderingDevicesRequest is the request struct for api ResetRenderingDevices
type ResetRenderingDevicesRequest struct {
	*requests.RpcRequest
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds string           `position:"Query" name:"InstanceIds"`
}

// ResetRenderingDevicesResponse is the response struct for api ResetRenderingDevices
type ResetRenderingDevicesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetRenderingDevicesRequest creates a request to invoke ResetRenderingDevices API
func CreateResetRenderingDevicesRequest() (request *ResetRenderingDevicesRequest) {
	request = &ResetRenderingDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ResetRenderingDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateResetRenderingDevicesResponse creates a response to parse from ResetRenderingDevices response
func CreateResetRenderingDevicesResponse() (response *ResetRenderingDevicesResponse) {
	response = &ResetRenderingDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
