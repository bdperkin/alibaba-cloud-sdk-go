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

// DescribeAlarmHistory invokes the cms.DescribeAlarmHistory API synchronously
// api document: https://help.aliyun.com/api/cms/describealarmhistory.html
func (client *Client) DescribeAlarmHistory(request *DescribeAlarmHistoryRequest) (response *DescribeAlarmHistoryResponse, err error) {
	response = CreateDescribeAlarmHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlarmHistoryWithChan invokes the cms.DescribeAlarmHistory API asynchronously
// api document: https://help.aliyun.com/api/cms/describealarmhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmHistoryWithChan(request *DescribeAlarmHistoryRequest) (<-chan *DescribeAlarmHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeAlarmHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlarmHistory(request)
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

// DescribeAlarmHistoryWithCallback invokes the cms.DescribeAlarmHistory API asynchronously
// api document: https://help.aliyun.com/api/cms/describealarmhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAlarmHistoryWithCallback(request *DescribeAlarmHistoryRequest, callback func(response *DescribeAlarmHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlarmHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlarmHistory(request)
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

// DescribeAlarmHistoryRequest is the request struct for api DescribeAlarmHistory
type DescribeAlarmHistoryRequest struct {
	*requests.RpcRequest
	AlertName  string           `position:"Query" name:"AlertName"`
	RuleName   string           `position:"Query" name:"RuleName"`
	Namespace  string           `position:"Query" name:"Namespace"`
	MetricName string           `position:"Query" name:"MetricName"`
	GroupId    string           `position:"Query" name:"GroupId"`
	Status     string           `position:"Query" name:"Status"`
	State      string           `position:"Query" name:"State"`
	Ascending  requests.Boolean `position:"Query" name:"Ascending"`
	OnlyCount  requests.Boolean `position:"Query" name:"OnlyCount"`
	StartTime  string           `position:"Query" name:"StartTime"`
	EndTime    string           `position:"Query" name:"EndTime"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Page       requests.Integer `position:"Query" name:"Page"`
}

// DescribeAlarmHistoryResponse is the response struct for api DescribeAlarmHistory
type DescribeAlarmHistoryResponse struct {
	*responses.BaseResponse
	Success          bool                                   `json:"Success" xml:"Success"`
	Code             string                                 `json:"Code" xml:"Code"`
	Message          string                                 `json:"Message" xml:"Message"`
	Total            string                                 `json:"Total" xml:"Total"`
	RequestId        string                                 `json:"RequestId" xml:"RequestId"`
	AlarmHistoryList AlarmHistoryListInDescribeAlarmHistory `json:"AlarmHistoryList" xml:"AlarmHistoryList"`
}

// CreateDescribeAlarmHistoryRequest creates a request to invoke DescribeAlarmHistory API
func CreateDescribeAlarmHistoryRequest() (request *DescribeAlarmHistoryRequest) {
	request = &DescribeAlarmHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "DescribeAlarmHistory", "cms", "openAPI")
	return
}

// CreateDescribeAlarmHistoryResponse creates a response to parse from DescribeAlarmHistory response
func CreateDescribeAlarmHistoryResponse() (response *DescribeAlarmHistoryResponse) {
	response = &DescribeAlarmHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
