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

// ListBookkeepingSubjectBalances invokes the companyreg.ListBookkeepingSubjectBalances API synchronously
func (client *Client) ListBookkeepingSubjectBalances(request *ListBookkeepingSubjectBalancesRequest) (response *ListBookkeepingSubjectBalancesResponse, err error) {
	response = CreateListBookkeepingSubjectBalancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListBookkeepingSubjectBalancesWithChan invokes the companyreg.ListBookkeepingSubjectBalances API asynchronously
func (client *Client) ListBookkeepingSubjectBalancesWithChan(request *ListBookkeepingSubjectBalancesRequest) (<-chan *ListBookkeepingSubjectBalancesResponse, <-chan error) {
	responseChan := make(chan *ListBookkeepingSubjectBalancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBookkeepingSubjectBalances(request)
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

// ListBookkeepingSubjectBalancesWithCallback invokes the companyreg.ListBookkeepingSubjectBalances API asynchronously
func (client *Client) ListBookkeepingSubjectBalancesWithCallback(request *ListBookkeepingSubjectBalancesRequest, callback func(response *ListBookkeepingSubjectBalancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBookkeepingSubjectBalancesResponse
		var err error
		defer close(result)
		response, err = client.ListBookkeepingSubjectBalances(request)
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

// ListBookkeepingSubjectBalancesRequest is the request struct for api ListBookkeepingSubjectBalances
type ListBookkeepingSubjectBalancesRequest struct {
	*requests.RpcRequest
	Period       requests.Integer `position:"Query" name:"Period"`
	Year         requests.Integer `position:"Query" name:"Year"`
	ProduceBizId string           `position:"Query" name:"ProduceBizId"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListBookkeepingSubjectBalancesResponse is the response struct for api ListBookkeepingSubjectBalances
type ListBookkeepingSubjectBalancesResponse struct {
	*responses.BaseResponse
	RequestId                  string                           `json:"RequestId" xml:"RequestId"`
	TotalCount                 int                              `json:"TotalCount" xml:"TotalCount"`
	PageSize                   int                              `json:"PageSize" xml:"PageSize"`
	PageNumber                 int                              `json:"PageNumber" xml:"PageNumber"`
	Success                    bool                             `json:"Success" xml:"Success"`
	ErrorMsg                   string                           `json:"ErrorMsg" xml:"ErrorMsg"`
	BookkeepingSubjectBalances []BookkeepingSubjectBalancesItem `json:"BookkeepingSubjectBalances" xml:"BookkeepingSubjectBalances"`
}

// CreateListBookkeepingSubjectBalancesRequest creates a request to invoke ListBookkeepingSubjectBalances API
func CreateListBookkeepingSubjectBalancesRequest() (request *ListBookkeepingSubjectBalancesRequest) {
	request = &ListBookkeepingSubjectBalancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ListBookkeepingSubjectBalances", "", "")
	request.Method = requests.POST
	return
}

// CreateListBookkeepingSubjectBalancesResponse creates a response to parse from ListBookkeepingSubjectBalances response
func CreateListBookkeepingSubjectBalancesResponse() (response *ListBookkeepingSubjectBalancesResponse) {
	response = &ListBookkeepingSubjectBalancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
