package aws

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apprunner"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppRunnerArgs struct {
	Image pulumi.StringInput `pulumi:"image"`
	Port  pulumi.StringInput `pulumi:"port"`
}

type AppRunner struct {
	pulumi.ResourceState
	Url pulumi.StringOutput `pulumi:"url"`
}

func NewAppRunner(ctx *pulumi.Context, name string, args *AppRunnerArgs, opts ...pulumi.ResourceOption) (*AppRunner, error) {
	if args == nil {
		args = &AppRunnerArgs{}
	}

	component := &AppRunner{}

	var err error

	err = ctx.RegisterComponentResource("rdc:index:AppRunner", name, component, opts...)
	if err != nil {
		return nil, err
	}

	app, err := apprunner.NewService(ctx, name, &apprunner.ServiceArgs{
		ServiceName: pulumi.String(name),
		SourceConfiguration: &apprunner.ServiceSourceConfigurationArgs{
			ImageRepository: &apprunner.ServiceSourceConfigurationImageRepositoryArgs{
				ImageIdentifier:     args.Image,
				ImageRepositoryType: pulumi.String("ECR_PUBLIC"),
				ImageConfiguration: &apprunner.ServiceSourceConfigurationImageRepositoryImageConfigurationArgs{
					Port: args.Port,
				},
			},
			AutoDeploymentsEnabled: pulumi.Bool(false),
		},
	}, pulumi.Parent(component))
	if err != nil {
		return nil, fmt.Errorf("error creating apprunner service: %v", err)
	}

	component.Url = pulumi.Sprintf("https://%s", app.ServiceUrl)

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"url": app.ServiceUrl,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
