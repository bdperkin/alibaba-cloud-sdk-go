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

// SetDcdnConfigOfVersion invokes the dcdn.SetDcdnConfigOfVersion API synchronously
// api document: https://help.aliyun.com/api/dcdn/setdcdnconfigofversion.html
func (client *Client) SetDcdnConfigOfVersion(request *SetDcdnConfigOfVersionRequest) (response *SetDcdnConfigOfVersionResponse, err error) {
	response = CreateSetDcdnConfigOfVersionResponse()
	err = client.DoAction(request, response)
	return
}

// SetDcdnConfigOfVersionWithChan invokes the dcdn.SetDcdnConfigOfVersion API asynchronously
// api document: https://help.aliyun.com/api/dcdn/setdcdnconfigofversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDcdnConfigOfVersionWithChan(request *SetDcdnConfigOfVersionRequest) (<-chan *SetDcdnConfigOfVersionResponse, <-chan error) {
	responseChan := make(chan *SetDcdnConfigOfVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDcdnConfigOfVersion(request)
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

// SetDcdnConfigOfVersionWithCallback invokes the dcdn.SetDcdnConfigOfVersion API asynchronously
// api document: https://help.aliyun.com/api/dcdn/setdcdnconfigofversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDcdnConfigOfVersionWithCallback(request *SetDcdnConfigOfVersionRequest, callback func(response *SetDcdnConfigOfVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDcdnConfigOfVersionResponse
		var err error
		defer close(result)
		response, err = client.SetDcdnConfigOfVersion(request)
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

// SetDcdnConfigOfVersionRequest is the request struct for api SetDcdnConfigOfVersion
type SetDcdnConfigOfVersionRequest struct {
	*requests.RpcRequest
	VersionId     string           `position:"Query" name:"VersionId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	FunctionName  string           `position:"Query" name:"FunctionName"`
	FunctionArgs  string           `position:"Query" name:"FunctionArgs"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	FunctionId    requests.Integer `position:"Query" name:"FunctionId"`
	ConfigId      string           `position:"Query" name:"ConfigId"`
}

// SetDcdnConfigOfVersionResponse is the response struct for api SetDcdnConfigOfVersion
type SetDcdnConfigOfVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDcdnConfigOfVersionRequest creates a request to invoke SetDcdnConfigOfVersion API
func CreateSetDcdnConfigOfVersionRequest() (request *SetDcdnConfigOfVersionRequest) {
	request = &SetDcdnConfigOfVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "SetDcdnConfigOfVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateSetDcdnConfigOfVersionResponse creates a response to parse from SetDcdnConfigOfVersion response
func CreateSetDcdnConfigOfVersionResponse() (response *SetDcdnConfigOfVersionResponse) {
	response = &SetDcdnConfigOfVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
