package vs

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

// DescribePresets invokes the vs.DescribePresets API synchronously
func (client *Client) DescribePresets(request *DescribePresetsRequest) (response *DescribePresetsResponse, err error) {
	response = CreateDescribePresetsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePresetsWithChan invokes the vs.DescribePresets API asynchronously
func (client *Client) DescribePresetsWithChan(request *DescribePresetsRequest) (<-chan *DescribePresetsResponse, <-chan error) {
	responseChan := make(chan *DescribePresetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePresets(request)
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

// DescribePresetsWithCallback invokes the vs.DescribePresets API asynchronously
func (client *Client) DescribePresetsWithCallback(request *DescribePresetsRequest, callback func(response *DescribePresetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePresetsResponse
		var err error
		defer close(result)
		response, err = client.DescribePresets(request)
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

// DescribePresetsRequest is the request struct for api DescribePresets
type DescribePresetsRequest struct {
	*requests.RpcRequest
	SubProtocol string           `position:"Query" name:"SubProtocol"`
	Id          string           `position:"Query" name:"Id"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribePresetsResponse is the response struct for api DescribePresets
type DescribePresetsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Id        string   `json:"Id" xml:"Id"`
	Presets   []Preset `json:"Presets" xml:"Presets"`
}

// CreateDescribePresetsRequest creates a request to invoke DescribePresets API
func CreateDescribePresetsRequest() (request *DescribePresetsRequest) {
	request = &DescribePresetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribePresets", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePresetsResponse creates a response to parse from DescribePresets response
func CreateDescribePresetsResponse() (response *DescribePresetsResponse) {
	response = &DescribePresetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
