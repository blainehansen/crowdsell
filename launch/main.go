package main

// docker build -f ../Dockerfile-crowdsell-go -t crowdsell-go .

import (
	// "log"
	"fmt"
	"time"
	"strings"
	"unicode"
	"io/ioutil"

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


func readEnvFile(filename string) map[string]string {
	env, err := godotenv.Read(filename)
	if err != nil {
		fmt.Printf("error reading %s\n", filename)
		panic(err)
	}
	return env
}

var env map[string]string = readEnvFile("/env/.env")
var keys map[string]string = readEnvFile("/keys/.keys")

var db *goqu.Database = func() *goqu.Database {
	bytesGolangDatabasePassword, readErr := ioutil.ReadFile("/keys/.keys.go-db")
	if readErr != nil {
		panic(readErr)
	}

	// CONNECTING TO DATABASE
	pgDb, connectionError := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=database port=5432 dbname=database user=golang_server_user password=%s sslmode=%s",
			strings.TrimSpace(string(bytesGolangDatabasePassword)),
			env["DATABASE_SSL"],
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


// TODO, the ci build scripts can kubectl apply/create different config maps
// use apply with multiple file switches depending on the env


func generateRandomToken() ([]byte, error) {
	tokenBytes := make([]byte, 64)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return nil, err
	}

	return encodeBase64(tokenBytes), nil
}



var shouldMail bool = func() bool {
	innerBool, parseError := strconv.ParseBool(env["SHOULD_MAIL"])
	if parseError != nil {
		panic(parseError)
	}

	return innerBool
}()

var serverDomain string = env["SERVER_DOMAIN"]
var serverProtocol string = env["SERVER_PROTOCOL"]

var mailPrivateAPIKey string = keys["MAIL_PRIVATE_API_KEY"]
var mailPublicKey string = keys["MAIL_PUBLIC_KEY"]

var mailgunClient mailgun.Mailgun = mailgun.NewMailgun(serverDomain, mailPrivateAPIKey, mailPublicKey)

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
	config.AllowOrigins = []string{env["ALLOW_ORIGIN"]}
	config.AllowMethods = []string{"OPTIONS", "POST"}
	config.MaxAge = 24 * time.Hour

	router.Use(cors.New(config))


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
			// in the event that a duplicate email comes in, we just pretend like everything's okay
			if strings.Contains(err.Error(), `unique constraint "emails_email_key"`) {
				c.Status(204); return
			}

			c.AbortWithError(500, err); return
		}

		validationUrl := fmt.Sprintf(`%s%s/recover-password?t=%s`, serverProtocol, serverDomain, validationToken)

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


	router.Run(fmt.Sprintf(":%s", "5050"))
}
