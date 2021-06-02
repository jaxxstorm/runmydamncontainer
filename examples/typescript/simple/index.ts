import * as pulumi from "@pulumi/pulumi";
import * as rdc from "@jaxxstorm/pulumi-rdc";

const awsapp = new rdc.AWSInstance("aws-example", {
    image: "public.ecr.aws/aws-containers/hello-app-runner:latest",
    port: 8000,
})

const azureapp = new rdc.AzureInstance("azure-example", {
    image: "mcr.microsoft.com/azuredocs/aci-helloworld",
    port: 80,
})

const gcpApp = new rdc.GCPInstance("gcp-example", {
    image: "gcr.io/cloudrun/hello:latest",
    port: 8000,
})

const urls = []

urls.push(
    gcpApp.url,
    azureapp.url,
    awsapp.url
)

export { urls }
