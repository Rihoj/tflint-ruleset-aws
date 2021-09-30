// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerNotebookInstanceInvalidInstanceTypeRule checks the pattern is valid
type AwsSagemakerNotebookInstanceInvalidInstanceTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSagemakerNotebookInstanceInvalidInstanceTypeRule returns new rule with default attributes
func NewAwsSagemakerNotebookInstanceInvalidInstanceTypeRule() *AwsSagemakerNotebookInstanceInvalidInstanceTypeRule {
	return &AwsSagemakerNotebookInstanceInvalidInstanceTypeRule{
		resourceType:  "aws_sagemaker_notebook_instance",
		attributeName: "instance_type",
		enum: []string{
			"ml.t2.medium",
			"ml.t2.large",
			"ml.t2.xlarge",
			"ml.t2.2xlarge",
			"ml.t3.medium",
			"ml.t3.large",
			"ml.t3.xlarge",
			"ml.t3.2xlarge",
			"ml.m4.xlarge",
			"ml.m4.2xlarge",
			"ml.m4.4xlarge",
			"ml.m4.10xlarge",
			"ml.m4.16xlarge",
			"ml.m5.xlarge",
			"ml.m5.2xlarge",
			"ml.m5.4xlarge",
			"ml.m5.12xlarge",
			"ml.m5.24xlarge",
			"ml.m5d.large",
			"ml.m5d.xlarge",
			"ml.m5d.2xlarge",
			"ml.m5d.4xlarge",
			"ml.m5d.8xlarge",
			"ml.m5d.12xlarge",
			"ml.m5d.16xlarge",
			"ml.m5d.24xlarge",
			"ml.c4.xlarge",
			"ml.c4.2xlarge",
			"ml.c4.4xlarge",
			"ml.c4.8xlarge",
			"ml.c5.xlarge",
			"ml.c5.2xlarge",
			"ml.c5.4xlarge",
			"ml.c5.9xlarge",
			"ml.c5.18xlarge",
			"ml.c5d.xlarge",
			"ml.c5d.2xlarge",
			"ml.c5d.4xlarge",
			"ml.c5d.9xlarge",
			"ml.c5d.18xlarge",
			"ml.p2.xlarge",
			"ml.p2.8xlarge",
			"ml.p2.16xlarge",
			"ml.p3.2xlarge",
			"ml.p3.8xlarge",
			"ml.p3.16xlarge",
			"ml.p3dn.24xlarge",
			"ml.g4dn.xlarge",
			"ml.g4dn.2xlarge",
			"ml.g4dn.4xlarge",
			"ml.g4dn.8xlarge",
			"ml.g4dn.12xlarge",
			"ml.g4dn.16xlarge",
			"ml.r5.large",
			"ml.r5.xlarge",
			"ml.r5.2xlarge",
			"ml.r5.4xlarge",
			"ml.r5.8xlarge",
			"ml.r5.12xlarge",
			"ml.r5.16xlarge",
			"ml.r5.24xlarge",
		},
	}
}

// Name returns the rule name
func (r *AwsSagemakerNotebookInstanceInvalidInstanceTypeRule) Name() string {
	return "aws_sagemaker_notebook_instance_invalid_instance_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerNotebookInstanceInvalidInstanceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerNotebookInstanceInvalidInstanceTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerNotebookInstanceInvalidInstanceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerNotebookInstanceInvalidInstanceTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as instance_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
