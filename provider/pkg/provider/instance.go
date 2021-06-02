// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"strconv"

	"github.com/jaxxstorm/pulumi-rdc/pkg/provider/aws"
	"github.com/jaxxstorm/pulumi-rdc/pkg/provider/azure"
	"github.com/jaxxstorm/pulumi-rdc/pkg/provider/gcp"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewAWSInstance creates a new AWS AppRunner resource.
func NewAWSInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var err error

	component := &Instance{}

	err = ctx.RegisterComponentResource("rdc:index:AWSInstance", name, component, opts...)
	if err != nil {
		return nil, err
	}

	var url pulumi.StringOutput

	// convert the port to an integer
	port := args.Port.ToIntOutput().ApplyT(func(port int) string {
		i := strconv.Itoa(port)
		return i
	}).(pulumi.StringOutput)

	app, err := aws.NewAppRunner(ctx, name, &aws.AppRunnerArgs{
		Image: args.Image,
		Port:  port,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}
	url = app.Url

	component.Url = url

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"url": url,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

func NewAzureInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var err error

	component := &Instance{}

	err = ctx.RegisterComponentResource("rdc:index:AzureInstance", name, component, opts...)
	if err != nil {
		return nil, err
	}

	app, err := azure.NewAci(ctx, name, &azure.AciArgs{
		Image: args.Image,
		Port:  args.Port,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}
	url := app.Url

	component.Url = url

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"url": url,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

func NewGCPInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var err error

	component := &Instance{}

	err = ctx.RegisterComponentResource("rdc:index:GCPInstance", name, component, opts...)
	if err != nil {
		return nil, err
	}

	app, err := gcp.NewCloudRun(ctx, name, &gcp.CloudRunArgs{
		Image:    args.Image,
		Port:     args.Port,
		Location: args.Location,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}
	url := app.Url

	component.Url = url

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"url": url,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
