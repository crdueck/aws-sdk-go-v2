// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an application. An application consists of one or more server groups.
// Each server group contain one or more servers.
func (c *Client) CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) {
	if params == nil {
		params = &CreateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApp", params, optFns, c.addOperationCreateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of application creation.
	ClientToken *string

	// The description of the new application
	Description *string

	// The name of the new application.
	Name *string

	// The name of the service role in the customer's account to be used by Server
	// Migration Service.
	RoleName *string

	// The server groups to include in the application.
	ServerGroups []types.ServerGroup

	// The tags to be associated with the application.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAppOutput struct {

	// A summary description of the application.
	AppSummary *types.AppSummary

	// The server groups included in the application.
	ServerGroups []types.ServerGroup

	// The tags associated with the application.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateApp{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "CreateApp",
	}
}
