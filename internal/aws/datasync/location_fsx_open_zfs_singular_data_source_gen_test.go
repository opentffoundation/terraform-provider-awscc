// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDataSyncLocationFSxOpenZFSDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataSync::LocationFSxOpenZFS", "awscc_datasync_location_fsx_open_zfs", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSDataSyncLocationFSxOpenZFSDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::DataSync::LocationFSxOpenZFS", "awscc_datasync_location_fsx_open_zfs", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
