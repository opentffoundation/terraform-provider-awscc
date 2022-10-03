// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

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
	registry.AddResourceFactory("awscc_ecs_service", serviceResource)
}

// serviceResource returns the Terraform awscc_ecs_service resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECS::Service resource.
func serviceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"capacity_provider_strategy": {
			// Property: CapacityProviderStrategy
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Base": {
			//         "type": "integer"
			//       },
			//       "CapacityProvider": {
			//         "type": "string"
			//       },
			//       "Weight": {
			//         "type": "integer"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"base": {
						// Property: Base
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"capacity_provider": {
						// Property: CapacityProvider
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"weight": {
						// Property: Weight
						Type:     types.Int64Type,
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
		"cluster": {
			// Property: Cluster
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
		"deployment_configuration": {
			// Property: DeploymentConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DeploymentCircuitBreaker": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Enable": {
			//           "type": "boolean"
			//         },
			//         "Rollback": {
			//           "type": "boolean"
			//         }
			//       },
			//       "required": [
			//         "Enable",
			//         "Rollback"
			//       ],
			//       "type": "object"
			//     },
			//     "MaximumPercent": {
			//       "type": "integer"
			//     },
			//     "MinimumHealthyPercent": {
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"deployment_circuit_breaker": {
						// Property: DeploymentCircuitBreaker
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enable": {
									// Property: Enable
									Type:     types.BoolType,
									Required: true,
								},
								"rollback": {
									// Property: Rollback
									Type:     types.BoolType,
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
					"maximum_percent": {
						// Property: MaximumPercent
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"minimum_healthy_percent": {
						// Property: MinimumHealthyPercent
						Type:     types.Int64Type,
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
		"deployment_controller": {
			// Property: DeploymentController
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Type": {
			//       "enum": [
			//         "CODE_DEPLOY",
			//         "ECS",
			//         "EXTERNAL"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"CODE_DEPLOY",
								"ECS",
								"EXTERNAL",
							}),
						},
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
		},
		"desired_count": {
			// Property: DesiredCount
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enable_ecs_managed_tags": {
			// Property: EnableECSManagedTags
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
		"enable_execute_command": {
			// Property: EnableExecuteCommand
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
		"health_check_grace_period_seconds": {
			// Property: HealthCheckGracePeriodSeconds
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"launch_type": {
			// Property: LaunchType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "EC2",
			//     "FARGATE",
			//     "EXTERNAL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"EC2",
					"FARGATE",
					"EXTERNAL",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"load_balancers": {
			// Property: LoadBalancers
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ContainerName": {
			//         "type": "string"
			//       },
			//       "ContainerPort": {
			//         "type": "integer"
			//       },
			//       "LoadBalancerName": {
			//         "type": "string"
			//       },
			//       "TargetGroupArn": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"container_name": {
						// Property: ContainerName
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"container_port": {
						// Property: ContainerPort
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"load_balancer_name": {
						// Property: LoadBalancerName
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"target_group_arn": {
						// Property: TargetGroupArn
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
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
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
		"network_configuration": {
			// Property: NetworkConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AwsvpcConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "AssignPublicIp": {
			//           "enum": [
			//             "DISABLED",
			//             "ENABLED"
			//           ],
			//           "type": "string"
			//         },
			//         "SecurityGroups": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "Subnets": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"awsvpc_configuration": {
						// Property: AwsvpcConfiguration
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"assign_public_ip": {
									// Property: AssignPublicIp
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"DISABLED",
											"ENABLED",
										}),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"security_groups": {
									// Property: SecurityGroups
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"subnets": {
									// Property: Subnets
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
		"placement_constraints": {
			// Property: PlacementConstraints
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Expression": {
			//         "type": "string"
			//       },
			//       "Type": {
			//         "enum": [
			//           "distinctInstance",
			//           "memberOf"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"expression": {
						// Property: Expression
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"distinctInstance",
								"memberOf",
							}),
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
		"placement_strategies": {
			// Property: PlacementStrategies
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Field": {
			//         "type": "string"
			//       },
			//       "Type": {
			//         "enum": [
			//           "binpack",
			//           "random",
			//           "spread"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"field": {
						// Property: Field
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"binpack",
								"random",
								"spread",
							}),
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
		"platform_version": {
			// Property: PlatformVersion
			// CloudFormation resource type schema:
			// {
			//   "default": "LATEST",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.String{Value: "LATEST"}),
				resource.UseStateForUnknown(),
			},
		},
		"propagate_tags": {
			// Property: PropagateTags
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "SERVICE",
			//     "TASK_DEFINITION"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SERVICE",
					"TASK_DEFINITION",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"role": {
			// Property: Role
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
		"scheduling_strategy": {
			// Property: SchedulingStrategy
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DAEMON",
			//     "REPLICA"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DAEMON",
					"REPLICA",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"service_arn": {
			// Property: ServiceArn
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
		"service_name": {
			// Property: ServiceName
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
		"service_registries": {
			// Property: ServiceRegistries
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ContainerName": {
			//         "type": "string"
			//       },
			//       "ContainerPort": {
			//         "type": "integer"
			//       },
			//       "Port": {
			//         "type": "integer"
			//       },
			//       "RegistryArn": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"container_name": {
						// Property: ContainerName
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"container_port": {
						// Property: ContainerPort
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"port": {
						// Property: Port
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"registry_arn": {
						// Property: RegistryArn
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
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
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
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
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
				resource.UseStateForUnknown(),
			},
		},
		"task_definition": {
			// Property: TaskDefinition
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
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
		Description: "Resource Type definition for AWS::ECS::Service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::Service").WithTerraformTypeName("awscc_ecs_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_public_ip":                  "AssignPublicIp",
		"awsvpc_configuration":              "AwsvpcConfiguration",
		"base":                              "Base",
		"capacity_provider":                 "CapacityProvider",
		"capacity_provider_strategy":        "CapacityProviderStrategy",
		"cluster":                           "Cluster",
		"container_name":                    "ContainerName",
		"container_port":                    "ContainerPort",
		"deployment_circuit_breaker":        "DeploymentCircuitBreaker",
		"deployment_configuration":          "DeploymentConfiguration",
		"deployment_controller":             "DeploymentController",
		"desired_count":                     "DesiredCount",
		"enable":                            "Enable",
		"enable_ecs_managed_tags":           "EnableECSManagedTags",
		"enable_execute_command":            "EnableExecuteCommand",
		"expression":                        "Expression",
		"field":                             "Field",
		"health_check_grace_period_seconds": "HealthCheckGracePeriodSeconds",
		"key":                               "Key",
		"launch_type":                       "LaunchType",
		"load_balancer_name":                "LoadBalancerName",
		"load_balancers":                    "LoadBalancers",
		"maximum_percent":                   "MaximumPercent",
		"minimum_healthy_percent":           "MinimumHealthyPercent",
		"name":                              "Name",
		"network_configuration":             "NetworkConfiguration",
		"placement_constraints":             "PlacementConstraints",
		"placement_strategies":              "PlacementStrategies",
		"platform_version":                  "PlatformVersion",
		"port":                              "Port",
		"propagate_tags":                    "PropagateTags",
		"registry_arn":                      "RegistryArn",
		"role":                              "Role",
		"rollback":                          "Rollback",
		"scheduling_strategy":               "SchedulingStrategy",
		"security_groups":                   "SecurityGroups",
		"service_arn":                       "ServiceArn",
		"service_name":                      "ServiceName",
		"service_registries":                "ServiceRegistries",
		"subnets":                           "Subnets",
		"tags":                              "Tags",
		"target_group_arn":                  "TargetGroupArn",
		"task_definition":                   "TaskDefinition",
		"type":                              "Type",
		"value":                             "Value",
		"weight":                            "Weight",
	})

	opts = opts.WithCreateTimeoutInMinutes(180).WithDeleteTimeoutInMinutes(30)

	opts = opts.WithUpdateTimeoutInMinutes(180)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
