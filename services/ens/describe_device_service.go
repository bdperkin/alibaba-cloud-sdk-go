package ens

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

// DescribeDeviceService invokes the ens.DescribeDeviceService API synchronously
func (client *Client) DescribeDeviceService(request *DescribeDeviceServiceRequest) (response *DescribeDeviceServiceResponse, err error) {
	response = CreateDescribeDeviceServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeviceServiceWithChan invokes the ens.DescribeDeviceService API asynchronously
func (client *Client) DescribeDeviceServiceWithChan(request *DescribeDeviceServiceRequest) (<-chan *DescribeDeviceServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeDeviceServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeviceService(request)
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

// DescribeDeviceServiceWithCallback invokes the ens.DescribeDeviceService API asynchronously
func (client *Client) DescribeDeviceServiceWithCallback(request *DescribeDeviceServiceRequest, callback func(response *DescribeDeviceServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeviceServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeviceService(request)
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

// DescribeDeviceServiceRequest is the request struct for api DescribeDeviceService
type DescribeDeviceServiceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	AppId      string `position:"Query" name:"AppId"`
}

// DescribeDeviceServiceResponse is the response struct for api DescribeDeviceService
type DescribeDeviceServiceResponse struct {
	*responses.BaseResponse
	RequestId           string               `json:"RequestId" xml:"RequestId"`
	ResourceDetailInfos []ResourceDetailInfo `json:"ResourceDetailInfos" xml:"ResourceDetailInfos"`
}

// CreateDescribeDeviceServiceRequest creates a request to invoke DescribeDeviceService API
func CreateDescribeDeviceServiceRequest() (request *DescribeDeviceServiceRequest) {
	request = &DescribeDeviceServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeDeviceService", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeDeviceServiceResponse creates a response to parse from DescribeDeviceService response
func CreateDescribeDeviceServiceResponse() (response *DescribeDeviceServiceResponse) {
	response = &DescribeDeviceServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
