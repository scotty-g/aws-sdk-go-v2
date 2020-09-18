// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a running multiplex. If the multiplex isn't running, this action has no
// effect.
func (c *Client) StopMultiplex(ctx context.Context, params *StopMultiplexInput, optFns ...func(*Options)) (*StopMultiplexOutput, error) {
	stack := middleware.NewStack("StopMultiplex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStopMultiplexMiddlewares(stack)
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
	addOpStopMultiplexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopMultiplex(options.Region), middleware.Before)

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
			OperationName: "StopMultiplex",
			Err:           err,
		}
	}
	out := result.(*StopMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for StopMultiplexRequest
type StopMultiplexInput struct {
	// The ID of the multiplex.
	MultiplexId *string
}

// Placeholder documentation for StopMultiplexResponse
type StopMultiplexOutput struct {
	// A list of the multiplex output destinations.
	Destinations []*types.MultiplexOutputDestination
	// The number of programs in the multiplex.
	ProgramCount *int32
	// The current state of the multiplex.
	State types.MultiplexState
	// The name of the multiplex.
	Name *string
	// The unique id of the multiplex.
	Id *string
	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings
	// The unique arn of the multiplex.
	Arn *string
	// The number of currently healthy pipelines.
	PipelinesRunningCount *int32
	// A list of availability zones for the multiplex.
	AvailabilityZones []*string
	// A collection of key-value pairs.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStopMultiplexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStopMultiplex{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStopMultiplex{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopMultiplex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "StopMultiplex",
	}
}