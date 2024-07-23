package app

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/codescalersinternships/Secret-noteAPI-SPA-NabilaSherif/middlewares"
	"github.com/codescalersinternships/Secret-noteAPI-SPA-NabilaSherif/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	port      string
	DB        models.GormDB
	router    *gin.Engine
	secretKey string
}

var ErrConnectingToDB = errors.New("failed to connect to db")
var ErrMigrating = errors.New("db migration failed")

func NewApp(port, dbFilePath, secretkey string) (App, error) {
	db, err := models.Connect(dbFilePath)
	if err != nil {
		return App{}, ErrConnectingToDB
	}
	err = db.Migrate()
	if err != nil {
		return App{}, ErrMigrating
	}
	middlewares.SetSecretKey(secretkey)
	app := App{port: port, DB: db, router: gin.Default(), secretKey: secretkey}
	return app, nil
}

func (app *App) Run(listenAddress string) error {
	c := cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	})
	app.router.Use(c)
	app.Routes()
	return app.router.Run(listenAddress)
}

func (app *App) Routes() {
	secretnotebase := app.router.Group("/secretnote")
	{
		secretnotebase.Any("/:username/createnote", middlewares.RequireAuthorization, app.createNote)
		secretnotebase.GET("/note/:url", app.getNote)
		secretnotebase.Any("/createuser", app.createUser)
		secretnotebase.Any("/login", app.loginUser)
		secretnotebase.GET("/:username/notes", middlewares.RequireAuthorization, app.getUserNotes)
		secretnotebase.POST("/logout", middlewares.RequireAuthorization, app.logoutUser)
	}
}

func (app *App) createNote(c *gin.Context) {
	fmt.Println("inside create note")
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found in token"})
		return
	}
	note, err := app.DB.NewNote(note.NoteText, note.ExpirationDate, note.MaxViews, fmt.Sprint(username))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, note)
}

func (app *App) getNote(c *gin.Context) {
	noteurl := c.Param("url")
	note, err := app.DB.GetNote(noteurl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, note)
}

func (app *App) getUserNotes(c *gin.Context) {
	username := c.Param("username")
	notes, err := app.DB.GetUserNotes(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func (app *App) createUser(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := app.DB.NewUser(user.Name, user.Password)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
}

func (app *App) loginUser(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := app.DB.LogIn(user.Name, user.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		expiration := time.Now().Add(24 * time.Hour)
		fmt.Println("expirationnnnnn :", expiration)
		token, err := middlewares.GenerateToken(user.Name, app.secretKey, time.Now().Add(24*time.Hour))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(token)
		c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
}
func (app *App) logoutUser(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/secretnote")
}
