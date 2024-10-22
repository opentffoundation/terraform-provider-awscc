// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package glue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_glue_job", jobDataSource)
}

// jobDataSource returns the Terraform awscc_glue_job data source.
// This Terraform data source corresponds to the CloudFormation AWS::Glue::Job resource.
func jobDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllocatedCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of capacity units that are allocated to this job.",
		//	  "type": "number"
		//	}
		"allocated_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of capacity units that are allocated to this job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Command
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The code that executes a job.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of the job command",
		//	      "type": "string"
		//	    },
		//	    "PythonVersion": {
		//	      "description": "The Python version being used to execute a Python shell job.",
		//	      "type": "string"
		//	    },
		//	    "Runtime": {
		//	      "description": "Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment",
		//	      "type": "string"
		//	    },
		//	    "ScriptLocation": {
		//	      "description": "Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"command": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the job command",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PythonVersion
				"python_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Python version being used to execute a Python shell job.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Runtime
				"runtime": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ScriptLocation
				"script_location": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The code that executes a job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Connections
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the connections used by a job",
		//	  "properties": {
		//	    "Connections": {
		//	      "description": "A list of connections used by the job.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"connections": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Connections
				"connections": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of connections used by the job.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the connections used by a job",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultArguments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default arguments for this job, specified as name-value pairs.",
		//	  "type": "object"
		//	}
		"default_arguments": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The default arguments for this job, specified as name-value pairs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the job.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExecutionClass
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the job is run with a standard or flexible execution class.",
		//	  "type": "string"
		//	}
		"execution_class": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the job is run with a standard or flexible execution class.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExecutionProperty
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The maximum number of concurrent runs that are allowed for this job.",
		//	  "properties": {
		//	    "MaxConcurrentRuns": {
		//	      "description": "The maximum number of concurrent runs allowed for the job.",
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"execution_property": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxConcurrentRuns
				"max_concurrent_runs": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum number of concurrent runs allowed for the job.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The maximum number of concurrent runs that are allowed for this job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlueVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Glue version determines the versions of Apache Spark and Python that AWS Glue supports.",
		//	  "type": "string"
		//	}
		"glue_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Glue version determines the versions of Apache Spark and Python that AWS Glue supports.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: JobMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Property description not available.",
		//	  "type": "string"
		//	}
		"job_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Property description not available.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: JobRunQueuingEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Property description not available.",
		//	  "type": "boolean"
		//	}
		"job_run_queuing_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Property description not available.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This field is reserved for future use.",
		//	  "type": "string"
		//	}
		"log_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This field is reserved for future use.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Property description not available.",
		//	  "type": "string"
		//	}
		"maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Property description not available.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.",
		//	  "type": "number"
		//	}
		"max_capacity": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxRetries
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum number of times to retry this job after a JobRun fails",
		//	  "type": "number"
		//	}
		"max_retries": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum number of times to retry this job after a JobRun fails",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name you assign to the job definition",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name you assign to the job definition",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NonOverridableArguments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Non-overridable arguments for this job, specified as name-value pairs.",
		//	  "type": "object"
		//	}
		"non_overridable_arguments": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "Non-overridable arguments for this job, specified as name-value pairs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotificationProperty
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies configuration properties of a notification.",
		//	  "properties": {
		//	    "NotifyDelayAfter": {
		//	      "description": "It is the number of minutes to wait before sending a job run delay notification after a job run starts",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"notification_property": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: NotifyDelayAfter
				"notify_delay_after": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "It is the number of minutes to wait before sending a job run delay notification after a job run starts",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies configuration properties of a notification.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NumberOfWorkers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of workers of a defined workerType that are allocated when a job runs.",
		//	  "type": "integer"
		//	}
		"number_of_workers": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of workers of a defined workerType that are allocated when a job runs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Role
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name or Amazon Resource Name (ARN) of the IAM role associated with this job.",
		//	  "type": "string"
		//	}
		"role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name or Amazon Resource Name (ARN) of the IAM role associated with this job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SecurityConfiguration structure to be used with this job.",
		//	  "type": "string"
		//	}
		"security_configuration": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SecurityConfiguration structure to be used with this job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to use with this job.",
		//	  "type": "object"
		//	}
		"tags": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The tags to use with this job.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Timeout
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.",
		//	  "type": "integer"
		//	}
		"timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkerType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "TThe type of predefined worker that is allocated when a job runs.",
		//	  "enum": [
		//	    "Standard",
		//	    "G.1X",
		//	    "G.2X",
		//	    "G.025X",
		//	    "G.4X",
		//	    "G.8X",
		//	    "Z.2X"
		//	  ],
		//	  "type": "string"
		//	}
		"worker_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "TThe type of predefined worker that is allocated when a job runs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Glue::Job",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::Job").WithTerraformTypeName("awscc_glue_job")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocated_capacity":        "AllocatedCapacity",
		"command":                   "Command",
		"connections":               "Connections",
		"default_arguments":         "DefaultArguments",
		"description":               "Description",
		"execution_class":           "ExecutionClass",
		"execution_property":        "ExecutionProperty",
		"glue_version":              "GlueVersion",
		"job_mode":                  "JobMode",
		"job_run_queuing_enabled":   "JobRunQueuingEnabled",
		"log_uri":                   "LogUri",
		"maintenance_window":        "MaintenanceWindow",
		"max_capacity":              "MaxCapacity",
		"max_concurrent_runs":       "MaxConcurrentRuns",
		"max_retries":               "MaxRetries",
		"name":                      "Name",
		"non_overridable_arguments": "NonOverridableArguments",
		"notification_property":     "NotificationProperty",
		"notify_delay_after":        "NotifyDelayAfter",
		"number_of_workers":         "NumberOfWorkers",
		"python_version":            "PythonVersion",
		"role":                      "Role",
		"runtime":                   "Runtime",
		"script_location":           "ScriptLocation",
		"security_configuration":    "SecurityConfiguration",
		"tags":                      "Tags",
		"timeout":                   "Timeout",
		"worker_type":               "WorkerType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
