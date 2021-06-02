package provider

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceArgs struct {
	Image         pulumi.StringInput `pulumi:"image"`
	Port          pulumi.IntInput    `pulumi:"port"`
	Location      pulumi.StringInput `pulumi:"location"`
}

// The Instance component resource.
type Instance struct {
	pulumi.ResourceState
	Url pulumi.StringOutput `pulumi:"url"`
}
