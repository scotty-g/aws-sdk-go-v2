// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing maintenance window. Only specified parameters are modified.
// The value you specify for Duration determines the specific end time for the
// maintenance window based on the time it begins. No maintenance window tasks are
// permitted to start after the resulting endtime minus the number of hours you
// specify for Cutoff. For example, if the maintenance window starts at 3 PM, the
// duration is three hours, and the value you specify for Cutoff is one hour, no
// maintenance window tasks can start after 5 PM.
func (c *Client) UpdateMaintenanceWindow(ctx context.Context, params *UpdateMaintenanceWindowInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowOutput, error) {
	stack := middleware.NewStack("UpdateMaintenanceWindow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateMaintenanceWindowMiddlewares(stack)
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
	addOpUpdateMaintenanceWindowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceWindow(options.Region), middleware.Before)

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
			OperationName: "UpdateMaintenanceWindow",
			Err:           err,
		}
	}
	out := result.(*UpdateMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMaintenanceWindowInput struct {
	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string
	// The duration of the maintenance window in hours.
	Duration *int32
	// Whether the maintenance window is enabled.
	Enabled *bool
	// Whether targets must be registered with the maintenance window before tasks can
	// be defined for those targets.
	AllowUnassociatedTargets *bool
	// An optional description for the update request.
	Description *string
	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "etc/UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string
	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int32
	// If True, then all fields that are required by the CreateMaintenanceWindow action
	// are also required for this API request. Optional fields that are not specified
	// are set to null.
	Replace *bool
	// The number of days to wait after the date and time specified by a CRON
	// expression before running the maintenance window. For example, the following
	// cron expression schedules a maintenance window to run the third Tuesday of every
	// month at 11:30 PM. cron(0 30 23 ? * TUE#3 *) If the schedule offset is 2, the
	// maintenance window won't run until two days later.
	ScheduleOffset *int32
	// The name of the maintenance window.
	Name *string
	// The ID of the maintenance window to update.
	WindowId *string
	// The date and time, in ISO-8601 Extended format, for when you want the
	// maintenance window to become inactive. EndDate allows you to set a date and time
	// in the future when the maintenance window will no longer run.
	EndDate *string
	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "etc/UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	StartDate *string
}

type UpdateMaintenanceWindowOutput struct {
	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "etc/UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string
	// The date and time, in ISO-8601 Extended format, for when the maintenance window
	// is scheduled to become active. The maintenance window will not run before this
	// specified time.
	StartDate *string
	// The duration of the maintenance window in hours.
	Duration *int32
	// The name of the maintenance window.
	Name *string
	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string
	// The date and time, in ISO-8601 Extended format, for when the maintenance window
	// is scheduled to become inactive. The maintenance window will not run after this
	// specified time.
	EndDate *string
	// Whether the maintenance window is enabled.
	Enabled *bool
	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int32
	// Whether targets must be registered with the maintenance window before tasks can
	// be defined for those targets.
	AllowUnassociatedTargets *bool
	// The ID of the created maintenance window.
	WindowId *string
	// An optional description of the update.
	Description *string
	// The number of days to wait to run a maintenance window after the scheduled CRON
	// expression date and time.
	ScheduleOffset *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateMaintenanceWindowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceWindow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceWindow{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMaintenanceWindow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateMaintenanceWindow",
	}
}