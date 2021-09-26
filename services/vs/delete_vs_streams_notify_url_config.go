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

// DeleteVsStreamsNotifyUrlConfig invokes the vs.DeleteVsStreamsNotifyUrlConfig API synchronously
func (client *Client) DeleteVsStreamsNotifyUrlConfig(request *DeleteVsStreamsNotifyUrlConfigRequest) (response *DeleteVsStreamsNotifyUrlConfigResponse, err error) {
	response = CreateDeleteVsStreamsNotifyUrlConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVsStreamsNotifyUrlConfigWithChan invokes the vs.DeleteVsStreamsNotifyUrlConfig API asynchronously
func (client *Client) DeleteVsStreamsNotifyUrlConfigWithChan(request *DeleteVsStreamsNotifyUrlConfigRequest) (<-chan *DeleteVsStreamsNotifyUrlConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteVsStreamsNotifyUrlConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVsStreamsNotifyUrlConfig(request)
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

// DeleteVsStreamsNotifyUrlConfigWithCallback invokes the vs.DeleteVsStreamsNotifyUrlConfig API asynchronously
func (client *Client) DeleteVsStreamsNotifyUrlConfigWithCallback(request *DeleteVsStreamsNotifyUrlConfigRequest, callback func(response *DeleteVsStreamsNotifyUrlConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVsStreamsNotifyUrlConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteVsStreamsNotifyUrlConfig(request)
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

// DeleteVsStreamsNotifyUrlConfigRequest is the request struct for api DeleteVsStreamsNotifyUrlConfig
type DeleteVsStreamsNotifyUrlConfigRequest struct {
	*requests.RpcRequest
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteVsStreamsNotifyUrlConfigResponse is the response struct for api DeleteVsStreamsNotifyUrlConfig
type DeleteVsStreamsNotifyUrlConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVsStreamsNotifyUrlConfigRequest creates a request to invoke DeleteVsStreamsNotifyUrlConfig API
func CreateDeleteVsStreamsNotifyUrlConfigRequest() (request *DeleteVsStreamsNotifyUrlConfigRequest) {
	request = &DeleteVsStreamsNotifyUrlConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DeleteVsStreamsNotifyUrlConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteVsStreamsNotifyUrlConfigResponse creates a response to parse from DeleteVsStreamsNotifyUrlConfig response
func CreateDeleteVsStreamsNotifyUrlConfigResponse() (response *DeleteVsStreamsNotifyUrlConfigResponse) {
	response = &DeleteVsStreamsNotifyUrlConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
