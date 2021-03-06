// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation returns all the domain names registered with Amazon Route 53 for
// the current AWS account.
func (c *Client) ListDomains(ctx context.Context, params *ListDomainsInput, optFns ...func(*Options)) (*ListDomainsOutput, error) {
	if params == nil {
		params = &ListDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomains", params, optFns, addOperationListDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListDomains request includes the following elements.
type ListDomainsInput struct {

	// For an initial request for a list of domains, omit this element. If the number
	// of domains that are associated with the current AWS account is greater than the
	// value that you specified for MaxItems, you can use Marker to return additional
	// domains. Get the value of NextPageMarker from the previous response, and submit
	// another request that includes the value of NextPageMarker in the Marker element.
	// Constraints: The marker must match the value specified in the previous request.
	Marker *string

	// Number of domains to be returned. Default: 20
	MaxItems *int32
}

// The ListDomains response includes the following elements.
type ListDomainsOutput struct {

	// A summary of domains.
	//
	// This member is required.
	Domains []*types.DomainSummary

	// If there are more domains than you specified for MaxItems in the request, submit
	// another request and include the value of NextPageMarker in the value of Marker.
	NextPageMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDomains{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomains(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ListDomains",
	}
}
