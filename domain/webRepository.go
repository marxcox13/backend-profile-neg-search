/*PACKAGE*/
package WebRespository

import (
	"context"
	"seachProfile/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

/**/

/*IMPORTS FROM PACKAGE*/

/**/

func configDynamo(ctx context.Context) (aws.Config, error) {
	return config.LoadDefaultConfig(
		ctx,
		config.WithRegion("us-east-1"),
	)
}

func GetUsersWebsPopularRepository(ctx context.Context) ([]models.Webs, error) {

	cfg, err := configDynamo(ctx)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.ScanInput{
		TableName: aws.String("webs"),
		Limit:     aws.Int32(10),
	}

	result, err := client.Scan(ctx, input)
	if err != nil {
		return nil, err
	}

	webs := []models.Webs{}
	for _, item := range result.Items {
		var web models.Webs
		err := attributevalue.UnmarshalMap(item, &web)
		if err != nil {
			return nil, err
		}
		webs = append(webs, web)
	}

	return webs, nil
}

func CreateWebRepository(ctx context.Context, web models.Webs) (int, error) {
	cfg, err := configDynamo(ctx)
	if err != nil {
		return 0, err
	}

	client := dynamodb.NewFromConfig(cfg)

	item, err := attributevalue.MarshalMap(web)
	if err != nil {
		return 0, nil
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("webs"),
		Item:      item,
	})

	return 200, nil
}
