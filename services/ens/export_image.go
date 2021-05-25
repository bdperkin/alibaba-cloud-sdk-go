package ens

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

// ExportImage invokes the ens.ExportImage API synchronously
func (client *Client) ExportImage(request *ExportImageRequest) (response *ExportImageResponse, err error) {
	response = CreateExportImageResponse()
	err = client.DoAction(request, response)
	return
}

// ExportImageWithChan invokes the ens.ExportImage API asynchronously
func (client *Client) ExportImageWithChan(request *ExportImageRequest) (<-chan *ExportImageResponse, <-chan error) {
	responseChan := make(chan *ExportImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportImage(request)
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

// ExportImageWithCallback invokes the ens.ExportImage API asynchronously
func (client *Client) ExportImageWithCallback(request *ExportImageRequest, callback func(response *ExportImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportImageResponse
		var err error
		defer close(result)
		response, err = client.ExportImage(request)
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

// ExportImageRequest is the request struct for api ExportImage
type ExportImageRequest struct {
	*requests.RpcRequest
	OSSRegionId string `position:"Query" name:"OSSRegionId"`
	OSSBucket   string `position:"Query" name:"OSSBucket"`
	RoleName    string `position:"Query" name:"RoleName"`
	OSSPrefix   string `position:"Query" name:"OSSPrefix"`
	ImageId     string `position:"Query" name:"ImageId"`
}

// ExportImageResponse is the response struct for api ExportImage
type ExportImageResponse struct {
	*responses.BaseResponse
	ExportedImageURL string `json:"ExportedImageURL" xml:"ExportedImageURL"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateExportImageRequest creates a request to invoke ExportImage API
func CreateExportImageRequest() (request *ExportImageRequest) {
	request = &ExportImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ExportImage", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportImageResponse creates a response to parse from ExportImage response
func CreateExportImageResponse() (response *ExportImageResponse) {
	response = &ExportImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
