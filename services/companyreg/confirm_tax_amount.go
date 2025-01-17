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

// ConfirmTaxAmount invokes the companyreg.ConfirmTaxAmount API synchronously
func (client *Client) ConfirmTaxAmount(request *ConfirmTaxAmountRequest) (response *ConfirmTaxAmountResponse, err error) {
	response = CreateConfirmTaxAmountResponse()
	err = client.DoAction(request, response)
	return
}

// ConfirmTaxAmountWithChan invokes the companyreg.ConfirmTaxAmount API asynchronously
func (client *Client) ConfirmTaxAmountWithChan(request *ConfirmTaxAmountRequest) (<-chan *ConfirmTaxAmountResponse, <-chan error) {
	responseChan := make(chan *ConfirmTaxAmountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfirmTaxAmount(request)
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

// ConfirmTaxAmountWithCallback invokes the companyreg.ConfirmTaxAmount API asynchronously
func (client *Client) ConfirmTaxAmountWithCallback(request *ConfirmTaxAmountRequest, callback func(response *ConfirmTaxAmountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfirmTaxAmountResponse
		var err error
		defer close(result)
		response, err = client.ConfirmTaxAmount(request)
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

// ConfirmTaxAmountRequest is the request struct for api ConfirmTaxAmount
type ConfirmTaxAmountRequest struct {
	*requests.RpcRequest
	BizId string `position:"Body" name:"BizId"`
}

// ConfirmTaxAmountResponse is the response struct for api ConfirmTaxAmount
type ConfirmTaxAmountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateConfirmTaxAmountRequest creates a request to invoke ConfirmTaxAmount API
func CreateConfirmTaxAmountRequest() (request *ConfirmTaxAmountRequest) {
	request = &ConfirmTaxAmountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ConfirmTaxAmount", "", "")
	request.Method = requests.POST
	return
}

// CreateConfirmTaxAmountResponse creates a response to parse from ConfirmTaxAmount response
func CreateConfirmTaxAmountResponse() (response *ConfirmTaxAmountResponse) {
	response = &ConfirmTaxAmountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
