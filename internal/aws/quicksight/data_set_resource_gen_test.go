// Code generated by generators/resource/main.go; DO NOT EDIT.

package quicksight_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSQuickSightDataSet_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::QuickSight::DataSet", "awscc_quicksight_data_set", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSQuickSightDataSet_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::QuickSight::DataSet", "awscc_quicksight_data_set", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
