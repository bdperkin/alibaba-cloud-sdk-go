package cms

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

func (client *Client) NodeStatusList(request *NodeStatusListRequest) (response *NodeStatusListResponse, err error) {
	response = CreateNodeStatusListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) NodeStatusListWithChan(request *NodeStatusListRequest) (<-chan *NodeStatusListResponse, <-chan error) {
	responseChan := make(chan *NodeStatusListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeStatusList(request)
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

func (client *Client) NodeStatusListWithCallback(request *NodeStatusListRequest, callback func(response *NodeStatusListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeStatusListResponse
		var err error
		defer close(result)
		response, err = client.NodeStatusList(request)
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

type NodeStatusListRequest struct {
	*requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

type NodeStatusListResponse struct {
	*responses.BaseResponse
	ErrorCode      int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	NodeStatusList struct {
		NodeStatus []struct {
			InstanceId  string `json:"InstanceId" xml:"InstanceId"`
			AutoInstall bool   `json:"AutoInstall" xml:"AutoInstall"`
			Status      string `json:"Status" xml:"Status"`
		} `json:"NodeStatus" xml:"NodeStatus"`
	} `json:"NodeStatusList" xml:"NodeStatusList"`
}

func CreateNodeStatusListRequest() (request *NodeStatusListRequest) {
	request = &NodeStatusListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "NodeStatusList", "cms", "openAPI")
	return
}

func CreateNodeStatusListResponse() (response *NodeStatusListResponse) {
	response = &NodeStatusListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
