# Web Service Example

This folder contains example code that shows how to use the [serverless-app module](/modules/serverless-app) to deploy
a "Hello, World" Node.js web service on top of [AWS Lambda](https://aws.amazon.com/lambda/) and 
[API Gateway](https://aws.amazon.com/api-gateway/). 

This code is used in the talk 
[How to test your infrastructure code: automated testing for Terraform, Docker, Packer, Kubernetes, and more](https://qconsf.com/sf2019/presentation/infrastructure-0) 
by [Yevgeniy Brikman](https://www.ybrikman.com/) as a representation of typical infrastructure code that deploys an 
external dependency, such as a backend web service, that you may wish to mock when writing automated tests. 

**Note**: This repo is for demonstration and learning purposes only and should NOT be used to run anything important. 
For production-ready versions of this code and many other types of infrastructure, check out 
[Gruntwork](https://gruntwork.io/).

**WARNING**: This module and the automated tests for it deploy real resources into your AWS account which can cost you
money. The resources are all part of the [AWS Free Tier](https://aws.amazon.com/free/), so if you haven't used that up,
it should be free, but you are completely responsible for all AWS charges.

## Running this example manually

1. Sign up for [AWS](https://aws.amazon.com/).
1. Configure your AWS credentials using one of the [supported methods for AWS CLI
   tools](https://blog.gruntwork.io/a-comprehensive-guide-to-authenticating-to-aws-on-the-command-line-63656a686799), 
   such as setting the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. 
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Create an S3 bucket (e.g., manually) and configure it as the 
   [Backend](https://www.terraform.io/docs/backends/index.html) used to store this module's state by creating a 
   `backend.hcl` file with the contents:
    ```hcl
    bucket = "YOUR BUCKET'S NAME"
    region = "YOUR BUCKET'S REGION"
    key    = "web-service/terraform.tfstate"
    ``` 
1. Run `terraform init -backend-config=backend.hcl`.
1. Run `terraform apply`.
1. This module will output the URL of the web service at the end of `apply`. Try this URL out in your browser or
   via `curl` to see if it's working!
1. When you're done, run `terraform destroy`.

## Running automated tests against this example

1. Sign up for [AWS](https://aws.amazon.com/).
1. Configure your AWS credentials using one of the [supported methods for AWS CLI
   tools](https://blog.gruntwork.io/a-comprehensive-guide-to-authenticating-to-aws-on-the-command-line-63656a686799), 
   such as setting the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. 
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Install [Golang](https://golang.org/), minimum version `1.13`.
1. `cd test`
1. To run the unit test for this example: `go test -v -timeout 15m -run '^TestWebServiceUnit$'`
1. To run the integration test for this example: `go test -v -timeout 15m -run '^TestProxyAppIntegration$'`
1. To run the integration test with test stages for this example: `go test -v -timeout 15m -run '^TestProxyAppIntegrationWithStages$'`