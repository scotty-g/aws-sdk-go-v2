// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information about the bootstrap actions associated with a cluster.
func (c *Client) ListBootstrapActions(ctx context.Context, params *ListBootstrapActionsInput, optFns ...func(*Options)) (*ListBootstrapActionsOutput, error) {
	if params == nil {
		params = &ListBootstrapActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBootstrapActions", params, optFns, addOperationListBootstrapActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBootstrapActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which bootstrap actions to retrieve.
type ListBootstrapActionsInput struct {

	// The cluster identifier for the bootstrap actions to list.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string
}

// This output contains the bootstrap actions detail.
type ListBootstrapActionsOutput struct {

	// The bootstrap actions associated with the cluster.
	BootstrapActions []*types.Command

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBootstrapActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBootstrapActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBootstrapActions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListBootstrapActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBootstrapActions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListBootstrapActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListBootstrapActions",
	}
}
