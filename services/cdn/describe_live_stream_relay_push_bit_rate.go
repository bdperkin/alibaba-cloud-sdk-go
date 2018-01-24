package cdn

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

func (client *Client) DescribeLiveStreamRelayPushBitRate(request *DescribeLiveStreamRelayPushBitRateRequest) (response *DescribeLiveStreamRelayPushBitRateResponse, err error) {
	response = CreateDescribeLiveStreamRelayPushBitRateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamRelayPushBitRateWithChan(request *DescribeLiveStreamRelayPushBitRateRequest) (<-chan *DescribeLiveStreamRelayPushBitRateResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRelayPushBitRateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRelayPushBitRate(request)
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

func (client *Client) DescribeLiveStreamRelayPushBitRateWithCallback(request *DescribeLiveStreamRelayPushBitRateRequest, callback func(response *DescribeLiveStreamRelayPushBitRateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRelayPushBitRateResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRelayPushBitRate(request)
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

type DescribeLiveStreamRelayPushBitRateRequest struct {
	*requests.RpcRequest
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamRelayPushBitRateResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	RelayPushBitRateModelList struct {
		RelayPushBitRateModel []struct {
			Time          string `json:"Time" xml:"Time"`
			VedioFrame    string `json:"VedioFrame" xml:"VedioFrame"`
			VedioTimstamp string `json:"VedioTimstamp" xml:"VedioTimstamp"`
			AudioFrame    string `json:"AudioFrame" xml:"AudioFrame"`
			AudioTimstamp string `json:"AudioTimstamp" xml:"AudioTimstamp"`
			RelayDomain   string `json:"RelayDomain" xml:"RelayDomain"`
		} `json:"RelayPushBitRateModel" xml:"RelayPushBitRateModel"`
	} `json:"RelayPushBitRateModelList" xml:"RelayPushBitRateModelList"`
}

func CreateDescribeLiveStreamRelayPushBitRateRequest() (request *DescribeLiveStreamRelayPushBitRateRequest) {
	request = &DescribeLiveStreamRelayPushBitRateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRelayPushBitRate", "", "")
	return
}

func CreateDescribeLiveStreamRelayPushBitRateResponse() (response *DescribeLiveStreamRelayPushBitRateResponse) {
	response = &DescribeLiveStreamRelayPushBitRateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
