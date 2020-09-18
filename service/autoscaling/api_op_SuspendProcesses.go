// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Suspends the specified automatic scaling processes, or all processes, for the
// specified Auto Scaling group. If you suspend either the Launch or Terminate
// process types, it can prevent other process types from functioning properly. For
// more information, see Suspending and Resuming Scaling Processes
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html)
// in the Amazon EC2 Auto Scaling User Guide. To resume processes that have been
// suspended, call the ResumeProcesses () API.
func (c *Client) SuspendProcesses(ctx context.Context, params *SuspendProcessesInput, optFns ...func(*Options)) (*SuspendProcessesOutput, error) {
	stack := middleware.NewStack("SuspendProcesses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSuspendProcessesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpSuspendProcessesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSuspendProcesses(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "SuspendProcesses",
			Err:           err,
		}
	}
	out := result.(*SuspendProcessesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SuspendProcessesInput struct {
	// One or more of the following processes:
	//
	//     * Launch
	//
	//     * Terminate
	//
	//     *
	// AddToLoadBalancer
	//
	//     * AlarmNotification
	//
	//     * AZRebalance
	//
	//     *
	// HealthCheck
	//
	//     * InstanceRefresh
	//
	//     * ReplaceUnhealthy
	//
	//     *
	// ScheduledActions
	//
	// If you omit this parameter, all processes are specified.
	ScalingProcesses []*string
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
}

type SuspendProcessesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSuspendProcessesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSuspendProcesses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSuspendProcesses{}, middleware.After)
}

func newServiceMetadataMiddleware_opSuspendProcesses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "SuspendProcesses",
	}
}