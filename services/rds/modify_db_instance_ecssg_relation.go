package rds

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

func (client *Client) ModifyDBInstanceECSSGRelation(request *ModifyDBInstanceECSSGRelationRequest) (response *ModifyDBInstanceECSSGRelationResponse, err error) {
	response = CreateModifyDBInstanceECSSGRelationResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyDBInstanceECSSGRelationWithChan(request *ModifyDBInstanceECSSGRelationRequest) (<-chan *ModifyDBInstanceECSSGRelationResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceECSSGRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceECSSGRelation(request)
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

func (client *Client) ModifyDBInstanceECSSGRelationWithCallback(request *ModifyDBInstanceECSSGRelationRequest, callback func(response *ModifyDBInstanceECSSGRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceECSSGRelationResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceECSSGRelation(request)
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

type ModifyDBInstanceECSSGRelationRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	EcsSecurityGroupId   string           `position:"Query" name:"EcsSecurityGroupId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type ModifyDBInstanceECSSGRelationResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	Items          struct {
		RdsEcsSecurityGroupRel []struct {
			RegionId        string `json:"RegionId" xml:"RegionId"`
			SecurityGroupId string `json:"SecurityGroupId" xml:"SecurityGroupId"`
			NetType         string `json:"NetType" xml:"NetType"`
		} `json:"RdsEcsSecurityGroupRel" xml:"RdsEcsSecurityGroupRel"`
	} `json:"Items" xml:"Items"`
}

func CreateModifyDBInstanceECSSGRelationRequest() (request *ModifyDBInstanceECSSGRelationRequest) {
	request = &ModifyDBInstanceECSSGRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceECSSGRelation", "rds", "openAPI")
	return
}

func CreateModifyDBInstanceECSSGRelationResponse() (response *ModifyDBInstanceECSSGRelationResponse) {
	response = &ModifyDBInstanceECSSGRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
