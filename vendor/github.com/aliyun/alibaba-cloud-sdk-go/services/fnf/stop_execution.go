package fnf

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

// StopExecution invokes the fnf.StopExecution API synchronously
// api document: https://help.aliyun.com/api/fnf/stopexecution.html
func (client *Client) StopExecution(request *StopExecutionRequest) (response *StopExecutionResponse, err error) {
	response = CreateStopExecutionResponse()
	err = client.DoAction(request, response)
	return
}

// StopExecutionWithChan invokes the fnf.StopExecution API asynchronously
// api document: https://help.aliyun.com/api/fnf/stopexecution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopExecutionWithChan(request *StopExecutionRequest) (<-chan *StopExecutionResponse, <-chan error) {
	responseChan := make(chan *StopExecutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopExecution(request)
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

// StopExecutionWithCallback invokes the fnf.StopExecution API asynchronously
// api document: https://help.aliyun.com/api/fnf/stopexecution.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopExecutionWithCallback(request *StopExecutionRequest, callback func(response *StopExecutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopExecutionResponse
		var err error
		defer close(result)
		response, err = client.StopExecution(request)
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

// StopExecutionRequest is the request struct for api StopExecution
type StopExecutionRequest struct {
	*requests.RpcRequest
	ExecutionName string `position:"Body" name:"ExecutionName"`
	RequestId     string `position:"Query" name:"RequestId"`
	Cause         string `position:"Body" name:"Cause"`
	FlowName      string `position:"Body" name:"FlowName"`
	Error         string `position:"Body" name:"Error"`
}

// StopExecutionResponse is the response struct for api StopExecution
type StopExecutionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Name           string `json:"Name" xml:"Name"`
	FlowName       string `json:"FlowName" xml:"FlowName"`
	FlowDefinition string `json:"FlowDefinition" xml:"FlowDefinition"`
	Input          string `json:"Input" xml:"Input"`
	Output         string `json:"Output" xml:"Output"`
	Status         string `json:"Status" xml:"Status"`
	StartedTime    string `json:"StartedTime" xml:"StartedTime"`
	StoppedTime    string `json:"StoppedTime" xml:"StoppedTime"`
}

// CreateStopExecutionRequest creates a request to invoke StopExecution API
func CreateStopExecutionRequest() (request *StopExecutionRequest) {
	request = &StopExecutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "StopExecution", "fnf", "openAPI")
	return
}

// CreateStopExecutionResponse creates a response to parse from StopExecution response
func CreateStopExecutionResponse() (response *StopExecutionResponse) {
	response = &StopExecutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}