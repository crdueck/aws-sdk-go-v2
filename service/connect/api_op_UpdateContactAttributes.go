// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates user-defined contact attributes associated with the
// specified contact. You can create or update user-defined attributes for both
// ongoing and completed contacts. For example, while the call is active, you can
// update the customer's name or the reason the customer called. You can add notes
// about steps that the agent took during the call that display to the next agent
// that takes the call. You can also update attributes for a contact using data
// from your CRM application and save the data with the contact in Amazon Connect.
// You could also flag calls for additional analysis, such as legal review or to
// identify abusive callers. Contact attributes are available in Amazon Connect for
// 24 months, and are then deleted. For information about contact record retention
// and the maximum size of the contact record attributes section, see Feature
// specifications (https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits)
// in the Amazon Connect Administrator Guide.
func (c *Client) UpdateContactAttributes(ctx context.Context, params *UpdateContactAttributesInput, optFns ...func(*Options)) (*UpdateContactAttributesOutput, error) {
	if params == nil {
		params = &UpdateContactAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContactAttributes", params, optFns, c.addOperationUpdateContactAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContactAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContactAttributesInput struct {

	// The Amazon Connect attributes. These attributes can be accessed in flows just
	// like any other contact attributes. You can have up to 32,768 UTF-8 bytes across
	// all attributes for a contact. Attribute keys can include only alphanumeric,
	// dash, and underscore characters.
	//
	// This member is required.
	Attributes map[string]string

	// The identifier of the contact. This is the identifier of the contact associated
	// with the first interaction with the contact center.
	//
	// This member is required.
	InitialContactId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	noSmithyDocumentSerde
}

type UpdateContactAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContactAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateContactAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateContactAttributes{}, middleware.After)
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
	if err = addOpUpdateContactAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContactAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContactAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateContactAttributes",
	}
}
