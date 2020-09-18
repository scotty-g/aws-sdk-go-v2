// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates an activity. An activity is a task that you write in any programming
// language and host on any machine that has access to AWS Step Functions.
// Activities must poll Step Functions using the GetActivityTask API action and
// respond using SendTask* API actions. This function lets Step Functions know the
// existence of your activity and returns an identifier for use in a state machine
// and when polling from the activity. This operation is eventually consistent. The
// results are best effort and may not reflect very recent updates and changes.
// CreateActivity is an idempotent API. Subsequent requests won’t create a
// duplicate resource if it was already created. CreateActivity's idempotency check
// is based on the activity name. If a following request has different tags values,
// Step Functions will ignore these differences and treat it as an idempotent
// request of the previous. In this case, tags will not be updated, even if they
// are different.
func (c *Client) CreateActivity(ctx context.Context, params *CreateActivityInput, optFns ...func(*Options)) (*CreateActivityOutput, error) {
	stack := middleware.NewStack("CreateActivity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpCreateActivityMiddlewares(stack)
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
	addOpCreateActivityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateActivity(options.Region), middleware.Before)

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
			OperationName: "CreateActivity",
			Err:           err,
		}
	}
	out := result.(*CreateActivityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateActivityInput struct {
	// The list of tags to add to a resource. An array of key-value pairs. For more
	// information, see Using Cost Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the AWS Billing and Cost Management User Guide, and Controlling Access Using
	// IAM Tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html). Tags
	// may only contain Unicode letters, digits, white space, or these symbols: _ . : /
	// = + - @.
	Tags []*types.Tag
	// The name of the activity to create. This name must be unique for your AWS
	// account and region for 90 days. For more information, see  Limits Related to
	// State Machine Executions
	// (https://docs.aws.amazon.com/step-functions/latest/dg/limits.html#service-limits-state-machine-executions)
	// in the AWS Step Functions Developer Guide. A name must not contain:
	//
	//     * white
	// space
	//
	//     * brackets < > { } [ ]
	//
	//     * wildcard characters ? *
	//
	//     * special
	// characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//     * control characters (U+0000-001F,
	// U+007F-009F)
	//
	// To enable logging with CloudWatch Logs, the name should only
	// contain 0-9, A-Z, a-z, - and _.
	Name *string
}

type CreateActivityOutput struct {
	// The Amazon Resource Name (ARN) that identifies the created activity.
	ActivityArn *string
	// The date the activity is created.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpCreateActivityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpCreateActivity{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateActivity{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateActivity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "CreateActivity",
	}
}