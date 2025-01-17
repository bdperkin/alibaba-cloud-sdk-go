package companyreg

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

// GetAlipayUrl invokes the companyreg.GetAlipayUrl API synchronously
func (client *Client) GetAlipayUrl(request *GetAlipayUrlRequest) (response *GetAlipayUrlResponse, err error) {
	response = CreateGetAlipayUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetAlipayUrlWithChan invokes the companyreg.GetAlipayUrl API asynchronously
func (client *Client) GetAlipayUrlWithChan(request *GetAlipayUrlRequest) (<-chan *GetAlipayUrlResponse, <-chan error) {
	responseChan := make(chan *GetAlipayUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAlipayUrl(request)
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

// GetAlipayUrlWithCallback invokes the companyreg.GetAlipayUrl API asynchronously
func (client *Client) GetAlipayUrlWithCallback(request *GetAlipayUrlRequest, callback func(response *GetAlipayUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAlipayUrlResponse
		var err error
		defer close(result)
		response, err = client.GetAlipayUrl(request)
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

// GetAlipayUrlRequest is the request struct for api GetAlipayUrl
type GetAlipayUrlRequest struct {
	*requests.RpcRequest
	OrderId   requests.Integer `position:"Query" name:"OrderId"`
	Type      string           `position:"Query" name:"Type"`
	BizType   string           `position:"Query" name:"BizType"`
	ReturnUrl string           `position:"Query" name:"ReturnUrl"`
}

// GetAlipayUrlResponse is the response struct for api GetAlipayUrl
type GetAlipayUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetAlipayUrlRequest creates a request to invoke GetAlipayUrl API
func CreateGetAlipayUrlRequest() (request *GetAlipayUrlRequest) {
	request = &GetAlipayUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "GetAlipayUrl", "", "")
	request.Method = requests.GET
	return
}

// CreateGetAlipayUrlResponse creates a response to parse from GetAlipayUrl response
func CreateGetAlipayUrlResponse() (response *GetAlipayUrlResponse) {
	response = &GetAlipayUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
