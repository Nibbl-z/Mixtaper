package main

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
)

func createAttribute(db *databases.Databases, collection string, key string, attributeType string, required bool, size int) {
	switch attributeType {
		case "string":
			db.CreateStringAttribute("mixtaper", collection, key, size, required)
		case "number":
			db.CreateFloatAttribute("mixtaper", collection, key, required, db.WithCreateFloatAttributeMin(-1000000), db.WithCreateFloatAttributeMax(1000000))
		case "array":
			db.CreateStringAttribute("mixtaper", collection, key, size, required, db.WithCreateStringAttributeArray(true))
	}
}

func createIndex(db *databases.Databases, collection string, key string, indexType string, attributes []string) {
	db.CreateIndex("mixtaper", collection, key, indexType, attributes)
}

type Attribute struct {
	Key string
	Type string
	Required bool
	Size int
}

type Index struct {
	Key string
	Type string
	Attributes []string
}

var LEVEL_ATTRIBUTES = []Attribute{
	{ Key: "songName", Type: "string", Required: true, Size: 200, },
	{ Key: "songArtist", Type: "string", Required: true, Size: 200, },
	{ Key: "description", Type: "string", Required: false, Size: 2000, },
	{ Key: "chartName", Type: "string", Required: false, Size: 200, },
	{ Key: "uploader", Type: "string", Required: true, Size: 100, },
	{ Key: "bpm", Type: "number", Required: false, Size: 200, },
	{ Key: "gamesUsed", Type: "array", Required: false, Size: 1000000, },
	{ Key: "youtubeVideo", Type: "string", Required: false, Size: 1000000, },
}

var LEVEL_INDEXES = []Index{
	{ Key: "uploader", Type: "key", Attributes: []string{"uploader"}},
	{ Key: "songName", Type: "fulltext", Attributes: []string{"songName"}},
	{ Key: "songArtist", Type: "fulltext", Attributes: []string{"songArtist"}},
	{ Key: "description", Type: "fulltext", Attributes: []string{"description"}},
	{ Key: "chartName", Type: "fulltext", Attributes: []string{"chartName"}},
	{ Key: "uploadDate", Type: "key", Attributes: []string{"$createdAt"}},
}

func createLevelAttributes(db *databases.Databases) {
	for _, attribute := range LEVEL_ATTRIBUTES {
		createAttribute(db, "levels", attribute.Key, attribute.Type, attribute.Required, attribute.Size)
	}
}

func createLevelIndexes(db *databases.Databases) {
	for _, index := range LEVEL_INDEXES {
		createIndex(db, "levels", index.Key, index.Type, index.Attributes)
	}
}

func main() {
	err := godotenv.Load()
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating client")
	fmt.Println(os.Getenv("APPWRITE_API_ENDPOINT"))
	client := appwrite.NewClient(
		appwrite.WithEndpoint(os.Getenv("APPWRITE_API_ENDPOINT")),
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)

	///// Database

	fmt.Println("Creating database")

	database := appwrite.NewDatabases(client)

	_, err = database.Create(
		"mixtaper",
		"Mixtaper",
		database.WithCreateEnabled(true),
	)
	
	if err != nil {panic(err)}

	// Levels

	fmt.Println("Creating levels collection")

	_, err = database.CreateCollection(
		"mixtaper",
		"levels",
		"Levels",
		database.WithCreateCollectionPermissions([]string{
			permission.Read(role.Any()),
			permission.Create(role.Users("verified")),
		}),
		database.WithCreateCollectionDocumentSecurity(true),
		database.WithCreateCollectionEnabled(true),
	)
	
	if err != nil {panic(err)}

	createLevelAttributes(database)
	createLevelIndexes(database)

	// Usernames

	fmt.Println("Creating usernames collection")

	_, err = database.CreateCollection(
		"mixtaper",
		"usernames",
		"Usernames",
		database.WithCreateCollectionPermissions([]string{
			permission.Read(role.Any()),
			permission.Create(role.Users("verified")),
		}),
		database.WithCreateCollectionDocumentSecurity(true),
		database.WithCreateCollectionEnabled(true),
	)
	
	if err != nil {panic(err)}

	createAttribute(database, "usernames", "username", "string", true, 20)
	createIndex(database, "usernames", "id", "unique", []string{"username"})

	///// Storage

	fmt.Println("Creating RIQ bucket")

	storage := appwrite.NewStorage(client)

	_, err = storage.CreateBucket(
		"riq_files",
		"RIQs",
		storage.WithCreateBucketEnabled(true),
		storage.WithCreateBucketPermissions([]string{permission.Create(role.Users("verified"))}),
		storage.WithCreateBucketFileSecurity(true),
		storage.WithCreateBucketAntivirus(true),
		storage.WithCreateBucketEncryption(true),
		storage.WithCreateBucketMaximumFileSize(20971520),
	)
	if err != nil {panic(err)}

	fmt.Println("Creating profile picture bucket")

	_, err = storage.CreateBucket(
		"profile_pictures",
		"Profile Pictures",
		storage.WithCreateBucketEnabled(true),
		storage.WithCreateBucketPermissions([]string{
			permission.Create(role.Users("")),
			permission.Read(role.Any()),
		}),
		storage.WithCreateBucketFileSecurity(true),
		storage.WithCreateBucketAntivirus(true),
		storage.WithCreateBucketEncryption(true),
		storage.WithCreateBucketAllowedFileExtensions([]string{
			"png", "jpg", "jpeg",
		}),
		storage.WithCreateBucketMaximumFileSize(5242880),
	)
	if err != nil {panic(err)}

	fmt.Println("Creating cover art bucket")

	_, err = storage.CreateBucket(
		"cover_art",
		"Cover Art",
		storage.WithCreateBucketEnabled(true),
		storage.WithCreateBucketPermissions([]string{
			permission.Create(role.Users("verified")),
			permission.Read(role.Any()),
		}),
		storage.WithCreateBucketFileSecurity(true),
		storage.WithCreateBucketAntivirus(true),
		storage.WithCreateBucketEncryption(true),
		storage.WithCreateBucketAllowedFileExtensions([]string{
			"png", "jpg", "jpeg",
		}),
		storage.WithCreateBucketMaximumFileSize(5242880),
	)
	if err != nil {panic(err)}

	fmt.Println("Appwrite setup complete!")
}