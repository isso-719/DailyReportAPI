package db

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conn DB 接続情報
var Conn *gorm.DB

// EnvLoad() .env から環境変数を取得する関数
func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Info("Info: No .env file found. Use System Environment Variables.")
	}
}

// init() 初期化処理
func init() {
	// 環境変数を .env からロードする
	envLoad()

	// DB接続情報
	// ユーザ
	user := os.Getenv("MYSQL_USER")
	// パスワード
	password := os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	host := os.Getenv("MYSQL_HOST")
	// 接続先ポート
	port := os.Getenv("MYSQL_PORT")
	// 接続先データベース
	database := os.Getenv("MYSQL_DATABASE")

	// MySQL に接続するための URL を作成する
	mySQLURL := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"

	// DB接続
	db, err := gorm.Open(mysql.Open(mySQLURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: DB connection failed.")
	}

	// DB接続情報を Conn に格納する
	Conn = db

}
