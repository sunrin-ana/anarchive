package main

import (
	"anarchive/entity"
	"anarchive/pkg"
	"crypto/tls"
	"database/sql"
	"os"
	"time"

	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/boot"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

func newBunDB() *bun.DB {
	host := pkg.GetEnv("DB_HOST", "localhost")
	port := pkg.GetEnv("DB_PORT", "5437")
	user := pkg.GetEnv("DB_USER", "test")
	password := pkg.GetEnv("DB_PASSWORD", "test")
	database := pkg.GetEnv("DB_NAME", "test")

	tlsConfig := &tls.Config{
		InsecureSkipVerify: os.Getenv("DB_TLS_SKIP_VERIFY") == "true",
	}

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(host+":"+port),
		pgdriver.WithTLSConfig(tlsConfig),
		pgdriver.WithUser(user),
		pgdriver.WithPassword(password),
		pgdriver.WithDatabase(database),
		pgdriver.WithApplicationName("anarchive"),
		pgdriver.WithTimeout(10*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(10*time.Second),
		pgdriver.WithWriteTimeout(10*time.Second),
		pgdriver.WithInsecure(true),
	)

	sqldb := sql.OpenDB(pgconn)

	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(5)
	sqldb.SetConnMaxLifetime(5 * time.Minute)
	sqldb.SetConnMaxIdleTime(10 * time.Minute)

	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}

func main() {
	if err := pkg.InitLogger(); err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer pkg.Logger.Sync()

	logger := pkg.GetLogger()
	logger.Info("Starting Analog Server")

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	app := spine.New()

	db := newBunDB()

	db.RegisterModel(
		// 모델
		(*entity.Archive)(nil),
		(*entity.ArchiveRef)(nil),

		(*entity.Bucket)(nil),
		(*entity.BucketRef)(nil),

		(*entity.Anhref)(nil),

		(*entity.UploadSession)(nil),
		(*entity.FilePart)(nil),
	)

	app.Constructor(
		// 디비
		func() *bun.DB { return db },

		// 로거
		func() *zap.Logger { return logger },
	)

	port := pkg.GetEnv("SERVER_PORT", "8080")
	logger.Info("Server starting", zap.String("port", port))
	app.Run(boot.Options{
		Address:                ":" + port,
		EnableGracefulShutdown: true,
	})
}
