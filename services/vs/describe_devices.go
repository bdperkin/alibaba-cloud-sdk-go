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

// DescribeDevices invokes the vs.DescribeDevices API synchronously
func (client *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
	response = CreateDescribeDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDevicesWithChan invokes the vs.DescribeDevices API asynchronously
func (client *Client) DescribeDevicesWithChan(request *DescribeDevicesRequest) (<-chan *DescribeDevicesResponse, <-chan error) {
	responseChan := make(chan *DescribeDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDevices(request)
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

// DescribeDevicesWithCallback invokes the vs.DescribeDevices API asynchronously
func (client *Client) DescribeDevicesWithCallback(request *DescribeDevicesRequest, callback func(response *DescribeDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDevicesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDevices(request)
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

// DescribeDevicesRequest is the request struct for api DescribeDevices
type DescribeDevicesRequest struct {
	*requests.RpcRequest
	SortDirection    string           `position:"Query" name:"SortDirection"`
	IncludeDirectory requests.Boolean `position:"Query" name:"IncludeDirectory"`
	GbId             string           `position:"Query" name:"GbId"`
	Type             string           `position:"Query" name:"Type"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	ParentId         string           `position:"Query" name:"ParentId"`
	IncludeStats     requests.Boolean `position:"Query" name:"IncludeStats"`
	Vendor           string           `position:"Query" name:"Vendor"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	DirectoryId      string           `position:"Query" name:"DirectoryId"`
	Id               string           `position:"Query" name:"Id"`
	ShowLog          string           `position:"Query" name:"ShowLog"`
	GroupId          string           `position:"Query" name:"GroupId"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	Name             string           `position:"Query" name:"Name"`
	SortBy           string           `position:"Query" name:"SortBy"`
	Status           string           `position:"Query" name:"Status"`
}

// DescribeDevicesResponse is the response struct for api DescribeDevices
type DescribeDevicesResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageSize   int64    `json:"PageSize" xml:"PageSize"`
	PageNum    int64    `json:"PageNum" xml:"PageNum"`
	PageCount  int64    `json:"PageCount" xml:"PageCount"`
	TotalCount int64    `json:"TotalCount" xml:"TotalCount"`
	Devices    []Device `json:"Devices" xml:"Devices"`
}

// CreateDescribeDevicesRequest creates a request to invoke DescribeDevices API
func CreateDescribeDevicesRequest() (request *DescribeDevicesRequest) {
	request = &DescribeDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDevicesResponse creates a response to parse from DescribeDevices response
func CreateDescribeDevicesResponse() (response *DescribeDevicesResponse) {
	response = &DescribeDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
