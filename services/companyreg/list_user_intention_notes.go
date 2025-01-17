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

// ListUserIntentionNotes invokes the companyreg.ListUserIntentionNotes API synchronously
func (client *Client) ListUserIntentionNotes(request *ListUserIntentionNotesRequest) (response *ListUserIntentionNotesResponse, err error) {
	response = CreateListUserIntentionNotesResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserIntentionNotesWithChan invokes the companyreg.ListUserIntentionNotes API asynchronously
func (client *Client) ListUserIntentionNotesWithChan(request *ListUserIntentionNotesRequest) (<-chan *ListUserIntentionNotesResponse, <-chan error) {
	responseChan := make(chan *ListUserIntentionNotesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserIntentionNotes(request)
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

// ListUserIntentionNotesWithCallback invokes the companyreg.ListUserIntentionNotes API asynchronously
func (client *Client) ListUserIntentionNotesWithCallback(request *ListUserIntentionNotesRequest, callback func(response *ListUserIntentionNotesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserIntentionNotesResponse
		var err error
		defer close(result)
		response, err = client.ListUserIntentionNotes(request)
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

// ListUserIntentionNotesRequest is the request struct for api ListUserIntentionNotes
type ListUserIntentionNotesRequest struct {
	*requests.RpcRequest
	BizType        string           `position:"Query" name:"BizType"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	IntentionBizId string           `position:"Query" name:"IntentionBizId"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
}

// ListUserIntentionNotesResponse is the response struct for api ListUserIntentionNotes
type ListUserIntentionNotesResponse struct {
	*responses.BaseResponse
	PageNum      int         `json:"PageNum" xml:"PageNum"`
	PageSize     int         `json:"PageSize" xml:"PageSize"`
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	Success      bool        `json:"Success" xml:"Success"`
	TotalItemNum int         `json:"TotalItemNum" xml:"TotalItemNum"`
	TotalPageNum int         `json:"TotalPageNum" xml:"TotalPageNum"`
	Data         []OpateLogs `json:"Data" xml:"Data"`
}

// CreateListUserIntentionNotesRequest creates a request to invoke ListUserIntentionNotes API
func CreateListUserIntentionNotesRequest() (request *ListUserIntentionNotesRequest) {
	request = &ListUserIntentionNotesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "ListUserIntentionNotes", "", "")
	request.Method = requests.POST
	return
}

// CreateListUserIntentionNotesResponse creates a response to parse from ListUserIntentionNotes response
func CreateListUserIntentionNotesResponse() (response *ListUserIntentionNotesResponse) {
	response = &ListUserIntentionNotesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
