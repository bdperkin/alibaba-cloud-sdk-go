package ram

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

func (client *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
	response = CreateGetUserResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetUserWithChan(request *GetUserRequest) (<-chan *GetUserResponse, <-chan error) {
	responseChan := make(chan *GetUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUser(request)
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

func (client *Client) GetUserWithCallback(request *GetUserRequest, callback func(response *GetUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserResponse
		var err error
		defer close(result)
		response, err = client.GetUser(request)
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

type GetUserRequest struct {
	*requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

type GetUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	User      struct {
		UserId        string `json:"UserId" xml:"UserId"`
		UserName      string `json:"UserName" xml:"UserName"`
		DisplayName   string `json:"DisplayName" xml:"DisplayName"`
		MobilePhone   string `json:"MobilePhone" xml:"MobilePhone"`
		Email         string `json:"Email" xml:"Email"`
		Comments      string `json:"Comments" xml:"Comments"`
		CreateDate    string `json:"CreateDate" xml:"CreateDate"`
		UpdateDate    string `json:"UpdateDate" xml:"UpdateDate"`
		LastLoginDate string `json:"LastLoginDate" xml:"LastLoginDate"`
	} `json:"User" xml:"User"`
}

func CreateGetUserRequest() (request *GetUserRequest) {
	request = &GetUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "GetUser", "", "")
	return
}

func CreateGetUserResponse() (response *GetUserResponse) {
	response = &GetUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
