package gcp

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudRunArgs struct {
	Image    pulumi.StringInput `pulumi:"image"`
	Port     pulumi.IntInput    `pulumi:"port"`
	Location pulumi.StringInput `pulumi:"location"`
}

type CloudRun struct {
	pulumi.ResourceState
	Url pulumi.StringOutput `pulumi:"url"`
}

func NewCloudRun(ctx *pulumi.Context, name string, args *CloudRunArgs, opts ...pulumi.ResourceOption) (*CloudRun, error) {

	if args == nil {
		args = &CloudRunArgs{}
	}

	var err error
	component := &CloudRun{}

	err = ctx.RegisterComponentResource("rdc:index:CloudRun", name, component, opts...)
	if err != nil {
		return nil, err
	}

	svc, err := cloudrun.NewService(ctx, name, &cloudrun.ServiceArgs{
		Location: args.Location,
		Template: &cloudrun.ServiceTemplateArgs{
			Spec: &cloudrun.ServiceTemplateSpecArgs{
				Containers: &cloudrun.ServiceTemplateSpecContainerArray{
					&cloudrun.ServiceTemplateSpecContainerArgs{
						Image: args.Image,
						Ports: &cloudrun.ServiceTemplateSpecContainerPortArray{
							&cloudrun.ServiceTemplateSpecContainerPortArgs{
								ContainerPort: args.Port,
							},
						},
					},
				},
			},
		},
	}, pulumi.Parent(component))
	if err != nil {
		return nil, fmt.Errorf("error creating cloudrun service: %v", err)
	}

	_, err = cloudrun.NewIamMember(ctx, name, &cloudrun.IamMemberArgs{
		Service:  svc.Name,
		Location: svc.Location,
		Role:     pulumi.String("roles/run.invoker"),
		Member:   pulumi.String("allUsers"),
	}, pulumi.Parent(svc))
	if err != nil {
		return nil, fmt.Errorf("error creating cloudrun iam member: %v", err)
	}

	url := pulumi.StringOutput(svc.Statuses.Index(pulumi.Int(0)).Url())

	component.Url = url

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"url": url,
	}); err != nil {
		return nil, err
	}

	return component, nil

}
