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

func (client *Client) AllocateEipAddress(request *AllocateEipAddressRequest) (response *AllocateEipAddressResponse, err error) {
	response = CreateAllocateEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AllocateEipAddressWithChan(request *AllocateEipAddressRequest) (<-chan *AllocateEipAddressResponse, <-chan error) {
	responseChan := make(chan *AllocateEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateEipAddress(request)
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

func (client *Client) AllocateEipAddressWithCallback(request *AllocateEipAddressRequest, callback func(response *AllocateEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateEipAddressResponse
		var err error
		defer close(result)
		response, err = client.AllocateEipAddress(request)
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

type AllocateEipAddressRequest struct {
	*requests.RpcRequest
	Netmode              string           `position:"Query" name:"Netmode"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
	ISP                  string           `position:"Query" name:"ISP"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Period               requests.Integer `position:"Query" name:"Period"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
}

type AllocateEipAddressResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	AllocationId string `json:"AllocationId" xml:"AllocationId"`
	EipAddress   string `json:"EipAddress" xml:"EipAddress"`
	OrderId      int    `json:"OrderId" xml:"OrderId"`
}

func CreateAllocateEipAddressRequest() (request *AllocateEipAddressRequest) {
	request = &AllocateEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AllocateEipAddress", "vpc", "openAPI")
	return
}

func CreateAllocateEipAddressResponse() (response *AllocateEipAddressResponse) {
	response = &AllocateEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
