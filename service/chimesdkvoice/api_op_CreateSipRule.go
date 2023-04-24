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

// Creates a SIP rule, which can be used to run a SIP media application as a
// target for a specific trigger type. For more information about SIP rules, see
// Managing SIP media applications and rules (https://docs.aws.amazon.com/chime-sdk/latest/ag/manage-sip-applications.html)
// in the Amazon Chime SDK Administrator Guide.
func (c *Client) CreateSipRule(ctx context.Context, params *CreateSipRuleInput, optFns ...func(*Options)) (*CreateSipRuleOutput, error) {
	if params == nil {
		params = &CreateSipRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSipRule", params, optFns, c.addOperationCreateSipRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSipRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSipRuleInput struct {

	// The name of the SIP rule.
	//
	// This member is required.
	Name *string

	// The type of trigger assigned to the SIP rule in TriggerValue , currently
	// RequestUriHostname or ToPhoneNumber .
	//
	// This member is required.
	TriggerType types.SipRuleTriggerType

	// If TriggerType is RequestUriHostname , the value can be the outbound host name
	// of a Voice Connector. If TriggerType is ToPhoneNumber , the value can be a
	// customer-owned phone number in the E164 format. The SipMediaApplication
	// specified in the SipRule is triggered if the request URI in an incoming SIP
	// request matches the RequestUriHostname , or if the To header in the incoming
	// SIP request matches the ToPhoneNumber value.
	//
	// This member is required.
	TriggerValue *string

	// Disables or enables a SIP rule. You must disable SIP rules before you can
	// delete them.
	Disabled *bool

	// List of SIP media applications, with priority and AWS Region. Only one SIP
	// application per AWS Region can be used.
	TargetApplications []types.SipRuleTargetApplication

	noSmithyDocumentSerde
}

type CreateSipRuleOutput struct {

	// The SIP rule information, including the rule ID, triggers, and target
	// applications.
	SipRule *types.SipRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSipRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSipRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSipRule{}, middleware.After)
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
	if err = addOpCreateSipRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSipRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSipRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateSipRule",
	}
}
