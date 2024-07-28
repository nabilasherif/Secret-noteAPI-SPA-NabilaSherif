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
	config := cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	//app.router.Use(cors.Default())
	app.router.Use(cors.New(config))
	app.Routes()
	return app.router.Run(listenAddress)
}

func (app *App) Routes() {
	secretnotebase := app.router.Group("/api")
	{
		secretnotebase.POST("/note", middlewares.RequireAuthorization, middlewares.RateLimiting, app.createNote)
		secretnotebase.GET("/note/:uuid", middlewares.RateLimiting, app.getNote)
		secretnotebase.POST("/users", middlewares.RateLimiting, app.createUser)
		secretnotebase.POST("/login", middlewares.RateLimiting, app.loginUser)
		secretnotebase.GET("/notes", middlewares.RequireAuthorization, app.getUserNotes)
		secretnotebase.POST("/logout", middlewares.RequireAuthorization, app.logoutUser)
	}
}

type NoteInput struct {
	NoteText       string    `json:"note_text" validate:"required"`
	ExpirationDate time.Time `json:"expiration_date" validate:"required"`
	MaxViews       int       `json:"max_viewers" validate:"required"`
}

func (app *App) createNote(c *gin.Context) {
	var input NoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Binding error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Binding error: %s", err.Error())})
		return
	}

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found in token"})
		return
	}

	note, err := app.DB.NewNote(input.NoteText, input.ExpirationDate, input.MaxViews, fmt.Sprint(username))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, note)
}

func (app *App) getNote(c *gin.Context) {
	noteurl := c.Param("uuid")
	note, err := app.DB.GetNote(noteurl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, note)
}

func (app *App) getUserNotes(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username not found in token"})
		return
	}
	notes, err := app.DB.GetUserNotes(fmt.Sprint(username))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func (app *App) createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := app.DB.NewUser(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}
	c.Header("Content-Type", "application/json") //remove?
	c.JSON(http.StatusCreated, user)
}

func (app *App) loginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := app.DB.LogIn(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	token, err := middlewares.GenerateToken(user.Name, app.secretKey, time.Now().Add(24*time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(token)

	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

func (app *App) logoutUser(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/secretnote")
}
