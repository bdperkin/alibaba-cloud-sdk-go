package live

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

// PlayChoosenShow invokes the live.PlayChoosenShow API synchronously
func (client *Client) PlayChoosenShow(request *PlayChoosenShowRequest) (response *PlayChoosenShowResponse, err error) {
	response = CreatePlayChoosenShowResponse()
	err = client.DoAction(request, response)
	return
}

// PlayChoosenShowWithChan invokes the live.PlayChoosenShow API asynchronously
func (client *Client) PlayChoosenShowWithChan(request *PlayChoosenShowRequest) (<-chan *PlayChoosenShowResponse, <-chan error) {
	responseChan := make(chan *PlayChoosenShowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PlayChoosenShow(request)
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

// PlayChoosenShowWithCallback invokes the live.PlayChoosenShow API asynchronously
func (client *Client) PlayChoosenShowWithCallback(request *PlayChoosenShowRequest, callback func(response *PlayChoosenShowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PlayChoosenShowResponse
		var err error
		defer close(result)
		response, err = client.PlayChoosenShow(request)
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

// PlayChoosenShowRequest is the request struct for api PlayChoosenShow
type PlayChoosenShowRequest struct {
	*requests.RpcRequest
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	ShowId   string           `position:"Query" name:"ShowId"`
}

// PlayChoosenShowResponse is the response struct for api PlayChoosenShow
type PlayChoosenShowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ShowId    string `json:"ShowId" xml:"ShowId"`
}

// CreatePlayChoosenShowRequest creates a request to invoke PlayChoosenShow API
func CreatePlayChoosenShowRequest() (request *PlayChoosenShowRequest) {
	request = &PlayChoosenShowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "PlayChoosenShow", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePlayChoosenShowResponse creates a response to parse from PlayChoosenShow response
func CreatePlayChoosenShowResponse() (response *PlayChoosenShowResponse) {
	response = &PlayChoosenShowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
