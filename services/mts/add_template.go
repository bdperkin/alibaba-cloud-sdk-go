package mts

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

func (client *Client) AddTemplate(request *AddTemplateRequest) (response *AddTemplateResponse, err error) {
	response = CreateAddTemplateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddTemplateWithChan(request *AddTemplateRequest) (<-chan *AddTemplateResponse, <-chan error) {
	responseChan := make(chan *AddTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTemplate(request)
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

func (client *Client) AddTemplateWithCallback(request *AddTemplateRequest, callback func(response *AddTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddTemplate(request)
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

type AddTemplateRequest struct {
	*requests.RpcRequest
	Container            string           `position:"Query" name:"Container"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	TransConfig          string           `position:"Query" name:"TransConfig"`
	MuxConfig            string           `position:"Query" name:"MuxConfig"`
	Video                string           `position:"Query" name:"Video"`
	Audio                string           `position:"Query" name:"Audio"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type AddTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Template  struct {
		Id        string `json:"Id" xml:"Id"`
		Name      string `json:"Name" xml:"Name"`
		State     string `json:"State" xml:"State"`
		Container struct {
			Format string `json:"Format" xml:"Format"`
		} `json:"Container" xml:"Container"`
		Video struct {
			Codec      string `json:"Codec" xml:"Codec"`
			Profile    string `json:"Profile" xml:"Profile"`
			Bitrate    string `json:"Bitrate" xml:"Bitrate"`
			Crf        string `json:"Crf" xml:"Crf"`
			Width      string `json:"Width" xml:"Width"`
			Height     string `json:"Height" xml:"Height"`
			Fps        string `json:"Fps" xml:"Fps"`
			Gop        string `json:"Gop" xml:"Gop"`
			Preset     string `json:"Preset" xml:"Preset"`
			ScanMode   string `json:"ScanMode" xml:"ScanMode"`
			Bufsize    string `json:"Bufsize" xml:"Bufsize"`
			Maxrate    string `json:"Maxrate" xml:"Maxrate"`
			PixFmt     string `json:"PixFmt" xml:"PixFmt"`
			Degrain    string `json:"Degrain" xml:"Degrain"`
			Qscale     string `json:"Qscale" xml:"Qscale"`
			Remove     string `json:"Remove" xml:"Remove"`
			Crop       string `json:"Crop" xml:"Crop"`
			Pad        string `json:"Pad" xml:"Pad"`
			MaxFps     string `json:"MaxFps" xml:"MaxFps"`
			BitrateBnd struct {
				Max string `json:"Max" xml:"Max"`
				Min string `json:"Min" xml:"Min"`
			} `json:"BitrateBnd" xml:"BitrateBnd"`
		} `json:"Video" xml:"Video"`
		Audio struct {
			Codec      string `json:"Codec" xml:"Codec"`
			Profile    string `json:"Profile" xml:"Profile"`
			Samplerate string `json:"Samplerate" xml:"Samplerate"`
			Bitrate    string `json:"Bitrate" xml:"Bitrate"`
			Channels   string `json:"Channels" xml:"Channels"`
			Qscale     string `json:"Qscale" xml:"Qscale"`
			Remove     string `json:"Remove" xml:"Remove"`
			Volume     struct {
				Level  string `json:"Level" xml:"Level"`
				Method string `json:"Method" xml:"Method"`
			} `json:"Volume" xml:"Volume"`
		} `json:"Audio" xml:"Audio"`
		TransConfig struct {
			TransMode               string `json:"TransMode" xml:"TransMode"`
			IsCheckReso             string `json:"IsCheckReso" xml:"IsCheckReso"`
			IsCheckResoFail         string `json:"IsCheckResoFail" xml:"IsCheckResoFail"`
			IsCheckVideoBitrate     string `json:"IsCheckVideoBitrate" xml:"IsCheckVideoBitrate"`
			IsCheckAudioBitrate     string `json:"IsCheckAudioBitrate" xml:"IsCheckAudioBitrate"`
			AdjDarMethod            string `json:"AdjDarMethod" xml:"AdjDarMethod"`
			IsCheckVideoBitrateFail string `json:"IsCheckVideoBitrateFail" xml:"IsCheckVideoBitrateFail"`
			IsCheckAudioBitrateFail string `json:"IsCheckAudioBitrateFail" xml:"IsCheckAudioBitrateFail"`
		} `json:"TransConfig" xml:"TransConfig"`
		MuxConfig struct {
			Segment struct {
				Duration string `json:"Duration" xml:"Duration"`
			} `json:"Segment" xml:"Segment"`
			Gif struct {
				Loop            string `json:"Loop" xml:"Loop"`
				FinalDelay      string `json:"FinalDelay" xml:"FinalDelay"`
				IsCustomPalette string `json:"IsCustomPalette" xml:"IsCustomPalette"`
				DitherMode      string `json:"DitherMode" xml:"DitherMode"`
			} `json:"Gif" xml:"Gif"`
		} `json:"MuxConfig" xml:"MuxConfig"`
	} `json:"Template" xml:"Template"`
}

func CreateAddTemplateRequest() (request *AddTemplateRequest) {
	request = &AddTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddTemplate", "mts", "openAPI")
	return
}

func CreateAddTemplateResponse() (response *AddTemplateResponse) {
	response = &AddTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
