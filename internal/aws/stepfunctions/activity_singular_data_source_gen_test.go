// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package stepfunctions_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSStepFunctionsActivityDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::StepFunctions::Activity", "awscc_stepfunctions_activity", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSStepFunctionsActivityDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::StepFunctions::Activity", "awscc_stepfunctions_activity", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
