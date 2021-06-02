package azure

import (
	"fmt"

	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/containerinstance"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AciArgs struct {
	Image pulumi.StringInput `pulumi:"image"`
	Port  pulumi.IntInput    `pulumi:"port"`
}

type Aci struct {
	pulumi.ResourceState
	Url pulumi.StringOutput `pulumi:"url"`
}

func NewAci(ctx *pulumi.Context, name string, args *AciArgs, opts ...pulumi.ResourceOption) (*Aci, error) {

	if args == nil {
		args = &AciArgs{}
	}

	var err error
	component := &Aci{}

	err = ctx.RegisterComponentResource("rdc:index:Aci", name, component, opts...)
	if err != nil {
		return nil, err
	}

	resourceGroup, err := resources.NewResourceGroup(ctx, name, nil, pulumi.Parent(component))
	if err != nil {
		return nil, fmt.Errorf("error creating resource group: %v", err)
	}

	container, err := containerinstance.NewContainerGroup(ctx, name, &containerinstance.ContainerGroupArgs{
		ResourceGroupName: resourceGroup.Name,
		OsType:            pulumi.String("Linux"),
		Containers: &containerinstance.ContainerArray{
			&containerinstance.ContainerArgs{
				Name:  pulumi.String(name),
				Image: args.Image,
				Ports: &containerinstance.ContainerPortArray{
					&containerinstance.ContainerPortArgs{Port: args.Port},
				},
				Resources: &containerinstance.ResourceRequirementsArgs{
					Requests: &containerinstance.ResourceRequestsArgs{
						Cpu:        pulumi.Float64(1.0),
						MemoryInGB: pulumi.Float64(1.5),
					},
				},
			},
		},
		IpAddress: &containerinstance.IpAddressArgs{
			Ports: &containerinstance.PortArray{
				&containerinstance.PortArgs{
					Port:     args.Port,
					Protocol: pulumi.String("Tcp"),
				},
			},
			Type: pulumi.String("Public"),
		},
	}, pulumi.Parent(resourceGroup))

	if err != nil {
		return nil, fmt.Errorf("error creating container instance: %v", err)
	}

	url := container.IpAddress.ApplyT(func(ip *containerinstance.IpAddressResponse) (string, error) {
		if ip == nil || ip.Ip == nil {
			return "", nil
		}
		return *ip.Ip, nil
	}).(pulumi.StringOutput)

	component.Url = pulumi.Sprintf("http://%s:%d", url, args.Port)

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"url": url,
	}); err != nil {
		return nil, err
	}

	return component, nil

}
