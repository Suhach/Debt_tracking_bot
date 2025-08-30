package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/Suhach/Debt_tracking_bot/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var Pool *pgxpool.Pool
													
func InitPGX() {
	DB_URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dbpool, err := pgxpool.New(context.Background(), DB_URL)
	if err != nil {
		logger.Log.Fatal("Unable to connect to database: ", zap.Error(err))
	}
	defer dbpool.Close()
}
