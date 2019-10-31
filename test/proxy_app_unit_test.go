package test

import (
	"fmt"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of a unit test for the Terraform module in examples/proxy-app
func TestProxyAppUnit(t *testing.T) {
	t.Parallel()

	uniqueId := random.UniqueId()

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: proxyAppPath,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name": fmt.Sprintf("proxy-app-%s", uniqueId),

			// To make this a unit test, we proxy a known, "mock" endpoint
			"url_to_proxy": "https://www.example.com",
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	url := terraform.Output(t, terraformOptions, "url")

	// Verify the app returns a 200 OK with the text from example.com
	maxRetries := 10
	timeBetweenRetries := 3 * time.Second
	http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, maxRetries, timeBetweenRetries, func(statusCode int, body string) bool {
		return statusCode == 200 && strings.Contains(body, "<h1>Example Domain</h1>")
	})
}
