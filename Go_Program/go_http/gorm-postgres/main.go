package main

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
    "time"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Email string `gorm:"unique"`

}

type Users struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Password string
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=host.docker.internal user=user password=password dbname=gormdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB = db
	fmt.Println("Connected to PostgreSQL!")
}

func CreateUser(name, email string) {
	user := User{Name: name, Email: email}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Println("Error creating user:", result.Error)
	}
	fmt.Printf("User %s created successfully!\n", user.Name)
}

func GetUsers() {
	var users []User
	DB.Find(&users)
	fmt.Println("Users:", users)
}

func UpdateUserEmail(id uint, newEmail string) {
	var user User
	DB.First(&user, id)
	user.Email = newEmail
	DB.Save(&user)
	fmt.Printf("User ID %d email updated!\n", id)
}

func DeleteUser(id uint) {
	DB.Delete(&User{}, id)
	fmt.Printf("User ID %d deleted!\n", id)
}

func MigrateDatabase() {
	DB.AutoMigrate(&User{})
	fmt.Println("Database migrated!")
}

var jwtKey = []byte("secretCode") // Replace with a secure secret key

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func generateToken(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
    claims := &Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func verifyToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
	ConnectDatabase()
	MigrateDatabase()

	ConnectDatabase()
	MigrateDatabase()

	CreateUser("John Doe", "john@example.com")
	CreateUser("Jane Doe", "jane@example.com")
	CreateUser("Jame Doe", "jame@example.com")

	GetUsers()
	UpdateUserEmail(1, "john.doe@updated.com")
	GetUsers()

	DeleteUser(2)
	GetUsers()
}