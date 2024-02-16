// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a delivery stream and its data. You can delete a delivery stream only
// if it is in one of the following states: ACTIVE , DELETING , CREATING_FAILED ,
// or DELETING_FAILED . You can't delete a delivery stream that is in the CREATING
// state. To check the state of a delivery stream, use DescribeDeliveryStream .
// DeleteDeliveryStream is an asynchronous API. When an API request to
// DeleteDeliveryStream succeeds, the delivery stream is marked for deletion, and
// it goes into the DELETING state.While the delivery stream is in the DELETING
// state, the service might continue to accept records, but it doesn't make any
// guarantees with respect to delivering the data. Therefore, as a best practice,
// first stop any applications that are sending records before you delete a
// delivery stream. Removal of a delivery stream that is in the DELETING state is
// a low priority operation for the service. A stream may remain in the DELETING
// state for several minutes. Therefore, as a best practice, applications should
// not wait for streams in the DELETING state to be removed.
func (c *Client) DeleteDeliveryStream(ctx context.Context, params *DeleteDeliveryStreamInput, optFns ...func(*Options)) (*DeleteDeliveryStreamOutput, error) {
	if params == nil {
		params = &DeleteDeliveryStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDeliveryStream", params, optFns, c.addOperationDeleteDeliveryStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDeliveryStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDeliveryStreamInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// Set this to true if you want to delete the delivery stream even if Firehose is
	// unable to retire the grant for the CMK. Firehose might be unable to retire the
	// grant due to a customer error, such as when the CMK or the grant are in an
	// invalid state. If you force deletion, you can then use the RevokeGrant (https://docs.aws.amazon.com/kms/latest/APIReference/API_RevokeGrant.html)
	// operation to revoke the grant you gave to Firehose. If a failure to retire the
	// grant happens due to an Amazon Web Services KMS issue, Firehose keeps retrying
	// the delete operation. The default value is false.
	AllowForceDelete *bool

	noSmithyDocumentSerde
}

type DeleteDeliveryStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDeliveryStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDeliveryStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteDeliveryStream"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteDeliveryStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDeliveryStream(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteDeliveryStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteDeliveryStream",
	}
}
