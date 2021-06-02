"""A Python Pulumi program"""

import pulumi
import jaxxstorm_pulumi_rdc as rdc

aws = rdc.aws_instance(
    "aws-example",
    image="public.ecr.aws/aws-containers/hello-app-runner:latest",
    port=8000
)

azure = rdc.azure_instance(
    "azure-example",
    image="mcr.microsoft.com/azuredocs/aci-helloworld",
    port=8000,
)

gcp = rdc.gcp_instance(
    "gcp-example",
    image="gcr.io/cloudrun/hello:latest",
    port=80
)