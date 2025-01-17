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

// ListProduceAuthorization invokes the companyreg.ListProduceAuthorization API synchronously
func (client *Client) ListProduceAuthorization(request *ListProduceAuthorizationRequest) (response *ListProduceAuthorizationResponse, err error) {
	response = CreateListProduceAuthorizationResponse()
	err = client.DoAction(request, response)
	return
}

// ListProduceAuthorizationWithChan invokes the companyreg.ListProduceAuthorization API asynchronously
func (client *Client) ListProduceAuthorizationWithChan(request *ListProduceAuthorizationRequest) (<-chan *ListProduceAuthorizationResponse, <-chan error) {
	responseChan := make(chan *ListProduceAuthorizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProduceAuthorization(request)
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

// ListProduceAuthorizationWithCallback invokes the companyreg.ListProduceAuthorization API asynchronously
func (client *Client) ListProduceAuthorizationWithCallback(request *ListProduceAuthorizationRequest, callback func(response *ListProduceAuthorizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProduceAuthorizationResponse
		var err error
		defer close(result)
		response, err = client.ListProduceAuthorization(request)
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

// ListProduceAuthorizationRequest is the request struct for api ListProduceAuthorization
type ListProduceAuthorizationRequest struct {
	*requests.RpcRequest
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	BizType  string           `position:"Query" name:"BizType"`
	BizId    string           `position:"Query" name:"BizId"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
}

// ListProduceAuthorizationResponse is the response struct for api ListProduceAuthorization
type ListProduceAuthorizationResponse struct {
	*responses.BaseResponse
	CurrentPageNum int                     `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int                     `json:"PageSize" xml:"PageSize"`
	RequestId      string                  `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                     `json:"TotalItemNum" xml:"TotalItemNum"`
	TotalPageNum   int                     `json:"TotalPageNum" xml:"TotalPageNum"`
	Success        bool                    `json:"Success" xml:"Success"`
	Data           []ProduceAuthorizations `json:"Data" xml:"Data"`
}

// CreateListProduceAuthorizationRequest creates a request to invoke ListProduceAuthorization API
func CreateListProduceAuthorizationRequest() (request *ListProduceAuthorizationRequest) {
	request = &ListProduceAuthorizationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "ListProduceAuthorization", "", "")
	request.Method = requests.POST
	return
}

// CreateListProduceAuthorizationResponse creates a response to parse from ListProduceAuthorization response
func CreateListProduceAuthorizationResponse() (response *ListProduceAuthorizationResponse) {
	response = &ListProduceAuthorizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
