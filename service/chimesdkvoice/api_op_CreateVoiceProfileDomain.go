// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a voice profile domain, a collection of voice profiles, their voice
// prints, and encrypted enrollment audio. Before creating any voice profiles, you
// must provide all notices and obtain all consents from the speaker as required
// under applicable privacy and biometrics laws, and as required under the AWS
// service terms (https://aws.amazon.com/service-terms/) for the Amazon Chime SDK.
// For more information about voice profile domains, see Using Amazon Chime SDK
// Voice Analytics (https://docs.aws.amazon.com/chime-sdk/latest/dg/pstn-voice-analytics.html)
// in the Amazon Chime SDK Developer Guide.
func (c *Client) CreateVoiceProfileDomain(ctx context.Context, params *CreateVoiceProfileDomainInput, optFns ...func(*Options)) (*CreateVoiceProfileDomainOutput, error) {
	if params == nil {
		params = &CreateVoiceProfileDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVoiceProfileDomain", params, optFns, c.addOperationCreateVoiceProfileDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVoiceProfileDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVoiceProfileDomainInput struct {

	// The name of the voice profile domain.
	//
	// This member is required.
	Name *string

	// The server-side encryption configuration for the request.
	//
	// This member is required.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The unique identifier for the client request. Use a different token for
	// different domain creation requests.
	ClientRequestToken *string

	// A description of the voice profile domain.
	Description *string

	// The tags assigned to the domain.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateVoiceProfileDomainOutput struct {

	// The requested voice profile domain.
	VoiceProfileDomain *types.VoiceProfileDomain

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVoiceProfileDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVoiceProfileDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVoiceProfileDomain{}, middleware.After)
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
	if err = addOpCreateVoiceProfileDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVoiceProfileDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVoiceProfileDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateVoiceProfileDomain",
	}
}
