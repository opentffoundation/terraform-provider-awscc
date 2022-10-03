// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediapackage_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSMediaPackageAssetDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaPackage::Asset", "awscc_mediapackage_asset", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSMediaPackageAssetDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::MediaPackage::Asset", "awscc_mediapackage_asset", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
