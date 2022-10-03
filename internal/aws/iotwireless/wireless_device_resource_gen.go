// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_iotwireless_wireless_device", wirelessDeviceResource)
}

// wirelessDeviceResource returns the Terraform awscc_iotwireless_wireless_device resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTWireless::WirelessDevice resource.
func wirelessDeviceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "Wireless device arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device description",
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Description: "Wireless device description",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"destination_name": {
			// Property: DestinationName
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device destination name",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "Wireless device destination name",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(128),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device Id. Returned after successful create.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Wireless device Id. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"last_uplink_received_at": {
			// Property: LastUplinkReceivedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the most recent uplink was received.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date and time when the most recent uplink was received.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.",
			//   "oneOf": [
			//     {
			//       "required": [
			//         "OtaaV11"
			//       ]
			//     },
			//     {
			//       "required": [
			//         "OtaaV10x"
			//       ]
			//     },
			//     {
			//       "required": [
			//         "AbpV11"
			//       ]
			//     },
			//     {
			//       "required": [
			//         "AbpV10x"
			//       ]
			//     }
			//   ],
			//   "properties": {
			//     "AbpV10x": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "DevAddr": {
			//           "pattern": "[a-fA-F0-9]{8}",
			//           "type": "string"
			//         },
			//         "SessionKeys": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "AppSKey": {
			//               "pattern": "[a-fA-F0-9]{32}",
			//               "type": "string"
			//             },
			//             "NwkSKey": {
			//               "pattern": "[a-fA-F0-9]{32}",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "NwkSKey",
			//             "AppSKey"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "DevAddr",
			//         "SessionKeys"
			//       ],
			//       "type": "object"
			//     },
			//     "AbpV11": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "DevAddr": {
			//           "pattern": "[a-fA-F0-9]{8}",
			//           "type": "string"
			//         },
			//         "SessionKeys": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "AppSKey": {
			//               "pattern": "[a-fA-F0-9]{32}",
			//               "type": "string"
			//             },
			//             "FNwkSIntKey": {
			//               "pattern": "[a-fA-F0-9]{32}",
			//               "type": "string"
			//             },
			//             "NwkSEncKey": {
			//               "pattern": "[a-fA-F0-9]{32}",
			//               "type": "string"
			//             },
			//             "SNwkSIntKey": {
			//               "pattern": "[a-fA-F0-9]{32}",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "FNwkSIntKey",
			//             "SNwkSIntKey",
			//             "NwkSEncKey",
			//             "AppSKey"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "DevAddr",
			//         "SessionKeys"
			//       ],
			//       "type": "object"
			//     },
			//     "DevEui": {
			//       "pattern": "[a-f0-9]{16}",
			//       "type": "string"
			//     },
			//     "DeviceProfileId": {
			//       "maxLength": 256,
			//       "type": "string"
			//     },
			//     "OtaaV10x": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "AppEui": {
			//           "pattern": "[a-fA-F0-9]{16}",
			//           "type": "string"
			//         },
			//         "AppKey": {
			//           "pattern": "[a-fA-F0-9]{32}",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "AppKey",
			//         "AppEui"
			//       ],
			//       "type": "object"
			//     },
			//     "OtaaV11": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "AppKey": {
			//           "pattern": "[a-fA-F0-9]{32}",
			//           "type": "string"
			//         },
			//         "JoinEui": {
			//           "pattern": "[a-fA-F0-9]{16}",
			//           "type": "string"
			//         },
			//         "NwkKey": {
			//           "pattern": "[a-fA-F0-9]{32}",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "AppKey",
			//         "NwkKey",
			//         "JoinEui"
			//       ],
			//       "type": "object"
			//     },
			//     "ServiceProfileId": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"abp_v10_x": {
						// Property: AbpV10x
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"dev_addr": {
									// Property: DevAddr
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{8}"), ""),
									},
								},
								"session_keys": {
									// Property: SessionKeys
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"app_s_key": {
												// Property: AppSKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
												},
											},
											"nwk_s_key": {
												// Property: NwkSKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"abp_v11": {
						// Property: AbpV11
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"dev_addr": {
									// Property: DevAddr
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{8}"), ""),
									},
								},
								"session_keys": {
									// Property: SessionKeys
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"app_s_key": {
												// Property: AppSKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
												},
											},
											"f_nwk_s_int_key": {
												// Property: FNwkSIntKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
												},
											},
											"nwk_s_enc_key": {
												// Property: NwkSEncKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
												},
											},
											"s_nwk_s_int_key": {
												// Property: SNwkSIntKey
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"dev_eui": {
						// Property: DevEui
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("[a-f0-9]{16}"), ""),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"device_profile_id": {
						// Property: DeviceProfileId
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"otaa_v10_x": {
						// Property: OtaaV10x
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"app_eui": {
									// Property: AppEui
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{16}"), ""),
									},
								},
								"app_key": {
									// Property: AppKey
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"otaa_v11": {
						// Property: OtaaV11
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"app_key": {
									// Property: AppKey
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
									},
								},
								"join_eui": {
									// Property: JoinEui
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{16}"), ""),
									},
								},
								"nwk_key": {
									// Property: NwkKey
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("[a-fA-F0-9]{32}"), ""),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"service_profile_id": {
						// Property: ServiceProfileId
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.RequiredAttributes(
					validate.OneOfRequired(
						validate.Required(
							"otaa_v11",
						),
						validate.Required(
							"otaa_v10_x",
						),
						validate.Required(
							"abp_v11",
						),
						validate.Required(
							"abp_v10_x",
						),
					),
				),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device name",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Wireless device name",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.",
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
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"thing_arn": {
			// Property: ThingArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Thing arn. Passed into update to associate Thing with Wireless device.",
			//   "type": "string"
			// }
			Description: "Thing arn. Passed into update to associate Thing with Wireless device.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"thing_name": {
			// Property: ThingName
			// CloudFormation resource type schema:
			// {
			//   "description": "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			//   "type": "string"
			// }
			Description: "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device type, currently only Sidewalk and LoRa",
			//   "enum": [
			//     "Sidewalk",
			//     "LoRaWAN"
			//   ],
			//   "type": "string"
			// }
			Description: "Wireless device type, currently only Sidewalk and LoRa",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"Sidewalk",
					"LoRaWAN",
				}),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Create and manage wireless gateways, including LoRa gateways.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::WirelessDevice").WithTerraformTypeName("awscc_iotwireless_wireless_device")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"abp_v10_x":               "AbpV10x",
		"abp_v11":                 "AbpV11",
		"app_eui":                 "AppEui",
		"app_key":                 "AppKey",
		"app_s_key":               "AppSKey",
		"arn":                     "Arn",
		"description":             "Description",
		"destination_name":        "DestinationName",
		"dev_addr":                "DevAddr",
		"dev_eui":                 "DevEui",
		"device_profile_id":       "DeviceProfileId",
		"f_nwk_s_int_key":         "FNwkSIntKey",
		"id":                      "Id",
		"join_eui":                "JoinEui",
		"key":                     "Key",
		"last_uplink_received_at": "LastUplinkReceivedAt",
		"lo_ra_wan":               "LoRaWAN",
		"name":                    "Name",
		"nwk_key":                 "NwkKey",
		"nwk_s_enc_key":           "NwkSEncKey",
		"nwk_s_key":               "NwkSKey",
		"otaa_v10_x":              "OtaaV10x",
		"otaa_v11":                "OtaaV11",
		"s_nwk_s_int_key":         "SNwkSIntKey",
		"service_profile_id":      "ServiceProfileId",
		"session_keys":            "SessionKeys",
		"tags":                    "Tags",
		"thing_arn":               "ThingArn",
		"thing_name":              "ThingName",
		"type":                    "Type",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
