// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides information about a user's device.
func (c *Client) DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) {
	stack := middleware.NewStack("DescribeDevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeDeviceMiddlewares(stack)
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
	addOpDescribeDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDevice(options.Region), middleware.Before)

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
			OperationName: "DescribeDevice",
			Err:           err,
		}
	}
	out := result.(*DescribeDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeviceInput struct {
	// The ARN of the fleet.
	FleetArn *string
	// A unique identifier for a registered user's device.
	DeviceId *string
}

type DescribeDeviceOutput struct {
	// The operating system patch level of the device.
	PatchLevel *string
	// The user name associated with the device.
	Username *string
	// The operating system of the device.
	OperatingSystem *string
	// The current state of the device.
	Status types.DeviceStatus
	// The date that the device first signed in to Amazon WorkLink.
	FirstAccessedTime *time.Time
	// The date that the device last accessed Amazon WorkLink.
	LastAccessedTime *time.Time
	// The model of the device.
	Model *string
	// The manufacturer of the device.
	Manufacturer *string
	// The operating system version of the device.
	OperatingSystemVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeDeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDevice{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeDevice",
	}
}