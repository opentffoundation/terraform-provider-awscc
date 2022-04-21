// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSConnectPhoneNumberDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Connect::PhoneNumber", "awscc_connect_phone_number", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSConnectPhoneNumberDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Connect::PhoneNumber", "awscc_connect_phone_number", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
