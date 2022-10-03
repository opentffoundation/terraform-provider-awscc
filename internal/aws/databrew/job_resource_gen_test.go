// Code generated by generators/resource/main.go; DO NOT EDIT.

package databrew_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDataBrewJob_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataBrew::Job", "awscc_databrew_job", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
