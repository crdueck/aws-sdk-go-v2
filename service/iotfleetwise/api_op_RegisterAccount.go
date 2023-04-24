// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Registers your Amazon Web Services account, IAM, and Amazon Timestream
// resources so Amazon Web Services IoT FleetWise can transfer your vehicle data to
// the Amazon Web Services Cloud. For more information, including step-by-step
// procedures, see Setting up Amazon Web Services IoT FleetWise (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/setting-up.html)
// . An Amazon Web Services account is not the same thing as a "user account". An
// Amazon Web Services user (https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction_identity-management.html#intro-identity-users)
// is an identity that you create using Identity and Access Management (IAM) and
// takes the form of either an IAM user (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html)
// or an IAM role, both with credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
// . A single Amazon Web Services account can, and typically does, contain many
// users and roles.
func (c *Client) RegisterAccount(ctx context.Context, params *RegisterAccountInput, optFns ...func(*Options)) (*RegisterAccountOutput, error) {
	if params == nil {
		params = &RegisterAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterAccount", params, optFns, c.addOperationRegisterAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterAccountInput struct {

	// The registered Amazon Timestream resources that Amazon Web Services IoT
	// FleetWise edge agent software can transfer your vehicle data to.
	//
	// This member is required.
	TimestreamResources *types.TimestreamResources

	// The IAM resource that allows Amazon Web Services IoT FleetWise to send data to
	// Amazon Timestream.
	//
	// Deprecated: iamResources is no longer used or needed as input
	IamResources *types.IamResources

	noSmithyDocumentSerde
}

type RegisterAccountOutput struct {

	// The time the account was registered, in seconds since epoch (January 1, 1970 at
	// midnight UTC time).
	//
	// This member is required.
	CreationTime *time.Time

	// The registered IAM resource that allows Amazon Web Services IoT FleetWise to
	// send data to Amazon Timestream.
	//
	// This member is required.
	IamResources *types.IamResources

	// The time this registration was last updated, in seconds since epoch (January 1,
	// 1970 at midnight UTC time).
	//
	// This member is required.
	LastModificationTime *time.Time

	// The status of registering your Amazon Web Services account, IAM role, and
	// Timestream resources.
	//
	// This member is required.
	RegisterAccountStatus types.RegistrationStatus

	// The registered Amazon Timestream resources that Amazon Web Services IoT
	// FleetWise edge agent software can transfer your vehicle data to.
	//
	// This member is required.
	TimestreamResources *types.TimestreamResources

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRegisterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRegisterAccount{}, middleware.After)
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
	if err = addOpRegisterAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "RegisterAccount",
	}
}
