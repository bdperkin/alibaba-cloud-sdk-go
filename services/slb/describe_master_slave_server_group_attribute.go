package slb

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

func (client *Client) DescribeMasterSlaveServerGroupAttribute(request *DescribeMasterSlaveServerGroupAttributeRequest) (response *DescribeMasterSlaveServerGroupAttributeResponse, err error) {
	response = CreateDescribeMasterSlaveServerGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeMasterSlaveServerGroupAttributeWithChan(request *DescribeMasterSlaveServerGroupAttributeRequest) (<-chan *DescribeMasterSlaveServerGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeMasterSlaveServerGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMasterSlaveServerGroupAttribute(request)
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

func (client *Client) DescribeMasterSlaveServerGroupAttributeWithCallback(request *DescribeMasterSlaveServerGroupAttributeRequest, callback func(response *DescribeMasterSlaveServerGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMasterSlaveServerGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeMasterSlaveServerGroupAttribute(request)
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

type DescribeMasterSlaveServerGroupAttributeRequest struct {
	*requests.RpcRequest
	Tags                     string           `position:"Query" name:"Tags"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId              string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	MasterSlaveServerGroupId string           `position:"Query" name:"MasterSlaveServerGroupId"`
}

type DescribeMasterSlaveServerGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId                  string `json:"RequestId" xml:"RequestId"`
	MasterSlaveServerGroupId   string `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
	MasterSlaveServerGroupName string `json:"MasterSlaveServerGroupName" xml:"MasterSlaveServerGroupName"`
	MasterSlaveBackendServers  struct {
		MasterSlaveBackendServer []struct {
			ServerId   string `json:"ServerId" xml:"ServerId"`
			Port       int    `json:"Port" xml:"Port"`
			Weight     int    `json:"Weight" xml:"Weight"`
			ServerType string `json:"ServerType" xml:"ServerType"`
		} `json:"MasterSlaveBackendServer" xml:"MasterSlaveBackendServer"`
	} `json:"MasterSlaveBackendServers" xml:"MasterSlaveBackendServers"`
}

func CreateDescribeMasterSlaveServerGroupAttributeRequest() (request *DescribeMasterSlaveServerGroupAttributeRequest) {
	request = &DescribeMasterSlaveServerGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveServerGroupAttribute", "slb", "openAPI")
	return
}

func CreateDescribeMasterSlaveServerGroupAttributeResponse() (response *DescribeMasterSlaveServerGroupAttributeResponse) {
	response = &DescribeMasterSlaveServerGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
