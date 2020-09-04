package dcdn

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

// DescribeDcdnDomainRealTimeReqHitRateData invokes the dcdn.DescribeDcdnDomainRealTimeReqHitRateData API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomainrealtimereqhitratedata.html
func (client *Client) DescribeDcdnDomainRealTimeReqHitRateData(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest) (response *DescribeDcdnDomainRealTimeReqHitRateDataResponse, err error) {
	response = CreateDescribeDcdnDomainRealTimeReqHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainRealTimeReqHitRateDataWithChan invokes the dcdn.DescribeDcdnDomainRealTimeReqHitRateData API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomainrealtimereqhitratedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnDomainRealTimeReqHitRateDataWithChan(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest) (<-chan *DescribeDcdnDomainRealTimeReqHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainRealTimeReqHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainRealTimeReqHitRateData(request)
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

// DescribeDcdnDomainRealTimeReqHitRateDataWithCallback invokes the dcdn.DescribeDcdnDomainRealTimeReqHitRateData API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomainrealtimereqhitratedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnDomainRealTimeReqHitRateDataWithCallback(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest, callback func(response *DescribeDcdnDomainRealTimeReqHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainRealTimeReqHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainRealTimeReqHitRateData(request)
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

// DescribeDcdnDomainRealTimeReqHitRateDataRequest is the request struct for api DescribeDcdnDomainRealTimeReqHitRateData
type DescribeDcdnDomainRealTimeReqHitRateDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnDomainRealTimeReqHitRateDataResponse is the response struct for api DescribeDcdnDomainRealTimeReqHitRateData
type DescribeDcdnDomainRealTimeReqHitRateDataResponse struct {
	*responses.BaseResponse
	RequestId string                                         `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDcdnDomainRealTimeReqHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeDcdnDomainRealTimeReqHitRateDataRequest creates a request to invoke DescribeDcdnDomainRealTimeReqHitRateData API
func CreateDescribeDcdnDomainRealTimeReqHitRateDataRequest() (request *DescribeDcdnDomainRealTimeReqHitRateDataRequest) {
	request = &DescribeDcdnDomainRealTimeReqHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainRealTimeReqHitRateData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDcdnDomainRealTimeReqHitRateDataResponse creates a response to parse from DescribeDcdnDomainRealTimeReqHitRateData response
func CreateDescribeDcdnDomainRealTimeReqHitRateDataResponse() (response *DescribeDcdnDomainRealTimeReqHitRateDataResponse) {
	response = &DescribeDcdnDomainRealTimeReqHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
