package devops_rdc

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

// ListDevopsProjectSprints invokes the devops_rdc.ListDevopsProjectSprints API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/listdevopsprojectsprints.html
func (client *Client) ListDevopsProjectSprints(request *ListDevopsProjectSprintsRequest) (response *ListDevopsProjectSprintsResponse, err error) {
	response = CreateListDevopsProjectSprintsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevopsProjectSprintsWithChan invokes the devops_rdc.ListDevopsProjectSprints API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/listdevopsprojectsprints.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDevopsProjectSprintsWithChan(request *ListDevopsProjectSprintsRequest) (<-chan *ListDevopsProjectSprintsResponse, <-chan error) {
	responseChan := make(chan *ListDevopsProjectSprintsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevopsProjectSprints(request)
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

// ListDevopsProjectSprintsWithCallback invokes the devops_rdc.ListDevopsProjectSprints API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/listdevopsprojectsprints.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDevopsProjectSprintsWithCallback(request *ListDevopsProjectSprintsRequest, callback func(response *ListDevopsProjectSprintsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevopsProjectSprintsResponse
		var err error
		defer close(result)
		response, err = client.ListDevopsProjectSprints(request)
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

// ListDevopsProjectSprintsRequest is the request struct for api ListDevopsProjectSprints
type ListDevopsProjectSprintsRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Body" name:"ProjectId"`
	OrgId     string `position:"Body" name:"OrgId"`
}

// ListDevopsProjectSprintsResponse is the response struct for api ListDevopsProjectSprints
type ListDevopsProjectSprintsResponse struct {
	*responses.BaseResponse
	Successful bool     `json:"Successful" xml:"Successful"`
	ErrorCode  string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string   `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Object     []Sprint `json:"Object" xml:"Object"`
}

// CreateListDevopsProjectSprintsRequest creates a request to invoke ListDevopsProjectSprints API
func CreateListDevopsProjectSprintsRequest() (request *ListDevopsProjectSprintsRequest) {
	request = &ListDevopsProjectSprintsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListDevopsProjectSprints", "", "")
	request.Method = requests.POST
	return
}

// CreateListDevopsProjectSprintsResponse creates a response to parse from ListDevopsProjectSprints response
func CreateListDevopsProjectSprintsResponse() (response *ListDevopsProjectSprintsResponse) {
	response = &ListDevopsProjectSprintsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
