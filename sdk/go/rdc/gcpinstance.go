// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rdc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GCPInstance struct {
	pulumi.ResourceState

	// url of your running container
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewGCPInstance registers a new resource with the given unique name, arguments, and options.
func NewGCPInstance(ctx *pulumi.Context,
	name string, args *GCPInstanceArgs, opts ...pulumi.ResourceOption) (*GCPInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Location == nil {
		args.Location = pulumi.StringPtr("us-central1")
	}
	var resource GCPInstance
	err := ctx.RegisterRemoteComponentResource("rdc:index:GCPInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type gcpinstanceArgs struct {
	// The image you want to run
	Image string `pulumi:"image"`
	// The location to run your cloudrun container
	Location *string `pulumi:"location"`
	Port     int     `pulumi:"port"`
}

// The set of arguments for constructing a GCPInstance resource.
type GCPInstanceArgs struct {
	// The image you want to run
	Image pulumi.StringInput
	// The location to run your cloudrun container
	Location pulumi.StringPtrInput
	Port     pulumi.IntInput
}

func (GCPInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpinstanceArgs)(nil)).Elem()
}

type GCPInstanceInput interface {
	pulumi.Input

	ToGCPInstanceOutput() GCPInstanceOutput
	ToGCPInstanceOutputWithContext(ctx context.Context) GCPInstanceOutput
}

func (*GCPInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPInstance)(nil))
}

func (i *GCPInstance) ToGCPInstanceOutput() GCPInstanceOutput {
	return i.ToGCPInstanceOutputWithContext(context.Background())
}

func (i *GCPInstance) ToGCPInstanceOutputWithContext(ctx context.Context) GCPInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPInstanceOutput)
}

func (i *GCPInstance) ToGCPInstancePtrOutput() GCPInstancePtrOutput {
	return i.ToGCPInstancePtrOutputWithContext(context.Background())
}

func (i *GCPInstance) ToGCPInstancePtrOutputWithContext(ctx context.Context) GCPInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPInstancePtrOutput)
}

type GCPInstancePtrInput interface {
	pulumi.Input

	ToGCPInstancePtrOutput() GCPInstancePtrOutput
	ToGCPInstancePtrOutputWithContext(ctx context.Context) GCPInstancePtrOutput
}

type gcpinstancePtrType GCPInstanceArgs

func (*gcpinstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GCPInstance)(nil))
}

func (i *gcpinstancePtrType) ToGCPInstancePtrOutput() GCPInstancePtrOutput {
	return i.ToGCPInstancePtrOutputWithContext(context.Background())
}

func (i *gcpinstancePtrType) ToGCPInstancePtrOutputWithContext(ctx context.Context) GCPInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPInstancePtrOutput)
}

// GCPInstanceArrayInput is an input type that accepts GCPInstanceArray and GCPInstanceArrayOutput values.
// You can construct a concrete instance of `GCPInstanceArrayInput` via:
//
//          GCPInstanceArray{ GCPInstanceArgs{...} }
type GCPInstanceArrayInput interface {
	pulumi.Input

	ToGCPInstanceArrayOutput() GCPInstanceArrayOutput
	ToGCPInstanceArrayOutputWithContext(context.Context) GCPInstanceArrayOutput
}

type GCPInstanceArray []GCPInstanceInput

func (GCPInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GCPInstance)(nil))
}

func (i GCPInstanceArray) ToGCPInstanceArrayOutput() GCPInstanceArrayOutput {
	return i.ToGCPInstanceArrayOutputWithContext(context.Background())
}

func (i GCPInstanceArray) ToGCPInstanceArrayOutputWithContext(ctx context.Context) GCPInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPInstanceArrayOutput)
}

// GCPInstanceMapInput is an input type that accepts GCPInstanceMap and GCPInstanceMapOutput values.
// You can construct a concrete instance of `GCPInstanceMapInput` via:
//
//          GCPInstanceMap{ "key": GCPInstanceArgs{...} }
type GCPInstanceMapInput interface {
	pulumi.Input

	ToGCPInstanceMapOutput() GCPInstanceMapOutput
	ToGCPInstanceMapOutputWithContext(context.Context) GCPInstanceMapOutput
}

type GCPInstanceMap map[string]GCPInstanceInput

func (GCPInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GCPInstance)(nil))
}

func (i GCPInstanceMap) ToGCPInstanceMapOutput() GCPInstanceMapOutput {
	return i.ToGCPInstanceMapOutputWithContext(context.Background())
}

func (i GCPInstanceMap) ToGCPInstanceMapOutputWithContext(ctx context.Context) GCPInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPInstanceMapOutput)
}

type GCPInstanceOutput struct {
	*pulumi.OutputState
}

func (GCPInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GCPInstance)(nil))
}

func (o GCPInstanceOutput) ToGCPInstanceOutput() GCPInstanceOutput {
	return o
}

func (o GCPInstanceOutput) ToGCPInstanceOutputWithContext(ctx context.Context) GCPInstanceOutput {
	return o
}

func (o GCPInstanceOutput) ToGCPInstancePtrOutput() GCPInstancePtrOutput {
	return o.ToGCPInstancePtrOutputWithContext(context.Background())
}

func (o GCPInstanceOutput) ToGCPInstancePtrOutputWithContext(ctx context.Context) GCPInstancePtrOutput {
	return o.ApplyT(func(v GCPInstance) *GCPInstance {
		return &v
	}).(GCPInstancePtrOutput)
}

type GCPInstancePtrOutput struct {
	*pulumi.OutputState
}

func (GCPInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GCPInstance)(nil))
}

func (o GCPInstancePtrOutput) ToGCPInstancePtrOutput() GCPInstancePtrOutput {
	return o
}

func (o GCPInstancePtrOutput) ToGCPInstancePtrOutputWithContext(ctx context.Context) GCPInstancePtrOutput {
	return o
}

type GCPInstanceArrayOutput struct{ *pulumi.OutputState }

func (GCPInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GCPInstance)(nil))
}

func (o GCPInstanceArrayOutput) ToGCPInstanceArrayOutput() GCPInstanceArrayOutput {
	return o
}

func (o GCPInstanceArrayOutput) ToGCPInstanceArrayOutputWithContext(ctx context.Context) GCPInstanceArrayOutput {
	return o
}

func (o GCPInstanceArrayOutput) Index(i pulumi.IntInput) GCPInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GCPInstance {
		return vs[0].([]GCPInstance)[vs[1].(int)]
	}).(GCPInstanceOutput)
}

type GCPInstanceMapOutput struct{ *pulumi.OutputState }

func (GCPInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GCPInstance)(nil))
}

func (o GCPInstanceMapOutput) ToGCPInstanceMapOutput() GCPInstanceMapOutput {
	return o
}

func (o GCPInstanceMapOutput) ToGCPInstanceMapOutputWithContext(ctx context.Context) GCPInstanceMapOutput {
	return o
}

func (o GCPInstanceMapOutput) MapIndex(k pulumi.StringInput) GCPInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GCPInstance {
		return vs[0].(map[string]GCPInstance)[vs[1].(string)]
	}).(GCPInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(GCPInstanceOutput{})
	pulumi.RegisterOutputType(GCPInstancePtrOutput{})
	pulumi.RegisterOutputType(GCPInstanceArrayOutput{})
	pulumi.RegisterOutputType(GCPInstanceMapOutput{})
}
