// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a pending maintenance action to a resource (for example, to a DB
// instance).
func (c *Client) ApplyPendingMaintenanceAction(ctx context.Context, params *ApplyPendingMaintenanceActionInput, optFns ...func(*Options)) (*ApplyPendingMaintenanceActionOutput, error) {
	if params == nil {
		params = &ApplyPendingMaintenanceActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplyPendingMaintenanceAction", params, optFns, c.addOperationApplyPendingMaintenanceActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplyPendingMaintenanceActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApplyPendingMaintenanceActionInput struct {

	// The pending maintenance action to apply to this resource. Valid values:
	// system-update , db-upgrade , hardware-maintenance , ca-certificate-rotation
	//
	// This member is required.
	ApplyAction *string

	// A value that specifies the type of opt-in request, or undoes an opt-in request.
	// An opt-in request of type immediate can't be undone. Valid values:
	//   - immediate - Apply the maintenance action immediately.
	//   - next-maintenance - Apply the maintenance action during the next maintenance
	//   window for the resource.
	//   - undo-opt-in - Cancel any existing next-maintenance opt-in requests.
	//
	// This member is required.
	OptInType *string

	// The RDS Amazon Resource Name (ARN) of the resource that the pending maintenance
	// action applies to. For information about creating an ARN, see Constructing an
	// RDS Amazon Resource Name (ARN) (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// .
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type ApplyPendingMaintenanceActionOutput struct {

	// Describes the pending maintenance actions for a resource.
	ResourcePendingMaintenanceActions *types.ResourcePendingMaintenanceActions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationApplyPendingMaintenanceActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpApplyPendingMaintenanceAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpApplyPendingMaintenanceAction{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpApplyPendingMaintenanceActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApplyPendingMaintenanceAction(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opApplyPendingMaintenanceAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ApplyPendingMaintenanceAction",
	}
}
