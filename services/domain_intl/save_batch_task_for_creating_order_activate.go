package domain_intl

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

func (client *Client) SaveBatchTaskForCreatingOrderActivate(request *SaveBatchTaskForCreatingOrderActivateRequest) (response *SaveBatchTaskForCreatingOrderActivateResponse, err error) {
	response = CreateSaveBatchTaskForCreatingOrderActivateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveBatchTaskForCreatingOrderActivateWithChan(request *SaveBatchTaskForCreatingOrderActivateRequest) (<-chan *SaveBatchTaskForCreatingOrderActivateResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForCreatingOrderActivateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForCreatingOrderActivate(request)
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

func (client *Client) SaveBatchTaskForCreatingOrderActivateWithCallback(request *SaveBatchTaskForCreatingOrderActivateRequest, callback func(response *SaveBatchTaskForCreatingOrderActivateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForCreatingOrderActivateResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForCreatingOrderActivate(request)
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

type SaveBatchTaskForCreatingOrderActivateRequest struct {
	*requests.RpcRequest
	UserClientIp       string                                                     `position:"Query" name:"UserClientIp"`
	Lang               string                                                     `position:"Query" name:"Lang"`
	OrderActivateParam *[]SaveBatchTaskForCreatingOrderActivateOrderActivateParam `position:"Query" name:"OrderActivateParam"  type:"Repeated"`
}

type SaveBatchTaskForCreatingOrderActivateOrderActivateParam struct {
	DomainName              string `name:"DomainName"`
	SubscriptionDuration    string `name:"SubscriptionDuration"`
	RegistrantProfileId     string `name:"RegistrantProfileId"`
	EnableDomainProxy       string `name:"EnableDomainProxy"`
	PermitPremiumActivation string `name:"PermitPremiumActivation"`
}

type SaveBatchTaskForCreatingOrderActivateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

func CreateSaveBatchTaskForCreatingOrderActivateRequest() (request *SaveBatchTaskForCreatingOrderActivateRequest) {
	request = &SaveBatchTaskForCreatingOrderActivateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForCreatingOrderActivate", "", "")
	return
}

func CreateSaveBatchTaskForCreatingOrderActivateResponse() (response *SaveBatchTaskForCreatingOrderActivateResponse) {
	response = &SaveBatchTaskForCreatingOrderActivateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
