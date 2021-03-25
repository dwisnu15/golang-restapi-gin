package configs

//this file concerns setting up the database connection
import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)


//var GormDB *gorm.DB
var DB *sql.DB

func InitViperConfig() {

	//get config file path
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(path + "/configs")
	//set the file name of config file
	viper.SetConfigName("dbconfig")
	//set path to look for config file (in: GinAPI)

	viper.SetConfigType("yml")
	var configurations Configurations

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	//db name is gin_framework
	viper.SetDefault("database.dbname", "gin_framework")

	err = viper.Unmarshal(&configurations)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configurations.Database.DBName)
	fmt.Println("Port is\t\t", configurations.Server.Port)
}

func InitDBConnection() *sql.DB {
	InitViperConfig()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"))
	engine := viper.GetString("database.engine")

	client, err := sql.Open(engine, psqlconn)
	if err != nil {
		log.Fatalf("Error database configurations, %v", err)
	}
	err = client.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	DB = client
	return DB
}

//func InitGormPostgres() {
//
//	//prepare connection string {host, port, username, password, dbname}
//	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		host,
//		port,
//		user,
//		pass,
//		dbname)
//
//	//create connection to database named 'gin_framework' using postgres dialect
//	database, err := gorm.Open("postgres", psqlconn)
//	if err != nil {
//		panic(err)
//	}
//
//	//migrate database schema on 'Items' model
//	//Make sure you call this method everytime you created a new model
//	database.AutoMigrate(&Items{})
//
//	GormDB = database
//}