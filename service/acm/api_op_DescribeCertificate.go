// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/DescribeCertificateRequest
type DescribeCertificateInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the ACM certificate. The ARN must have
	// the following form:
	//
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// CertificateArn is a required field
	CertificateArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificateInput"}

	if s.CertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateArn"))
	}
	if s.CertificateArn != nil && len(*s.CertificateArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/DescribeCertificateResponse
type DescribeCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Metadata about an ACM certificate.
	Certificate *CertificateDetail `type:"structure"`
}

// String returns the string representation
func (s DescribeCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCertificate = "DescribeCertificate"

// DescribeCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager.
//
// Returns detailed metadata about the specified ACM certificate.
//
//    // Example sending a request using DescribeCertificateRequest.
//    req := client.DescribeCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-2015-12-08/DescribeCertificate
func (c *Client) DescribeCertificateRequest(input *DescribeCertificateInput) DescribeCertificateRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificateInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificateOutput{})
	return DescribeCertificateRequest{Request: req, Input: input, Copy: c.DescribeCertificateRequest}
}

// DescribeCertificateRequest is the request type for the
// DescribeCertificate API operation.
type DescribeCertificateRequest struct {
	*aws.Request
	Input *DescribeCertificateInput
	Copy  func(*DescribeCertificateInput) DescribeCertificateRequest
}

// Send marshals and sends the DescribeCertificate API request.
func (r DescribeCertificateRequest) Send(ctx context.Context) (*DescribeCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificateResponse{
		DescribeCertificateOutput: r.Request.Data.(*DescribeCertificateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificateResponse is the response type for the
// DescribeCertificate API operation.
type DescribeCertificateResponse struct {
	*DescribeCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificate request.
func (r *DescribeCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}