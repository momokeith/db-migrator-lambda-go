package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
     _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/momokeith/wallet-lambda-go/rds"
	"gopkg.in/gormigrate.v1"
	"log"
	"github.com/jinzhu/gorm"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db := rds.GetDB()
	defer db.Close()

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201806051400",
			Migrate: func(tx *gorm.DB) error {
				type Wallet struct {
					Uuid string
					User_id string
				}
				return tx.AutoMigrate(&Wallet{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("wallets").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

		return events.APIGatewayProxyResponse{
		Body:	 "{\"migrated\":true}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}