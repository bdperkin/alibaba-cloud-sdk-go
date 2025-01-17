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

// UpdateFee invokes the companyreg.UpdateFee API synchronously
func (client *Client) UpdateFee(request *UpdateFeeRequest) (response *UpdateFeeResponse, err error) {
	response = CreateUpdateFeeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFeeWithChan invokes the companyreg.UpdateFee API asynchronously
func (client *Client) UpdateFeeWithChan(request *UpdateFeeRequest) (<-chan *UpdateFeeResponse, <-chan error) {
	responseChan := make(chan *UpdateFeeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFee(request)
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

// UpdateFeeWithCallback invokes the companyreg.UpdateFee API asynchronously
func (client *Client) UpdateFeeWithCallback(request *UpdateFeeRequest, callback func(response *UpdateFeeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFeeResponse
		var err error
		defer close(result)
		response, err = client.UpdateFee(request)
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

// UpdateFeeRequest is the request struct for api UpdateFee
type UpdateFeeRequest struct {
	*requests.RpcRequest
	Kind            string           `position:"Query" name:"Kind"`
	Use             string           `position:"Query" name:"Use"`
	BaseTotalAmount string           `position:"Query" name:"BaseTotalAmount"`
	Payer           string           `position:"Query" name:"Payer"`
	SecondKey       string           `position:"Query" name:"SecondKey"`
	PayMethod       string           `position:"Query" name:"PayMethod"`
	FirstKey        string           `position:"Query" name:"FirstKey"`
	BizId           string           `position:"Query" name:"BizId"`
	FeeType         string           `position:"Query" name:"FeeType"`
	Id              requests.Integer `position:"Query" name:"Id"`
}

// UpdateFeeResponse is the response struct for api UpdateFee
type UpdateFeeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateUpdateFeeRequest creates a request to invoke UpdateFee API
func CreateUpdateFeeRequest() (request *UpdateFeeRequest) {
	request = &UpdateFeeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "UpdateFee", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateFeeResponse creates a response to parse from UpdateFee response
func CreateUpdateFeeResponse() (response *UpdateFeeResponse) {
	response = &UpdateFeeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
