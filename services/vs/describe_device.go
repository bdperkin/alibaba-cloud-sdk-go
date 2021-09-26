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

// DescribeDevice invokes the vs.DescribeDevice API synchronously
func (client *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
	response = CreateDescribeDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeviceWithChan invokes the vs.DescribeDevice API asynchronously
func (client *Client) DescribeDeviceWithChan(request *DescribeDeviceRequest) (<-chan *DescribeDeviceResponse, <-chan error) {
	responseChan := make(chan *DescribeDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDevice(request)
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

// DescribeDeviceWithCallback invokes the vs.DescribeDevice API asynchronously
func (client *Client) DescribeDeviceWithCallback(request *DescribeDeviceRequest, callback func(response *DescribeDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeviceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDevice(request)
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

// DescribeDeviceRequest is the request struct for api DescribeDevice
type DescribeDeviceRequest struct {
	*requests.RpcRequest
	IncludeDirectory requests.Boolean `position:"Query" name:"IncludeDirectory"`
	IncludeStats     requests.Boolean `position:"Query" name:"IncludeStats"`
	Id               string           `position:"Query" name:"Id"`
	ShowLog          string           `position:"Query" name:"ShowLog"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDeviceResponse is the response struct for api DescribeDevice
type DescribeDeviceResponse struct {
	*responses.BaseResponse
	RequestId       string    `json:"RequestId" xml:"RequestId"`
	Id              string    `json:"Id" xml:"Id"`
	Name            string    `json:"Name" xml:"Name"`
	Description     string    `json:"Description" xml:"Description"`
	GroupId         string    `json:"GroupId" xml:"GroupId"`
	ParentId        string    `json:"ParentId" xml:"ParentId"`
	DirectoryId     string    `json:"DirectoryId" xml:"DirectoryId"`
	Type            string    `json:"Type" xml:"Type"`
	AutoStart       bool      `json:"AutoStart" xml:"AutoStart"`
	Enabled         bool      `json:"Enabled" xml:"Enabled"`
	Protocol        string    `json:"Protocol" xml:"Protocol"`
	Status          string    `json:"Status" xml:"Status"`
	GbId            string    `json:"GbId" xml:"GbId"`
	Ip              string    `json:"Ip" xml:"Ip"`
	Port            int64     `json:"Port" xml:"Port"`
	Url             string    `json:"Url" xml:"Url"`
	Username        string    `json:"Username" xml:"Username"`
	Password        string    `json:"Password" xml:"Password"`
	Vendor          string    `json:"Vendor" xml:"Vendor"`
	Dsn             string    `json:"Dsn" xml:"Dsn"`
	Longitude       string    `json:"Longitude" xml:"Longitude"`
	Latitude        string    `json:"Latitude" xml:"Latitude"`
	AutoPos         bool      `json:"AutoPos" xml:"AutoPos"`
	PosInterval     int64     `json:"PosInterval" xml:"PosInterval"`
	AlarmMethod     string    `json:"AlarmMethod" xml:"AlarmMethod"`
	CreatedTime     string    `json:"CreatedTime" xml:"CreatedTime"`
	RegisteredTime  string    `json:"RegisteredTime" xml:"RegisteredTime"`
	ChannelSyncTime string    `json:"ChannelSyncTime" xml:"ChannelSyncTime"`
	Params          string    `json:"Params" xml:"Params"`
	Stats           Stats     `json:"Stats" xml:"Stats"`
	Directory       Directory `json:"Directory" xml:"Directory"`
}

// CreateDescribeDeviceRequest creates a request to invoke DescribeDevice API
func CreateDescribeDeviceRequest() (request *DescribeDeviceRequest) {
	request = &DescribeDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDeviceResponse creates a response to parse from DescribeDevice response
func CreateDescribeDeviceResponse() (response *DescribeDeviceResponse) {
	response = &DescribeDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
