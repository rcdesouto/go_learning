package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title"`
	Info  ItemInfo `json:"info"`
}

func main() {
	var myCreds = credentials.NewSharedCredentials("/Users/jpcutler/.aws/config", "jpc-test")
	fmt.Println("myCreds", myCreds)
	var awsConfig = &aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: myCreds,
	}

	fmt.Println("Credentials", awsConfig)

	sess, err := session.NewSession(awsConfig)

	if err != nil {
		fmt.Println("Session error: %s", err)
		os.Exit(1)
	}

	svc := dynamodb.New(sess)

	fmt.Println("List Tables")
	listTables(svc)
	fmt.Println("*********************")
	fmt.Println("Create Table")
	createTable(svc)
	printTable(svc, "Movies")
	fmt.Println("*********************")
	createFromJson(svc)
	fmt.Println("*********************")
	readTableItem(svc)
	fmt.Println("*********************")
	readTableItemScan(svc)
	fmt.Println("*********************")
	updateTableItem(svc)
	readTableItemScan(svc)
	fmt.Println("*********************")
	deleteTableItem(svc)
	readTableItemScan(svc)
}

func listTables(svc *dynamodb.DynamoDB) (err error) {
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		fmt.Println("Listing error:", err)
		os.Exit(1)
	}

	fmt.Println("Tables:")
	fmt.Println("")

	for _, n := range result.TableNames {
		fmt.Println(*n)
	}
	return err
}

func printTable(svc *dynamodb.DynamoDB, tName string) (err error) {
	req := &dynamodb.DescribeTableInput{
		TableName: &tName,
	}
	result, err := svc.DescribeTable(req)
	if err != nil {
		return err
	}
	table := result.Table
	fmt.Printf("done", table)
	return err
}

func createTable(svc *dynamodb.DynamoDB) (err error) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Movies"),
	}

	_, err = svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:", err)
		return err
	}

	fmt.Println("Create the table movies")
	return err
}

func createFromJson(svc *dynamodb.DynamoDB) {
	info := ItemInfo{
		Plot:   "Nothing happens at all.",
		Rating: 1.0,
	}

	item := Item{
		Year:  2015,
		Title: "The Big New Movie",
		Info:  info,
	}

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Movies"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added 'The Big New Movie' (2015) to Movies table")
}

func readTableItem(svc *dynamodb.DynamoDB) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String("2015"),
			},
			"title": {
				S: aws.String("The Big New Movie"),
			},
		},
	})

	if err != nil {
		fmt.Println(error.Error)
		return
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if item.Title == "" {
		fmt.Println("Could not find 'The Big New Movie' (2015)")
		return
	}

	fmt.Println("Found item:")
	fmt.Println("Year:   ", item.Year)
	fmt.Println("Title:  ", item.Title)
	fmt.Println("Plot:   ", item.Info.Plot)
	fmt.Println("Rating: ", item.Info.Rating)
}

func readTableItemScan(svc *dynamodb.DynamoDB) {
	fmt.Println("readTableItemScan")
	year := 2015
	num_items := 0
	filt := expression.Name("year").Equal(expression.Value(year))

	proj := expression.NamesList(expression.Name("title"), expression.Name("year"), expression.Name("info.rating"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Movies"),
	}

	result, err := svc.Scan(params)

	if len(result.Items) == 0 {
		fmt.Println("No items found matched the criteria")
	}

	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if item.Info.Rating > 0 {
			num_items += 1

			fmt.Println("Title: ", item.Title)
			fmt.Println("Rating:", item.Info.Rating)
			fmt.Println()
		}
	}
}

func updateTableItem(svc *dynamodb.DynamoDB) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String("0.5"),
			},
		},
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String("2015"),
			},
			"title": {
				S: aws.String("The Big New Movie"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set info.rating = :r"),
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated 'The Big New Movie' (2015) rating to 0.5")
}

func deleteTableItem(svc *dynamodb.DynamoDB) {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String("2015"),
			},
			"title": {
				S: aws.String("The Big New Movie"),
			},
		},
		TableName: aws.String("Movies"),
	}

	_, err := svc.DeleteItem(input)

	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Deleted 'The Big New Movie' (2015)")
}
