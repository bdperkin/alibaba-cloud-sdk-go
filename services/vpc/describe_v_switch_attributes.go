package vpc

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

func (client *Client) DescribeVSwitchAttributes(request *DescribeVSwitchAttributesRequest) (response *DescribeVSwitchAttributesResponse, err error) {
	response = CreateDescribeVSwitchAttributesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVSwitchAttributesWithChan(request *DescribeVSwitchAttributesRequest) (<-chan *DescribeVSwitchAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeVSwitchAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVSwitchAttributes(request)
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

func (client *Client) DescribeVSwitchAttributesWithCallback(request *DescribeVSwitchAttributesRequest, callback func(response *DescribeVSwitchAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVSwitchAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVSwitchAttributes(request)
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

type DescribeVSwitchAttributesRequest struct {
	*requests.RpcRequest
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeVSwitchAttributesResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	VSwitchId               string `json:"VSwitchId" xml:"VSwitchId"`
	VpcId                   string `json:"VpcId" xml:"VpcId"`
	Status                  string `json:"Status" xml:"Status"`
	CidrBlock               string `json:"CidrBlock" xml:"CidrBlock"`
	ZoneId                  string `json:"ZoneId" xml:"ZoneId"`
	AvailableIpAddressCount int    `json:"AvailableIpAddressCount" xml:"AvailableIpAddressCount"`
	Description             string `json:"Description" xml:"Description"`
	VSwitchName             string `json:"VSwitchName" xml:"VSwitchName"`
	CreationTime            string `json:"CreationTime" xml:"CreationTime"`
	IsDefault               bool   `json:"IsDefault" xml:"IsDefault"`
	CloudResources          struct {
		CloudResourceSetType []struct {
			ResourceType  string `json:"ResourceType" xml:"ResourceType"`
			ResourceCount int    `json:"ResourceCount" xml:"ResourceCount"`
		} `json:"CloudResourceSetType" xml:"CloudResourceSetType"`
	} `json:"CloudResources" xml:"CloudResources"`
}

func CreateDescribeVSwitchAttributesRequest() (request *DescribeVSwitchAttributesRequest) {
	request = &DescribeVSwitchAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVSwitchAttributes", "vpc", "openAPI")
	return
}

func CreateDescribeVSwitchAttributesResponse() (response *DescribeVSwitchAttributesResponse) {
	response = &DescribeVSwitchAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
