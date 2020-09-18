// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves entries from the specified fleet's event log. You can specify a time
// range to limit the result set. Use the pagination parameters to retrieve results
// as a set of sequential pages. If successful, a collection of event log entries
// matching the request are returned. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * Describe fleets:
//
//         * DescribeFleetAttributes ()
//
//         *
// DescribeFleetCapacity ()
//
//         * DescribeFleetPortSettings ()
//
//         *
// DescribeFleetUtilization ()
//
//         * DescribeRuntimeConfiguration ()
//
//
// * DescribeEC2InstanceLimits ()
//
//         * DescribeFleetEvents ()
//
//     *
// UpdateFleetAttributes ()
//
//     * StartFleetActions () or StopFleetActions ()
func (c *Client) DescribeFleetEvents(ctx context.Context, params *DescribeFleetEventsInput, optFns ...func(*Options)) (*DescribeFleetEventsOutput, error) {
	stack := middleware.NewStack("DescribeFleetEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeFleetEventsMiddlewares(stack)
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
	addOpDescribeFleetEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetEvents(options.Region), middleware.Before)

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
			OperationName: "DescribeFleetEvents",
			Err:           err,
		}
	}
	out := result.(*DescribeFleetEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeFleetEventsInput struct {
	// Most recent date to retrieve event logs for. If no end time is specified, this
	// call returns entries from the specified start time up to the present. Format is
	// a number expressed in Unix time as milliseconds (ex: "1469498468.057").
	EndTime *time.Time
	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string
	// Earliest date to retrieve event logs for. If no start time is specified, this
	// call returns entries starting from when the fleet was created to the specified
	// end time. Format is a number expressed in Unix time as milliseconds (ex:
	// "1469498468.057").
	StartTime *time.Time
	// A unique identifier for a fleet to get event logs for. You can use either the
	// fleet ID or ARN value.
	FleetId *string
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32
}

// Represents the returned data in response to a request action.
type DescribeFleetEventsOutput struct {
	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string
	// A collection of objects containing event log entries for the specified fleet.
	Events []*types.Event

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeFleetEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFleetEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetEvents",
	}
}