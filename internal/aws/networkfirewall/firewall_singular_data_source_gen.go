// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkfirewall_firewall", firewallDataSource)
}

// firewallDataSource returns the Terraform awscc_networkfirewall_firewall data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkFirewall::Firewall resource.
func firewallDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"delete_protection": {
			// Property: DeleteProtection
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "pattern": "^.*$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"endpoint_ids": {
			// Property: EndpointIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "An endpoint Id.",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"firewall_arn": {
			// Property: FirewallArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws.*$",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"firewall_id": {
			// Property: FirewallId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "^([0-9a-f]{8})-([0-9a-f]{4}-){3}([0-9a-f]{12})$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"firewall_name": {
			// Property: FirewallName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"firewall_policy_arn": {
			// Property: FirewallPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws.*$",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"firewall_policy_change_protection": {
			// Property: FirewallPolicyChangeProtection
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"subnet_change_protection": {
			// Property: SubnetChangeProtection
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"subnet_mappings": {
			// Property: SubnetMappings
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "SubnetId": {
			//         "description": "A SubnetId.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "SubnetId"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"subnet_id": {
						// Property: SubnetId
						Description: "A SubnetId.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "^vpc-[0-9a-f]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::NetworkFirewall::Firewall",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::Firewall").WithTerraformTypeName("awscc_networkfirewall_firewall")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"delete_protection":                 "DeleteProtection",
		"description":                       "Description",
		"endpoint_ids":                      "EndpointIds",
		"firewall_arn":                      "FirewallArn",
		"firewall_id":                       "FirewallId",
		"firewall_name":                     "FirewallName",
		"firewall_policy_arn":               "FirewallPolicyArn",
		"firewall_policy_change_protection": "FirewallPolicyChangeProtection",
		"key":                               "Key",
		"subnet_change_protection":          "SubnetChangeProtection",
		"subnet_id":                         "SubnetId",
		"subnet_mappings":                   "SubnetMappings",
		"tags":                              "Tags",
		"value":                             "Value",
		"vpc_id":                            "VpcId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
