// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a protected query that is started by AWS Clean Rooms.
func (c *Client) StartProtectedQuery(ctx context.Context, params *StartProtectedQueryInput, optFns ...func(*Options)) (*StartProtectedQueryOutput, error) {
	if params == nil {
		params = &StartProtectedQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartProtectedQuery", params, optFns, c.addOperationStartProtectedQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartProtectedQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartProtectedQueryInput struct {

	// A unique identifier for the membership to run this query against. Currently
	// accepts a membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// The details needed to write the query results.
	//
	// This member is required.
	ResultConfiguration *types.ProtectedQueryResultConfiguration

	// The protected SQL query parameters.
	//
	// This member is required.
	SqlParameters *types.ProtectedQuerySQLParameters

	// The type of the protected query to be started.
	//
	// This member is required.
	Type types.ProtectedQueryType

	noSmithyDocumentSerde
}

type StartProtectedQueryOutput struct {

	// The protected query.
	//
	// This member is required.
	ProtectedQuery *types.ProtectedQuery

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartProtectedQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartProtectedQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartProtectedQuery{}, middleware.After)
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
	if err = addOpStartProtectedQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartProtectedQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartProtectedQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "StartProtectedQuery",
	}
}
