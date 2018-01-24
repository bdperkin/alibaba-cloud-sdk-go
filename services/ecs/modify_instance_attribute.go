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

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
	response = CreateModifyInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyInstanceAttributeWithChan(request *ModifyInstanceAttributeRequest) (<-chan *ModifyInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceAttribute(request)
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

func (client *Client) ModifyInstanceAttributeWithCallback(request *ModifyInstanceAttributeRequest, callback func(response *ModifyInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceAttribute(request)
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

type ModifyInstanceAttributeRequest struct {
	*requests.RpcRequest
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Password             string           `position:"Query" name:"Password"`
	HostName             string           `position:"Query" name:"HostName"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	InstanceName         string           `position:"Query" name:"InstanceName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Recyclable           requests.Boolean `position:"Query" name:"Recyclable"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type ModifyInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyInstanceAttributeRequest() (request *ModifyInstanceAttributeRequest) {
	request = &ModifyInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAttribute", "ecs", "openAPI")
	return
}

func CreateModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
	response = &ModifyInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
