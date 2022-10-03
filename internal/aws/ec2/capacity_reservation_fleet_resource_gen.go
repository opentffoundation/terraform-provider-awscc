// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_capacity_reservation_fleet", capacityReservationFleetResource)
}

// capacityReservationFleetResource returns the Terraform awscc_ec2_capacity_reservation_fleet resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::CapacityReservationFleet resource.
func capacityReservationFleetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"allocation_strategy": {
			// Property: AllocationStrategy
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"capacity_reservation_fleet_id": {
			// Property: CapacityReservationFleetId
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
		"end_date": {
			// Property: EndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"instance_match_criteria": {
			// Property: InstanceMatchCriteria
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "open"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"open",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"instance_type_specifications": {
			// Property: InstanceTypeSpecifications
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AvailabilityZone": {
			//         "type": "string"
			//       },
			//       "AvailabilityZoneId": {
			//         "type": "string"
			//       },
			//       "EbsOptimized": {
			//         "type": "boolean"
			//       },
			//       "InstancePlatform": {
			//         "type": "string"
			//       },
			//       "InstanceType": {
			//         "type": "string"
			//       },
			//       "Priority": {
			//         "maximum": 999,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "Weight": {
			//         "type": "number"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"availability_zone": {
						// Property: AvailabilityZone
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"availability_zone_id": {
						// Property: AvailabilityZoneId
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ebs_optimized": {
						// Property: EbsOptimized
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"instance_platform": {
						// Property: InstancePlatform
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"instance_type": {
						// Property: InstanceType
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"priority": {
						// Property: Priority
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 999),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"weight": {
						// Property: Weight
						Type:     types.Float64Type,
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
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"no_remove_end_date": {
			// Property: NoRemoveEndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"remove_end_date": {
			// Property: RemoveEndDate
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tag_specifications": {
			// Property: TagSpecifications
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ResourceType": {
			//         "type": "string"
			//       },
			//       "Tags": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Key": {
			//               "type": "string"
			//             },
			//             "Value": {
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Value",
			//             "Key"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": false
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"resource_type": {
						// Property: ResourceType
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"tags": {
						// Property: Tags
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
									Required: true,
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
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tenancy": {
			// Property: Tenancy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "default"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"default",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"total_target_capacity": {
			// Property: TotalTargetCapacity
			// CloudFormation resource type schema:
			// {
			//   "maximum": 25000,
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(1, 25000),
			},
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
		Description: "Resource Type definition for AWS::EC2::CapacityReservationFleet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CapacityReservationFleet").WithTerraformTypeName("awscc_ec2_capacity_reservation_fleet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_strategy":           "AllocationStrategy",
		"availability_zone":             "AvailabilityZone",
		"availability_zone_id":          "AvailabilityZoneId",
		"capacity_reservation_fleet_id": "CapacityReservationFleetId",
		"ebs_optimized":                 "EbsOptimized",
		"end_date":                      "EndDate",
		"instance_match_criteria":       "InstanceMatchCriteria",
		"instance_platform":             "InstancePlatform",
		"instance_type":                 "InstanceType",
		"instance_type_specifications":  "InstanceTypeSpecifications",
		"key":                           "Key",
		"no_remove_end_date":            "NoRemoveEndDate",
		"priority":                      "Priority",
		"remove_end_date":               "RemoveEndDate",
		"resource_type":                 "ResourceType",
		"tag_specifications":            "TagSpecifications",
		"tags":                          "Tags",
		"tenancy":                       "Tenancy",
		"total_target_capacity":         "TotalTargetCapacity",
		"value":                         "Value",
		"weight":                        "Weight",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
