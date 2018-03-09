package slb

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

// DescribeAccessLogsDownloadAttribute invokes the slb.DescribeAccessLogsDownloadAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describeaccesslogsdownloadattribute.html
func (client *Client) DescribeAccessLogsDownloadAttribute(request *DescribeAccessLogsDownloadAttributeRequest) (response *DescribeAccessLogsDownloadAttributeResponse, err error) {
	response = CreateDescribeAccessLogsDownloadAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccessLogsDownloadAttributeWithChan invokes the slb.DescribeAccessLogsDownloadAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeaccesslogsdownloadattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccessLogsDownloadAttributeWithChan(request *DescribeAccessLogsDownloadAttributeRequest) (<-chan *DescribeAccessLogsDownloadAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeAccessLogsDownloadAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccessLogsDownloadAttribute(request)
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

// DescribeAccessLogsDownloadAttributeWithCallback invokes the slb.DescribeAccessLogsDownloadAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describeaccesslogsdownloadattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccessLogsDownloadAttributeWithCallback(request *DescribeAccessLogsDownloadAttributeRequest, callback func(response *DescribeAccessLogsDownloadAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccessLogsDownloadAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccessLogsDownloadAttribute(request)
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

// DescribeAccessLogsDownloadAttributeRequest is the request struct for api DescribeAccessLogsDownloadAttribute
type DescribeAccessLogsDownloadAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	LogType              string           `position:"Query" name:"LogType"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeAccessLogsDownloadAttributeResponse is the response struct for api DescribeAccessLogsDownloadAttribute
type DescribeAccessLogsDownloadAttributeResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	PageNumber             int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize               int                    `json:"PageSize" xml:"PageSize"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	Count                  int                    `json:"Count" xml:"Count"`
	LogsDownloadAttributes LogsDownloadAttributes `json:"LogsDownloadAttributes" xml:"LogsDownloadAttributes"`
}

// CreateDescribeAccessLogsDownloadAttributeRequest creates a request to invoke DescribeAccessLogsDownloadAttribute API
func CreateDescribeAccessLogsDownloadAttributeRequest() (request *DescribeAccessLogsDownloadAttributeRequest) {
	request = &DescribeAccessLogsDownloadAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeAccessLogsDownloadAttribute", "slb", "openAPI")
	return
}

// CreateDescribeAccessLogsDownloadAttributeResponse creates a response to parse from DescribeAccessLogsDownloadAttribute response
func CreateDescribeAccessLogsDownloadAttributeResponse() (response *DescribeAccessLogsDownloadAttributeResponse) {
	response = &DescribeAccessLogsDownloadAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
