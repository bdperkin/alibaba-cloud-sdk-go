package lrg

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

// QueryNodeParent invokes the lrg.QueryNodeParent API synchronously
// api document: https://help.aliyun.com/api/lrg/querynodeparent.html
func (client *Client) QueryNodeParent(request *QueryNodeParentRequest) (response *QueryNodeParentResponse, err error) {
	response = CreateQueryNodeParentResponse()
	err = client.DoAction(request, response)
	return
}

// QueryNodeParentWithChan invokes the lrg.QueryNodeParent API asynchronously
// api document: https://help.aliyun.com/api/lrg/querynodeparent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNodeParentWithChan(request *QueryNodeParentRequest) (<-chan *QueryNodeParentResponse, <-chan error) {
	responseChan := make(chan *QueryNodeParentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryNodeParent(request)
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

// QueryNodeParentWithCallback invokes the lrg.QueryNodeParent API asynchronously
// api document: https://help.aliyun.com/api/lrg/querynodeparent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNodeParentWithCallback(request *QueryNodeParentRequest, callback func(response *QueryNodeParentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryNodeParentResponse
		var err error
		defer close(result)
		response, err = client.QueryNodeParent(request)
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

// QueryNodeParentRequest is the request struct for api QueryNodeParent
type QueryNodeParentRequest struct {
	*requests.RoaRequest
	NodeName    string `position:"Path" name:"nodeName"`
	ProductName string `position:"Path" name:"productName"`
}

// QueryNodeParentResponse is the response struct for api QueryNodeParent
type QueryNodeParentResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Success bool                   `json:"success" xml:"success"`
	Message string                 `json:"message" xml:"message"`
	Data    map[string]interface{} `json:"data" xml:"data"`
}

// CreateQueryNodeParentRequest creates a request to invoke QueryNodeParent API
func CreateQueryNodeParentRequest() (request *QueryNodeParentRequest) {
	request = &QueryNodeParentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "QueryNodeParent", "/api/v2/[productName]/node/[nodeName]?action=queryParent", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryNodeParentResponse creates a response to parse from QueryNodeParent response
func CreateQueryNodeParentResponse() (response *QueryNodeParentResponse) {
	response = &QueryNodeParentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}