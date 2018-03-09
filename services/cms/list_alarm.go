package cms

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

// ListAlarm invokes the cms.ListAlarm API synchronously
// api document: https://help.aliyun.com/api/cms/listalarm.html
func (client *Client) ListAlarm(request *ListAlarmRequest) (response *ListAlarmResponse, err error) {
	response = CreateListAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// ListAlarmWithChan invokes the cms.ListAlarm API asynchronously
// api document: https://help.aliyun.com/api/cms/listalarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAlarmWithChan(request *ListAlarmRequest) (<-chan *ListAlarmResponse, <-chan error) {
	responseChan := make(chan *ListAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAlarm(request)
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

// ListAlarmWithCallback invokes the cms.ListAlarm API asynchronously
// api document: https://help.aliyun.com/api/cms/listalarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAlarmWithCallback(request *ListAlarmRequest, callback func(response *ListAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAlarmResponse
		var err error
		defer close(result)
		response, err = client.ListAlarm(request)
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

// ListAlarmRequest is the request struct for api ListAlarm
type ListAlarmRequest struct {
	*requests.RpcRequest
	CallbyCmsOwner string           `position:"Query" name:"callby_cms_owner"`
	Id             string           `position:"Query" name:"Id"`
	Name           string           `position:"Query" name:"Name"`
	Namespace      string           `position:"Query" name:"Namespace"`
	Dimension      string           `position:"Query" name:"Dimension"`
	State          string           `position:"Query" name:"State"`
	IsEnable       requests.Boolean `position:"Query" name:"IsEnable"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// ListAlarmResponse is the response struct for api ListAlarm
type ListAlarmResponse struct {
	*responses.BaseResponse
	Success   bool      `json:"Success" xml:"Success"`
	Code      string    `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	NextToken int       `json:"NextToken" xml:"NextToken"`
	Total     int       `json:"Total" xml:"Total"`
	RequestId string    `json:"RequestId" xml:"RequestId"`
	AlarmList AlarmList `json:"AlarmList" xml:"AlarmList"`
}

// CreateListAlarmRequest creates a request to invoke ListAlarm API
func CreateListAlarmRequest() (request *ListAlarmRequest) {
	request = &ListAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ListAlarm", "cms", "openAPI")
	return
}

// CreateListAlarmResponse creates a response to parse from ListAlarm response
func CreateListAlarmResponse() (response *ListAlarmResponse) {
	response = &ListAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
