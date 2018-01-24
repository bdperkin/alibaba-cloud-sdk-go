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

func (client *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
	response = CreateRunInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RunInstancesWithChan(request *RunInstancesRequest) (<-chan *RunInstancesResponse, <-chan error) {
	responseChan := make(chan *RunInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunInstances(request)
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

func (client *Client) RunInstancesWithCallback(request *RunInstancesRequest, callback func(response *RunInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunInstancesResponse
		var err error
		defer close(result)
		response, err = client.RunInstances(request)
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

type RunInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	SecurityEnhancementStrategy string                          `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                          `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              requests.Float                  `position:"Query" name:"SpotPriceLimit"`
	HostName                    string                          `position:"Query" name:"HostName"`
	Password                    string                          `position:"Query" name:"Password"`
	Tag                         *[]RunInstancesTag              `position:"Query" name:"Tag"  type:"Repeated"`
	OwnerId                     requests.Integer                `position:"Query" name:"OwnerId"`
	VSwitchId                   string                          `position:"Query" name:"VSwitchId"`
	SpotStrategy                string                          `position:"Query" name:"SpotStrategy"`
	InstanceName                string                          `position:"Query" name:"InstanceName"`
	InternetChargeType          string                          `position:"Query" name:"InternetChargeType"`
	ZoneId                      string                          `position:"Query" name:"ZoneId"`
	InternetMaxBandwidthIn      requests.Integer                `position:"Query" name:"InternetMaxBandwidthIn"`
	ImageId                     string                          `position:"Query" name:"ImageId"`
	ClientToken                 string                          `position:"Query" name:"ClientToken"`
	IoOptimized                 string                          `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                          `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     requests.Integer                `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                 string                          `position:"Query" name:"Description"`
	SystemDiskCategory          string                          `position:"Query" name:"SystemDisk.Category"`
	UserData                    string                          `position:"Query" name:"UserData"`
	InstanceType                string                          `position:"Query" name:"InstanceType"`
	NetworkInterface            *[]RunInstancesNetworkInterface `position:"Query" name:"NetworkInterface"  type:"Repeated"`
	Amount                      requests.Integer                `position:"Query" name:"Amount"`
	ResourceOwnerAccount        string                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                          `position:"Query" name:"OwnerAccount"`
	SystemDiskDiskName          string                          `position:"Query" name:"SystemDisk.DiskName"`
	RamRoleName                 string                          `position:"Query" name:"RamRoleName"`
	AutoReleaseTime             string                          `position:"Query" name:"AutoReleaseTime"`
	DataDisk                    *[]RunInstancesDataDisk         `position:"Query" name:"DataDisk"  type:"Repeated"`
	SystemDiskSize              string                          `position:"Query" name:"SystemDisk.Size"`
	SystemDiskDescription       string                          `position:"Query" name:"SystemDisk.Description"`
}

type RunInstancesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}
type RunInstancesNetworkInterface struct {
	PrimaryIpAddress     string `name:"PrimaryIpAddress"`
	VSwitchId            string `name:"VSwitchId"`
	SecurityGroupId      string `name:"SecurityGroupId"`
	NetworkInterfaceName string `name:"NetworkInterfaceName"`
	Description          string `name:"Description"`
}
type RunInstancesDataDisk struct {
	Size               string `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	Encrypted          string `name:"Encrypted"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
}

type RunInstancesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	InstanceIdSets struct {
		InstanceIdSet []string `json:"InstanceIdSet" xml:"InstanceIdSet"`
	} `json:"InstanceIdSets" xml:"InstanceIdSets"`
}

func CreateRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RunInstances", "ecs", "openAPI")
	return
}

func CreateRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
