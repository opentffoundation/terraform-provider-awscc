// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotsitewise_gateway", gatewayDataSource)
}

// gatewayDataSource returns the Terraform awscc_iotsitewise_gateway data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTSiteWise::Gateway resource.
func gatewayDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"gateway_capability_summaries": {
			// Property: GatewayCapabilitySummaries
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of gateway capability summaries that each contain a namespace and status.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Contains a summary of a gateway capability configuration.",
			//     "properties": {
			//       "CapabilityConfiguration": {
			//         "description": "The JSON document that defines the gateway capability's configuration.",
			//         "type": "string"
			//       },
			//       "CapabilityNamespace": {
			//         "description": "The namespace of the capability configuration.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "CapabilityNamespace"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of gateway capability summaries that each contain a namespace and status.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"capability_configuration": {
						// Property: CapabilityConfiguration
						Description: "The JSON document that defines the gateway capability's configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"capability_namespace": {
						// Property: CapabilityNamespace
						Description: "The namespace of the capability configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"gateway_id": {
			// Property: GatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the gateway device.",
			//   "type": "string"
			// }
			Description: "The ID of the gateway device.",
			Type:        types.StringType,
			Computed:    true,
		},
		"gateway_name": {
			// Property: GatewayName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique, friendly name for the gateway.",
			//   "type": "string"
			// }
			Description: "A unique, friendly name for the gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"gateway_platform": {
			// Property: GatewayPlatform
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The gateway's platform. You can only specify one platform in a gateway.",
			//   "oneOf": [
			//     {
			//       "required": [
			//         "Greengrass"
			//       ]
			//     },
			//     {
			//       "required": [
			//         "GreengrassV2"
			//       ]
			//     }
			//   ],
			//   "properties": {
			//     "Greengrass": {
			//       "additionalProperties": false,
			//       "description": "A gateway that runs on AWS IoT Greengrass V1.",
			//       "properties": {
			//         "GroupArn": {
			//           "description": "The ARN of the Greengrass group.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "GroupArn"
			//       ],
			//       "type": "object"
			//     },
			//     "GreengrassV2": {
			//       "additionalProperties": false,
			//       "description": "A gateway that runs on AWS IoT Greengrass V2.",
			//       "properties": {
			//         "CoreDeviceThingName": {
			//           "description": "The name of the CoreDevice in GreenGrass V2.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "CoreDeviceThingName"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The gateway's platform. You can only specify one platform in a gateway.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"greengrass": {
						// Property: Greengrass
						Description: "A gateway that runs on AWS IoT Greengrass V1.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"group_arn": {
									// Property: GroupArn
									Description: "The ARN of the Greengrass group.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"greengrass_v2": {
						// Property: GreengrassV2
						Description: "A gateway that runs on AWS IoT Greengrass V2.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"core_device_thing_name": {
									// Property: CoreDeviceThingName
									Description: "The name of the CoreDevice in GreenGrass V2.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the gateway.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the gateway.",
			Attributes: tfsdk.ListNestedAttributes(
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::Gateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Gateway").WithTerraformTypeName("awscc_iotsitewise_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"capability_configuration":     "CapabilityConfiguration",
		"capability_namespace":         "CapabilityNamespace",
		"core_device_thing_name":       "CoreDeviceThingName",
		"gateway_capability_summaries": "GatewayCapabilitySummaries",
		"gateway_id":                   "GatewayId",
		"gateway_name":                 "GatewayName",
		"gateway_platform":             "GatewayPlatform",
		"greengrass":                   "Greengrass",
		"greengrass_v2":                "GreengrassV2",
		"group_arn":                    "GroupArn",
		"key":                          "Key",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
