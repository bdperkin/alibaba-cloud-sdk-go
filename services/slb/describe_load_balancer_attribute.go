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

func (client *Client) DescribeLoadBalancerAttribute(request *DescribeLoadBalancerAttributeRequest) (response *DescribeLoadBalancerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLoadBalancerAttributeWithChan(request *DescribeLoadBalancerAttributeRequest) (<-chan *DescribeLoadBalancerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerAttribute(request)
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

func (client *Client) DescribeLoadBalancerAttributeWithCallback(request *DescribeLoadBalancerAttributeRequest, callback func(response *DescribeLoadBalancerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerAttribute(request)
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

type DescribeLoadBalancerAttributeRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeLoadBalancerAttributeResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	LoadBalancerId     string `json:"LoadBalancerId" xml:"LoadBalancerId"`
	ResourceGroupId    string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	LoadBalancerName   string `json:"LoadBalancerName" xml:"LoadBalancerName"`
	LoadBalancerStatus string `json:"LoadBalancerStatus" xml:"LoadBalancerStatus"`
	RegionId           string `json:"RegionId" xml:"RegionId"`
	RegionIdAlias      string `json:"RegionIdAlias" xml:"RegionIdAlias"`
	Address            string `json:"Address" xml:"Address"`
	AddressType        string `json:"AddressType" xml:"AddressType"`
	VpcId              string `json:"VpcId" xml:"VpcId"`
	VSwitchId          string `json:"VSwitchId" xml:"VSwitchId"`
	NetworkType        string `json:"NetworkType" xml:"NetworkType"`
	InternetChargeType string `json:"InternetChargeType" xml:"InternetChargeType"`
	AutoReleaseTime    int    `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
	Bandwidth          int    `json:"Bandwidth" xml:"Bandwidth"`
	LoadBalancerSpec   string `json:"LoadBalancerSpec" xml:"LoadBalancerSpec"`
	CreateTime         string `json:"CreateTime" xml:"CreateTime"`
	CreateTimeStamp    int    `json:"CreateTimeStamp" xml:"CreateTimeStamp"`
	EndTime            string `json:"EndTime" xml:"EndTime"`
	EndTimeStamp       int    `json:"EndTimeStamp" xml:"EndTimeStamp"`
	PayType            string `json:"PayType" xml:"PayType"`
	MasterZoneId       string `json:"MasterZoneId" xml:"MasterZoneId"`
	SlaveZoneId        string `json:"SlaveZoneId" xml:"SlaveZoneId"`
	ListenerPorts      struct {
		ListenerPort []string `json:"ListenerPort" xml:"ListenerPort"`
	} `json:"ListenerPorts" xml:"ListenerPorts"`
	ListenerPortsAndProtocal struct {
		ListenerPortAndProtocal []struct {
			ListenerPort     int    `json:"ListenerPort" xml:"ListenerPort"`
			ListenerProtocal string `json:"ListenerProtocal" xml:"ListenerProtocal"`
		} `json:"ListenerPortAndProtocal" xml:"ListenerPortAndProtocal"`
	} `json:"ListenerPortsAndProtocal" xml:"ListenerPortsAndProtocal"`
	ListenerPortsAndProtocol struct {
		ListenerPortAndProtocol []struct {
			ListenerPort     int    `json:"ListenerPort" xml:"ListenerPort"`
			ListenerProtocol string `json:"ListenerProtocol" xml:"ListenerProtocol"`
		} `json:"ListenerPortAndProtocol" xml:"ListenerPortAndProtocol"`
	} `json:"ListenerPortsAndProtocol" xml:"ListenerPortsAndProtocol"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `json:"ServerId" xml:"ServerId"`
			Weight   int    `json:"Weight" xml:"Weight"`
		} `json:"BackendServer" xml:"BackendServer"`
	} `json:"BackendServers" xml:"BackendServers"`
}

func CreateDescribeLoadBalancerAttributeRequest() (request *DescribeLoadBalancerAttributeRequest) {
	request = &DescribeLoadBalancerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerAttribute", "slb", "openAPI")
	return
}

func CreateDescribeLoadBalancerAttributeResponse() (response *DescribeLoadBalancerAttributeResponse) {
	response = &DescribeLoadBalancerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
