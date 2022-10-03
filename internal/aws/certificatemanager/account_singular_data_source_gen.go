// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package certificatemanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_certificatemanager_account", accountDataSource)
}

// accountDataSource returns the Terraform awscc_certificatemanager_account data source.
// This Terraform data source corresponds to the CloudFormation AWS::CertificateManager::Account resource.
func accountDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_id": {
			// Property: AccountId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"expiry_events_configuration": {
			// Property: ExpiryEventsConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DaysBeforeExpiry": {
			//       "maximum": 45,
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"days_before_expiry": {
						// Property: DaysBeforeExpiry
						Type:     types.Int64Type,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CertificateManager::Account",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CertificateManager::Account").WithTerraformTypeName("awscc_certificatemanager_account")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                  "AccountId",
		"days_before_expiry":          "DaysBeforeExpiry",
		"expiry_events_configuration": "ExpiryEventsConfiguration",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
