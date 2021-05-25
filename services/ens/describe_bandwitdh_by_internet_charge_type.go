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

// DescribeBandwitdhByInternetChargeType invokes the ens.DescribeBandwitdhByInternetChargeType API synchronously
func (client *Client) DescribeBandwitdhByInternetChargeType(request *DescribeBandwitdhByInternetChargeTypeRequest) (response *DescribeBandwitdhByInternetChargeTypeResponse, err error) {
	response = CreateDescribeBandwitdhByInternetChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBandwitdhByInternetChargeTypeWithChan invokes the ens.DescribeBandwitdhByInternetChargeType API asynchronously
func (client *Client) DescribeBandwitdhByInternetChargeTypeWithChan(request *DescribeBandwitdhByInternetChargeTypeRequest) (<-chan *DescribeBandwitdhByInternetChargeTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeBandwitdhByInternetChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBandwitdhByInternetChargeType(request)
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

// DescribeBandwitdhByInternetChargeTypeWithCallback invokes the ens.DescribeBandwitdhByInternetChargeType API asynchronously
func (client *Client) DescribeBandwitdhByInternetChargeTypeWithCallback(request *DescribeBandwitdhByInternetChargeTypeRequest, callback func(response *DescribeBandwitdhByInternetChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBandwitdhByInternetChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeBandwitdhByInternetChargeType(request)
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

// DescribeBandwitdhByInternetChargeTypeRequest is the request struct for api DescribeBandwitdhByInternetChargeType
type DescribeBandwitdhByInternetChargeTypeRequest struct {
	*requests.RpcRequest
	Isp         string `position:"Query" name:"Isp"`
	StartTime   string `position:"Query" name:"StartTime"`
	EnsRegionId string `position:"Query" name:"EnsRegionId"`
	EndTime     string `position:"Query" name:"EndTime"`
}

// DescribeBandwitdhByInternetChargeTypeResponse is the response struct for api DescribeBandwitdhByInternetChargeType
type DescribeBandwitdhByInternetChargeTypeResponse struct {
	*responses.BaseResponse
	BandwidthValue     int64  `json:"BandwidthValue" xml:"BandwidthValue"`
	InternetChargeType string `json:"InternetChargeType" xml:"InternetChargeType"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	TimeStamp          string `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateDescribeBandwitdhByInternetChargeTypeRequest creates a request to invoke DescribeBandwitdhByInternetChargeType API
func CreateDescribeBandwitdhByInternetChargeTypeRequest() (request *DescribeBandwitdhByInternetChargeTypeRequest) {
	request = &DescribeBandwitdhByInternetChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeBandwitdhByInternetChargeType", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBandwitdhByInternetChargeTypeResponse creates a response to parse from DescribeBandwitdhByInternetChargeType response
func CreateDescribeBandwitdhByInternetChargeTypeResponse() (response *DescribeBandwitdhByInternetChargeTypeResponse) {
	response = &DescribeBandwitdhByInternetChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
