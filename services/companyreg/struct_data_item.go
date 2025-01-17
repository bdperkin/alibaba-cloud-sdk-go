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

// DataItem is a nested struct in companyreg response
type DataItem struct {
	ProductName   string `json:"ProductName" xml:"ProductName"`
	Count         int    `json:"Count" xml:"Count"`
	IsSaveVoucher bool   `json:"IsSaveVoucher" xml:"IsSaveVoucher"`
	Date          string `json:"Date" xml:"Date"`
	Id            int64  `json:"Id" xml:"Id"`
	Amount        string `json:"Amount" xml:"Amount"`
	TaxAndAmount  string `json:"TaxAndAmount" xml:"TaxAndAmount"`
	Period        string `json:"Period" xml:"Period"`
	Url           string `json:"Url" xml:"Url"`
	Use           bool   `json:"Use" xml:"Use"`
	Tax           string `json:"Tax" xml:"Tax"`
	InvoiceNo     string `json:"InvoiceNo" xml:"InvoiceNo"`
	OrgName       string `json:"OrgName" xml:"OrgName"`
	Type          string `json:"Type" xml:"Type"`
}
