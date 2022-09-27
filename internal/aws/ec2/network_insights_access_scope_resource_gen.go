// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_network_insights_access_scope", networkInsightsAccessScopeResourceType)
}

// networkInsightsAccessScopeResourceType returns the Terraform awscc_ec2_network_insights_access_scope resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::NetworkInsightsAccessScope resource type.
func networkInsightsAccessScopeResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_date": {
			// Property: CreatedDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"exclude_paths": {
			// Property: ExcludePaths
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Destination": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Source": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ThroughResources": {
			//         "insertionOrder": true,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ResourceStatement": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "ResourceTypes": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 },
			//                 "Resources": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayForEach(validate.StringInSlice([]string{
														"tcp",
														"udp",
													})),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"source": {
						// Property: Source
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayForEach(validate.StringInSlice([]string{
														"tcp",
														"udp",
													})),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"through_resources": {
						// Property: ThroughResources
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// ExcludePaths is a write-only property.
		},
		"match_paths": {
			// Property: MatchPaths
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Destination": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Source": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "PacketHeaderStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DestinationAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "DestinationPrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Protocols": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "enum": [
			//                     "tcp",
			//                     "udp"
			//                   ],
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourceAddresses": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePorts": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "SourcePrefixLists": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "ResourceStatement": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ResourceTypes": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               },
			//               "Resources": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "ThroughResources": {
			//         "insertionOrder": true,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ResourceStatement": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "ResourceTypes": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 },
			//                 "Resources": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "type": "string"
			//                   },
			//                   "type": "array"
			//                 }
			//               },
			//               "type": "object"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayForEach(validate.StringInSlice([]string{
														"tcp",
														"udp",
													})),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"source": {
						// Property: Source
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"packet_header_statement": {
									// Property: PacketHeaderStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"destination_addresses": {
												// Property: DestinationAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_ports": {
												// Property: DestinationPorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"destination_prefix_lists": {
												// Property: DestinationPrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"protocols": {
												// Property: Protocols
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayForEach(validate.StringInSlice([]string{
														"tcp",
														"udp",
													})),
												},
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_addresses": {
												// Property: SourceAddresses
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_ports": {
												// Property: SourcePorts
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"source_prefix_lists": {
												// Property: SourcePrefixLists
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"through_resources": {
						// Property: ThroughResources
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"resource_statement": {
									// Property: ResourceStatement
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"resource_types": {
												// Property: ResourceTypes
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
												},
											},
											"resources": {
												// Property: Resources
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Computed: true,
												PlanModifiers: []tfsdk.AttributePlanModifier{
													resource.UseStateForUnknown(),
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
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// MatchPaths is a write-only property.
		},
		"network_insights_access_scope_arn": {
			// Property: NetworkInsightsAccessScopeArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"network_insights_access_scope_id": {
			// Property: NetworkInsightsAccessScopeId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"updated_date": {
			// Property: UpdatedDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::EC2::NetworkInsightsAccessScope",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInsightsAccessScope").WithTerraformTypeName("awscc_ec2_network_insights_access_scope")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_date":                      "CreatedDate",
		"destination":                       "Destination",
		"destination_addresses":             "DestinationAddresses",
		"destination_ports":                 "DestinationPorts",
		"destination_prefix_lists":          "DestinationPrefixLists",
		"exclude_paths":                     "ExcludePaths",
		"key":                               "Key",
		"match_paths":                       "MatchPaths",
		"network_insights_access_scope_arn": "NetworkInsightsAccessScopeArn",
		"network_insights_access_scope_id":  "NetworkInsightsAccessScopeId",
		"packet_header_statement":           "PacketHeaderStatement",
		"protocols":                         "Protocols",
		"resource_statement":                "ResourceStatement",
		"resource_types":                    "ResourceTypes",
		"resources":                         "Resources",
		"source":                            "Source",
		"source_addresses":                  "SourceAddresses",
		"source_ports":                      "SourcePorts",
		"source_prefix_lists":               "SourcePrefixLists",
		"tags":                              "Tags",
		"through_resources":                 "ThroughResources",
		"updated_date":                      "UpdatedDate",
		"value":                             "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/MatchPaths",
		"/properties/ExcludePaths",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
