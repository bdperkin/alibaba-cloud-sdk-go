package cloudphoto

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

// GetEvent invokes the cloudphoto.GetEvent API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/getevent.html
func (client *Client) GetEvent(request *GetEventRequest) (response *GetEventResponse, err error) {
	response = CreateGetEventResponse()
	err = client.DoAction(request, response)
	return
}

// GetEventWithChan invokes the cloudphoto.GetEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetEventWithChan(request *GetEventRequest) (<-chan *GetEventResponse, <-chan error) {
	responseChan := make(chan *GetEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEvent(request)
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

// GetEventWithCallback invokes the cloudphoto.GetEvent API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/getevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetEventWithCallback(request *GetEventRequest, callback func(response *GetEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEventResponse
		var err error
		defer close(result)
		response, err = client.GetEvent(request)
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

// GetEventRequest is the request struct for api GetEvent
type GetEventRequest struct {
	*requests.RpcRequest
	EventId   requests.Integer `position:"Query" name:"EventId"`
	StoreName string           `position:"Query" name:"StoreName"`
	LibraryId string           `position:"Query" name:"LibraryId"`
}

// GetEventResponse is the response struct for api GetEvent
type GetEventResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Event     Event  `json:"Event" xml:"Event"`
}

// CreateGetEventRequest creates a request to invoke GetEvent API
func CreateGetEventRequest() (request *GetEventRequest) {
	request = &GetEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetEvent", "cloudphoto", "openAPI")
	return
}

// CreateGetEventResponse creates a response to parse from GetEvent response
func CreateGetEventResponse() (response *GetEventResponse) {
	response = &GetEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
