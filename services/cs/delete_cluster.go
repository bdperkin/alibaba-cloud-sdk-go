package cs

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

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
	response = CreateDeleteClusterResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteClusterWithChan(request *DeleteClusterRequest) (<-chan *DeleteClusterResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCluster(request)
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

func (client *Client) DeleteClusterWithCallback(request *DeleteClusterRequest, callback func(response *DeleteClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterResponse
		var err error
		defer close(result)
		response, err = client.DeleteCluster(request)
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

type DeleteClusterRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

type DeleteClusterResponse struct {
	*responses.BaseResponse
}

func CreateDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DeleteCluster", "/clusters/[ClusterId]", "", "")
	request.Method = requests.DELETE
	return
}

func CreateDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
