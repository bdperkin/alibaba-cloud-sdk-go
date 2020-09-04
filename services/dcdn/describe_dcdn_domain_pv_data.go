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

// DescribeDcdnDomainPvData invokes the dcdn.DescribeDcdnDomainPvData API synchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomainpvdata.html
func (client *Client) DescribeDcdnDomainPvData(request *DescribeDcdnDomainPvDataRequest) (response *DescribeDcdnDomainPvDataResponse, err error) {
	response = CreateDescribeDcdnDomainPvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainPvDataWithChan invokes the dcdn.DescribeDcdnDomainPvData API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomainpvdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnDomainPvDataWithChan(request *DescribeDcdnDomainPvDataRequest) (<-chan *DescribeDcdnDomainPvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainPvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainPvData(request)
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

// DescribeDcdnDomainPvDataWithCallback invokes the dcdn.DescribeDcdnDomainPvData API asynchronously
// api document: https://help.aliyun.com/api/dcdn/describedcdndomainpvdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDcdnDomainPvDataWithCallback(request *DescribeDcdnDomainPvDataRequest, callback func(response *DescribeDcdnDomainPvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainPvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainPvData(request)
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

// DescribeDcdnDomainPvDataRequest is the request struct for api DescribeDcdnDomainPvData
type DescribeDcdnDomainPvDataRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnDomainPvDataResponse is the response struct for api DescribeDcdnDomainPvData
type DescribeDcdnDomainPvDataResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	DomainName     string         `json:"DomainName" xml:"DomainName"`
	DataInterval   string         `json:"DataInterval" xml:"DataInterval"`
	StartTime      string         `json:"StartTime" xml:"StartTime"`
	EndTime        string         `json:"EndTime" xml:"EndTime"`
	PvDataInterval PvDataInterval `json:"PvDataInterval" xml:"PvDataInterval"`
}

// CreateDescribeDcdnDomainPvDataRequest creates a request to invoke DescribeDcdnDomainPvData API
func CreateDescribeDcdnDomainPvDataRequest() (request *DescribeDcdnDomainPvDataRequest) {
	request = &DescribeDcdnDomainPvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainPvData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainPvDataResponse creates a response to parse from DescribeDcdnDomainPvData response
func CreateDescribeDcdnDomainPvDataResponse() (response *DescribeDcdnDomainPvDataResponse) {
	response = &DescribeDcdnDomainPvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
