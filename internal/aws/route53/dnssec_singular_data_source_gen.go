// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53_dnssec", dNSSECDataSource)
}

// dNSSECDataSource returns the Terraform awscc_route53_dnssec data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53::DNSSEC resource.
func dNSSECDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"hosted_zone_id": {
			// Property: HostedZoneId
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique string (ID) used to identify a hosted zone.",
			//   "pattern": "^[A-Z0-9]{1,32}$",
			//   "type": "string"
			// }
			Description: "The unique string (ID) used to identify a hosted zone.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Route53::DNSSEC",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::DNSSEC").WithTerraformTypeName("awscc_route53_dnssec")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"hosted_zone_id": "HostedZoneId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
