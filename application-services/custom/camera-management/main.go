//
// Copyright (c) 2022 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"

	appsdk "github.com/edgexfoundry/app-functions-sdk-go/v2/pkg"
	"github.com/edgexfoundry/edgex-examples/application-services/custom/camera-management/appcamera"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

const (
	serviceKey = "app-camera-management"
)

func main() {
	service, ok := appsdk.NewAppService(serviceKey)
	if !ok {
		fmt.Printf("error: unable to create new app service %s!\n", serviceKey)
		os.Exit(-1)
	}

	app := appcamera.NewAppServiceWithTargetType(service, &dtos.SystemEvent{})

	// TODO: Replace below functions with built in and/or your custom functions for your use case
	//       or remove if using Pipeline By Topics below.
	//       See https://docs.edgexfoundry.org/latest/microservices/application/BuiltIn/ for list of built-in functions
	// err = app.service.SetDefaultFunctionsPipeline(
	// transforms.NewFilterFor(deviceNames).FilterByDeviceName,
	// sample.LogEventDetails,
	// sample.ConvertEventToXML,
	// sample.OutputXML)
	// if err != nil {
	// 	app.lc.Errorf("SetFunctionsPipeline returned error: %s", err.Error())
	// 	return -1
	// }

	// TODO: Remove adding functions pipelines by topic if default pipeline above is all your Use Case needs.
	//       Or remove default above if your use case needs multiple pipelines by topic.
	// Example of adding functions pipelines by topic.
	// These pipelines will only execute if the specified topic match the incoming topic.
	// Note: Device services publish to the 'edgex/events/device/<device-service-name><profile-name>/<device-name>/<source-name>' topic
	//       Core Data publishes to the 'edgex/events/core/<device-service-name>/<profile-name>/<device-name>/<source-name>' topic
	// Note: This example with default above causes Events from Random-Float-Device device to be processed twice
	//       resulting in the XML to be published back to the MessageBus twice.
	// See https://docs.edgexfoundry.org/latest/microservices/application/AdvancedTopics/#pipeline-per-topics for more details.
	// 	err = app.service.AddFunctionsPipelineForTopics("Floats", []string{"events/device/device-virtual/+/Random-Float-Device/#"},
	// 		sample.LogEventDetails,
	// 		sample.ConvertEventToXML,
	// 		sample.OutputXML)
	// 	if err != nil {
	// 		app.lc.Errorf("AddFunctionsPipelineForTopic returned error: %s", err.Error())
	// 		return -1
	// 	}
	// 	// Note: This example with default above causes Events from Int32 source to be processed twice
	// 	//       resulting in the XML to be published back to the MessageBus twice.
	// 	err = app.service.AddFunctionsPipelineForTopics("Int32s", []string{"events/device/device-virtual/+/+/Int32"},
	// 		sample.LogEventDetails,
	// 		sample.SendGetCommand,
	// 		sample.ConvertEventToXML,
	// 		sample.OutputXML)
	// 	if err != nil {
	// 		app.lc.Errorf("AddFunctionsPipelineForTopic returned error: %s", err.Error())
	// 		return -1
	// 	}

	// 	if err := app.service.MakeItRun(); err != nil {
	// 		app.lc.Errorf("MakeItRun returned error: %s", err.Error())
	// 		return -1
	// 	}

	// 	// TODO: Do any required cleanup here, if needed

	// 	return 0
	// }

	if err := app.Run(); err != nil {
		service.LoggingClient().Error(err.Error())
		os.Exit(-1)
	}

	os.Exit(0)
}
