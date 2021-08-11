# Run My Damn Container

Run My Damn Container (RDC) is the multi cloud solution you've been waiting for. Deploy your container to the 3 major cloud providers with just a few lines of code.

RDC is a [Pulumi](https://pulumi.com) package that allows you to ship a built Docker Image into the major cloud providers available container running engines. It's designed to have absolutely no configurability whatsoever, if you want to configure your application or image in any way at all, hardcode it into your Docker image - there's no room for best practices in a multi cloud world.

## Installing

### Installing the Plugin Binary

You'll need to ensure the Pulumi plugin is installed. Automatic plugin acquisition doesn't currently work for third party components, so you'll need to install manually:

```
pulumi plugin install resource rdc v0.0.1-alpha.1622659303+64064f31 --server https://lbriggs.jfrog.io/artifactory/pulumi-packages/runmydamncontainer
```

Update the version to refer to the [latest release]()

## Usage


