// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to change the ability of your account to send email.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutAccountSendingAttributesRequest
type PutAccountSendingAttributesInput struct {
	_ struct{} `type:"structure"`

	// Enables or disables your account's ability to send email. Set to true to
	// enable email sending, or set to false to disable email sending.
	//
	// If AWS paused your account's ability to send email, you can't use this operation
	// to resume your account's ability to send email.
	SendingEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s PutAccountSendingAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutAccountSendingAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.SendingEnabled != nil {
		v := *s.SendingEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SendingEnabled", protocol.BoolValue(v), metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutAccountSendingAttributesResponse
type PutAccountSendingAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAccountSendingAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutAccountSendingAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutAccountSendingAttributes = "PutAccountSendingAttributes"

// PutAccountSendingAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Enable or disable the ability of your account to send email.
//
//    // Example sending a request using PutAccountSendingAttributesRequest.
//    req := client.PutAccountSendingAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutAccountSendingAttributes
func (c *Client) PutAccountSendingAttributesRequest(input *PutAccountSendingAttributesInput) PutAccountSendingAttributesRequest {
	op := &aws.Operation{
		Name:       opPutAccountSendingAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/account/sending",
	}

	if input == nil {
		input = &PutAccountSendingAttributesInput{}
	}

	req := c.newRequest(op, input, &PutAccountSendingAttributesOutput{})
	return PutAccountSendingAttributesRequest{Request: req, Input: input, Copy: c.PutAccountSendingAttributesRequest}
}

// PutAccountSendingAttributesRequest is the request type for the
// PutAccountSendingAttributes API operation.
type PutAccountSendingAttributesRequest struct {
	*aws.Request
	Input *PutAccountSendingAttributesInput
	Copy  func(*PutAccountSendingAttributesInput) PutAccountSendingAttributesRequest
}

// Send marshals and sends the PutAccountSendingAttributes API request.
func (r PutAccountSendingAttributesRequest) Send(ctx context.Context) (*PutAccountSendingAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAccountSendingAttributesResponse{
		PutAccountSendingAttributesOutput: r.Request.Data.(*PutAccountSendingAttributesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAccountSendingAttributesResponse is the response type for the
// PutAccountSendingAttributes API operation.
type PutAccountSendingAttributesResponse struct {
	*PutAccountSendingAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAccountSendingAttributes request.
func (r *PutAccountSendingAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}