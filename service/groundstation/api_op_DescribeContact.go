// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an existing contact.
func (c *Client) DescribeContact(ctx context.Context, params *DescribeContactInput, optFns ...func(*Options)) (*DescribeContactOutput, error) {
	stack := middleware.NewStack("DescribeContact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeContactMiddlewares(stack)
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
	addOpDescribeContactValidationMiddleware(stack)

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
			OperationName: "DescribeContact",
			Err:           err,
		}
	}
	out := result.(*DescribeContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeContactInput struct {
	// UUID of a contact.
	ContactId *string
}

//
type DescribeContactOutput struct {
	// UUID of a contact.
	ContactId *string
	// Status of a contact.
	ContactStatus types.ContactStatus
	// End time of a contact.
	EndTime *time.Time
	// Error message for a contact.
	ErrorMessage *string
	// Ground station for a contact.
	GroundStation *string
	// Maximum elevation angle of a contact.
	MaximumElevation *types.Elevation
	// ARN of a mission profile.
	MissionProfileArn *string
	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	PostPassEndTime *time.Time
	// Amount of time prior to contact start you’d like to receive a CloudWatch event
	// indicating an upcoming pass.
	PrePassStartTime *time.Time
	// Region of a contact.
	Region *string
	// ARN of a satellite.
	SatelliteArn *string
	// Start time of a contact.
	StartTime *time.Time
	// Tags assigned to a contact.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeContactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeContact{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeContact{}, middleware.After)
}