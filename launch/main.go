package main

import (
	// "os"
	// "log"
	"fmt"
	"time"
	"unicode"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/blainehansen/goqu"
	_ "github.com/blainehansen/goqu/adapters/postgres"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"
	"github.com/iancoleman/strcase"

	"github.com/badoux/checkmail"

	"github.com/json-iterator/go/extra"

	"crypto/rand"
	"encoding/base64"

	"strconv"
	"gopkg.in/mailgun/mailgun-go.v1"
)


var environment map[string]string = func() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println("error reading .env file")
		panic(err)
	}
	return env
}()


var db *goqu.Database = func() *goqu.Database {
	// CONNECTING TO DATABASE
	pgDb, connectionError := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			environment["DOCKER_DATABASE_HOST"],
			environment["DATABASE_PORT"],
			environment["DATABASE_DB_NAME"],
			"golang_server_user",
			environment["GOLANG_DATABASE_PASSWORD"],
			environment["DATABASE_SSL"],
		),
	)
	if connectionError != nil {
		fmt.Println("failed to connect to database")
		panic(connectionError.Error())
	}

	finalDb := goqu.New("postgres", pgDb)

	_, err := finalDb.Exec("SHOW server_version")
	if err != nil {
		fmt.Println("failed to connect to database")
		panic(err.Error())
	}

	return finalDb
}()


func lowercaseFirstLetter(name string) string {
	out := []rune(name)
	out[0] = unicode.ToLower(out[0])
	return string(out)
}


var localEncoding *base64.Encoding = base64.RawURLEncoding
func encodeBase64(data []byte) []byte {
	encodedData := make([]byte, localEncoding.EncodedLen(len(data)))
	localEncoding.Encode(encodedData, data)
	return encodedData
}

func generateRandomToken() ([]byte, error) {
	tokenBytes := make([]byte, 64)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}

	return encodeBase64(tokenBytes), nil
}



var shouldMail bool = func() bool {
	innerBool, parseError := strconv.ParseBool(environment["SHOULD_MAIL"])
	if parseError != nil {
		panic(parseError)
	}

	return innerBool
}()

var domain string = environment["SERVER_DOMAIN"]
var privateAPIKey string = environment["MAIL_PRIVATE_API_KEY"]
var publicValidationKey string = environment["MAIL_PUBLIC_KEY"]

var mailgunClient mailgun.Mailgun = mailgun.NewMailgun(domain, privateAPIKey, publicValidationKey)

func sendMessage(sender string, subject string, body string, recipient string) error {
	message := mailgunClient.NewMessage(sender, subject, body, recipient)
	if shouldMail {
		_, _, err := mailgunClient.Send(message)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("fake emailed message:")
		fmt.Println(sender, subject, body, recipient)
	}

	return nil
}


func main() {
	extra.SetNamingStrategy(lowercaseFirstLetter)
	goqu.SetColumnRenameFunction(strcase.ToSnake)

	// SETTING UP ROUTER
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	// config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowMethods = []string{"OPTIONS", "POST"}
	config.MaxAge = 24 * time.Hour

	router.Use(cors.New(config))


	// sudo docker-compose -f docker-compose-launch.yml up
	router.POST("/new-email", func(c *gin.Context) {
		input := struct {
			Email string
		}{}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithError(422, err); return
		}

		if err := checkmail.ValidateFormat(input.Email); err != nil {
			c.AbortWithError(400, err); return
		}

		byteValidationToken, generationError := generateRandomToken()
		if generationError != nil {
			c.AbortWithError(500, generationError)
		}
		validationToken := string(byteValidationToken)

		insert := db.From("emails").Insert(
			goqu.Record{ "email": input.Email, "validation_token": validationToken },
		)
		if _, err := insert.Exec(); err != nil{
			c.AbortWithError(500, err); return
		}

		validationUrl := fmt.Sprintf(`%s%s/recover-password?t=%s`, environment["SERVER_PROTOCOL"], environment["SERVER_DOMAIN"], validationToken)

		body := "Hello! Thank you for signing up to join the Crowdsell private beta.\n\n" +
			"Click this link to validate your email: \n" +
			validationUrl

		sendMessage("no-reply@crowdsell.io", "Crowdsell - Validation Email", body, input.Email)

		c.Status(204)
	})

	router.POST("/validate", func(c *gin.Context) {
		input := struct {
			ValidationToken string
		}{}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithError(422, err); return
		}

		updateResult, updateErr := db.From("emails").Where(
			goqu.I("validation_token").Eq(input.ValidationToken),
		).Update(
			goqu.Record{ "validation_token": nil },
		).Exec()

		if updateErr != nil{
			c.AbortWithError(500, updateErr); return
		}

		rowsAffected, err := updateResult.RowsAffected()
		if rowsAffected != 1 {
			c.AbortWithError(400, fmt.Errorf("got a rowsAffected of %s\n", rowsAffected)); return
		}
		if err != nil {
			c.AbortWithError(500, err); return
		}

		c.Status(204)
	})


	router.Run(fmt.Sprintf(":%s", environment["API_PORT"]))
}
