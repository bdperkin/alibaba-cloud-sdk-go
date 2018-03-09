package ddospro

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

// DescribeDomainConfig invokes the ddospro.DescribeDomainConfig API synchronously
// api document: https://help.aliyun.com/api/ddospro/describedomainconfig.html
func (client *Client) DescribeDomainConfig(request *DescribeDomainConfigRequest) (response *DescribeDomainConfigResponse, err error) {
	response = CreateDescribeDomainConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainConfigWithChan invokes the ddospro.DescribeDomainConfig API asynchronously
// api document: https://help.aliyun.com/api/ddospro/describedomainconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainConfigWithChan(request *DescribeDomainConfigRequest) (<-chan *DescribeDomainConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainConfig(request)
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

// DescribeDomainConfigWithCallback invokes the ddospro.DescribeDomainConfig API asynchronously
// api document: https://help.aliyun.com/api/ddospro/describedomainconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainConfigWithCallback(request *DescribeDomainConfigRequest, callback func(response *DescribeDomainConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainConfig(request)
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

// DescribeDomainConfigRequest is the request struct for api DescribeDomainConfig
type DescribeDomainConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Domain          string           `position:"Query" name:"Domain"`
}

// DescribeDomainConfigResponse is the response struct for api DescribeDomainConfig
type DescribeDomainConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Config    Config `json:"Config" xml:"Config"`
}

// CreateDescribeDomainConfigRequest creates a request to invoke DescribeDomainConfig API
func CreateDescribeDomainConfigRequest() (request *DescribeDomainConfigRequest) {
	request = &DescribeDomainConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeDomainConfig", "", "")
	return
}

// CreateDescribeDomainConfigResponse creates a response to parse from DescribeDomainConfig response
func CreateDescribeDomainConfigResponse() (response *DescribeDomainConfigResponse) {
	response = &DescribeDomainConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
