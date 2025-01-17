package companyreg

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

// UpdateCustomerInfo invokes the companyreg.UpdateCustomerInfo API synchronously
func (client *Client) UpdateCustomerInfo(request *UpdateCustomerInfoRequest) (response *UpdateCustomerInfoResponse, err error) {
	response = CreateUpdateCustomerInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCustomerInfoWithChan invokes the companyreg.UpdateCustomerInfo API asynchronously
func (client *Client) UpdateCustomerInfoWithChan(request *UpdateCustomerInfoRequest) (<-chan *UpdateCustomerInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateCustomerInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCustomerInfo(request)
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

// UpdateCustomerInfoWithCallback invokes the companyreg.UpdateCustomerInfo API asynchronously
func (client *Client) UpdateCustomerInfoWithCallback(request *UpdateCustomerInfoRequest, callback func(response *UpdateCustomerInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCustomerInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateCustomerInfo(request)
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

// UpdateCustomerInfoRequest is the request struct for api UpdateCustomerInfo
type UpdateCustomerInfoRequest struct {
	*requests.RpcRequest
	ContactTelInfo            string           `position:"Query" name:"ContactTelInfo"`
	TaxPreparerName           string           `position:"Query" name:"TaxPreparerName"`
	TaxAgentSecret            string           `position:"Query" name:"TaxAgentSecret"`
	LegalRepresentative       string           `position:"Query" name:"LegalRepresentative"`
	TaxPreparerPassword       string           `position:"Query" name:"TaxPreparerPassword"`
	EstablishmentDate         string           `position:"Query" name:"EstablishmentDate"`
	ContactName               string           `position:"Query" name:"ContactName"`
	BizScope                  string           `position:"Query" name:"BizScope"`
	IncomeDeclarationPassword string           `position:"Query" name:"IncomeDeclarationPassword"`
	CompanyType               string           `position:"Query" name:"CompanyType"`
	CorpAddress               string           `position:"Query" name:"CorpAddress"`
	BizId                     string           `position:"Query" name:"BizId"`
	Name                      string           `position:"Query" name:"Name"`
	IsRefreshInfo             requests.Boolean `position:"Query" name:"IsRefreshInfo"`
	RegisteredCapital         string           `position:"Query" name:"RegisteredCapital"`
	OperatingPeriod           string           `position:"Query" name:"OperatingPeriod"`
}

// UpdateCustomerInfoResponse is the response struct for api UpdateCustomerInfo
type UpdateCustomerInfoResponse struct {
	*responses.BaseResponse
	AcctgSystem          string `json:"AcctgSystem" xml:"AcctgSystem"`
	BizScope             string `json:"BizScope" xml:"BizScope"`
	CompanyType          string `json:"CompanyType" xml:"CompanyType"`
	CorpAddress          string `json:"CorpAddress" xml:"CorpAddress"`
	EstablishmentDate    string `json:"EstablishmentDate" xml:"EstablishmentDate"`
	ExternalUniqueId     string `json:"ExternalUniqueId" xml:"ExternalUniqueId"`
	IncomeDeclarationPsw string `json:"IncomeDeclarationPsw" xml:"IncomeDeclarationPsw"`
	LegalRepresentative  string `json:"LegalRepresentative" xml:"LegalRepresentative"`
	Name                 string `json:"Name" xml:"Name"`
	OrgName              string `json:"OrgName" xml:"OrgName"`
	RegisteredCaptial    string `json:"RegisteredCaptial" xml:"RegisteredCaptial"`
	RequestId            string `json:"RequestId" xml:"RequestId"`
	TaxArea              string `json:"TaxArea" xml:"TaxArea"`
	TaxNo                string `json:"TaxNo" xml:"TaxNo"`
	TaxPreparerName      string `json:"TaxPreparerName" xml:"TaxPreparerName"`
	TaxPreparerPsw       string `json:"TaxPreparerPsw" xml:"TaxPreparerPsw"`
	TaxTypes             string `json:"TaxTypes" xml:"TaxTypes"`
	TaxiationAgentSecret string `json:"TaxiationAgentSecret" xml:"TaxiationAgentSecret"`
	TaxpayerType         string `json:"TaxpayerType" xml:"TaxpayerType"`
	TenantId             int64  `json:"TenantId" xml:"TenantId"`
}

// CreateUpdateCustomerInfoRequest creates a request to invoke UpdateCustomerInfo API
func CreateUpdateCustomerInfoRequest() (request *UpdateCustomerInfoRequest) {
	request = &UpdateCustomerInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "UpdateCustomerInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateCustomerInfoResponse creates a response to parse from UpdateCustomerInfo response
func CreateUpdateCustomerInfoResponse() (response *UpdateCustomerInfoResponse) {
	response = &UpdateCustomerInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
