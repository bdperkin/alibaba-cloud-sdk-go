package ecs

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

// DescribeInstancesFullStatus invokes the ecs.DescribeInstancesFullStatus API synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancesfullstatus.html
func (client *Client) DescribeInstancesFullStatus(request *DescribeInstancesFullStatusRequest) (response *DescribeInstancesFullStatusResponse, err error) {
	response = CreateDescribeInstancesFullStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstancesFullStatusWithChan invokes the ecs.DescribeInstancesFullStatus API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancesfullstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstancesFullStatusWithChan(request *DescribeInstancesFullStatusRequest) (<-chan *DescribeInstancesFullStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeInstancesFullStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstancesFullStatus(request)
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

// DescribeInstancesFullStatusWithCallback invokes the ecs.DescribeInstancesFullStatus API asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancesfullstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstancesFullStatusWithCallback(request *DescribeInstancesFullStatusRequest, callback func(response *DescribeInstancesFullStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstancesFullStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstancesFullStatus(request)
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

// DescribeInstancesFullStatusRequest is the request struct for api DescribeInstancesFullStatus
type DescribeInstancesFullStatusRequest struct {
	*requests.RpcRequest
}

// DescribeInstancesFullStatusResponse is the response struct for api DescribeInstancesFullStatus
type DescribeInstancesFullStatusResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	TotalCount            int                   `json:"TotalCount" xml:"TotalCount"`
	PageNumber            int                   `json:"PageNumber" xml:"PageNumber"`
	PageSize              int                   `json:"PageSize" xml:"PageSize"`
	InstanceFullStatusSet InstanceFullStatusSet `json:"InstanceFullStatusSet" xml:"InstanceFullStatusSet"`
}

// CreateDescribeInstancesFullStatusRequest creates a request to invoke DescribeInstancesFullStatus API
func CreateDescribeInstancesFullStatusRequest() (request *DescribeInstancesFullStatusRequest) {
	request = &DescribeInstancesFullStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstancesFullStatus", "ecs", "openAPI")
	return
}

// CreateDescribeInstancesFullStatusResponse creates a response to parse from DescribeInstancesFullStatus response
func CreateDescribeInstancesFullStatusResponse() (response *DescribeInstancesFullStatusResponse) {
	response = &DescribeInstancesFullStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
