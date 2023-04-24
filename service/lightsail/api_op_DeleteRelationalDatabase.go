// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a database in Amazon Lightsail. The delete relational database
// operation supports tag-based access control via resource tags applied to the
// resource identified by relationalDatabaseName. For more information, see the
// Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags)
// .
func (c *Client) DeleteRelationalDatabase(ctx context.Context, params *DeleteRelationalDatabaseInput, optFns ...func(*Options)) (*DeleteRelationalDatabaseOutput, error) {
	if params == nil {
		params = &DeleteRelationalDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRelationalDatabase", params, optFns, c.addOperationDeleteRelationalDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRelationalDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRelationalDatabaseInput struct {

	// The name of the database that you are deleting.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The name of the database snapshot created if skip final snapshot is false ,
	// which is the default value for that parameter. Specifying this parameter and
	// also specifying the skip final snapshot parameter to true results in an error.
	// Constraints:
	//   - Must contain from 2 to 255 alphanumeric characters, or hyphens.
	//   - The first and last character must be a letter or number.
	FinalRelationalDatabaseSnapshotName *string

	// Determines whether a final database snapshot is created before your database is
	// deleted. If true is specified, no database snapshot is created. If false is
	// specified, a database snapshot is created before your database is deleted. You
	// must specify the final relational database snapshot name parameter if the skip
	// final snapshot parameter is false . Default: false
	SkipFinalSnapshot *bool

	noSmithyDocumentSerde
}

type DeleteRelationalDatabaseOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRelationalDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRelationalDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRelationalDatabase{}, middleware.After)
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
	if err = addOpDeleteRelationalDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRelationalDatabase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRelationalDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DeleteRelationalDatabase",
	}
}
