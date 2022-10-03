// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkmanager_core_network", coreNetworkDataSource)
}

// coreNetworkDataSource returns the Terraform awscc_networkmanager_core_network data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkManager::CoreNetwork resource.
func coreNetworkDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"core_network_arn": {
			// Property: CoreNetworkArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN (Amazon resource name) of core network",
			//   "type": "string"
			// }
			Description: "The ARN (Amazon resource name) of core network",
			Type:        types.StringType,
			Computed:    true,
		},
		"core_network_id": {
			// Property: CoreNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The Id of core network",
			//   "type": "string"
			// }
			Description: "The Id of core network",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The creation time of core network",
			//   "type": "string"
			// }
			Description: "The creation time of core network",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of core network",
			//   "type": "string"
			// }
			Description: "The description of core network",
			Type:        types.StringType,
			Computed:    true,
		},
		"edges": {
			// Property: Edges
			// CloudFormation resource type schema:
			// {
			//   "description": "The edges within a core network.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Asn": {
			//         "description": "The ASN of a core network edge.",
			//         "type": "number"
			//       },
			//       "EdgeLocation": {
			//         "description": "The Region where a core network edge is located.",
			//         "type": "string"
			//       },
			//       "InsideCidrBlocks": {
			//         "insertionOrder": false,
			//         "items": {
			//           "description": "The inside IP addresses used for core network edges.",
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The edges within a core network.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"asn": {
						// Property: Asn
						Description: "The ASN of a core network edge.",
						Type:        types.Float64Type,
						Computed:    true,
					},
					"edge_location": {
						// Property: EdgeLocation
						Description: "The Region where a core network edge is located.",
						Type:        types.StringType,
						Computed:    true,
					},
					"inside_cidr_blocks": {
						// Property: InsideCidrBlocks
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the global network that your core network is a part of.",
			//   "type": "string"
			// }
			Description: "The ID of the global network that your core network is a part of.",
			Type:        types.StringType,
			Computed:    true,
		},
		"owner_account": {
			// Property: OwnerAccount
			// CloudFormation resource type schema:
			// {
			//   "description": "Owner of the core network",
			//   "type": "string"
			// }
			Description: "Owner of the core network",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy_document": {
			// Property: PolicyDocument
			// CloudFormation resource type schema:
			// {
			//   "description": "Live policy document for the core network, you must provide PolicyDocument in Json Format",
			//   "type": "object"
			// }
			Description: "Live policy document for the core network, you must provide PolicyDocument in Json Format",
			Type:        JSONStringType,
			Computed:    true,
		},
		"segments": {
			// Property: Segments
			// CloudFormation resource type schema:
			// {
			//   "description": "The segments within a core network.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "EdgeLocations": {
			//         "insertionOrder": false,
			//         "items": {
			//           "description": "The Regions where the edges are located.",
			//           "type": "string"
			//         },
			//         "type": "array"
			//       },
			//       "Name": {
			//         "description": "Name of segment",
			//         "type": "string"
			//       },
			//       "SharedSegments": {
			//         "insertionOrder": false,
			//         "items": {
			//           "description": "The shared segments of a core network.",
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The segments within a core network.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"edge_locations": {
						// Property: EdgeLocations
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"name": {
						// Property: Name
						Description: "Name of segment",
						Type:        types.StringType,
						Computed:    true,
					},
					"shared_segments": {
						// Property: SharedSegments
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of core network",
			//   "type": "string"
			// }
			Description: "The state of core network",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the global network.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The tags for the global network.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::NetworkManager::CoreNetwork",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::CoreNetwork").WithTerraformTypeName("awscc_networkmanager_core_network")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asn":                "Asn",
		"core_network_arn":   "CoreNetworkArn",
		"core_network_id":    "CoreNetworkId",
		"created_at":         "CreatedAt",
		"description":        "Description",
		"edge_location":      "EdgeLocation",
		"edge_locations":     "EdgeLocations",
		"edges":              "Edges",
		"global_network_id":  "GlobalNetworkId",
		"inside_cidr_blocks": "InsideCidrBlocks",
		"key":                "Key",
		"name":               "Name",
		"owner_account":      "OwnerAccount",
		"policy_document":    "PolicyDocument",
		"segments":           "Segments",
		"shared_segments":    "SharedSegments",
		"state":              "State",
		"tags":               "Tags",
		"value":              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
