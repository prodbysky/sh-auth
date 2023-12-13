package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/prodbysky/sh-auth/pkg/models"
)

var users []models.User

// Router related functions
func authHash() string {
    curr_time := time.Now().Unix()
    num := make([]byte, 8)
    num = binary.LittleEndian.AppendUint64(num, uint64(curr_time))
    sha := sha256.New()
    sha.Write(num)
    hashed := sha.Sum(nil)
    return hex.EncodeToString(hashed)
}

func status(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
    })
}

func getUser(c *gin.Context) {
	name := c.Query("name")
	pass := c.Query("password")
	if name == "" || pass == "" {
		return
	}

	user := models.User{Name: name, Password: pass}

    users = append(users, user)
    
    fmt.Println(users)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   fmt.Sprintf("User: %s, password: %s", user.Name, user.Password),
	})
}

// Data management functions
func readUsers(filename string, users *[]models.User) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatalf("Non-existent filename provided %s\n", filename)
    }

    var result []models.User
    err = json.Unmarshal(data, &result)
    if err != nil {
        log.Fatalln("Failed unmarshalling data")
    }
}

func saveUsers(filename string, users []models.User) {
    fmt.Println(users)
    bytes, err := json.Marshal(users)
    if err != nil {
        log.Fatalln("Failed to marshal users:", err)
    }

    fmt.Println("Marshaled JSON:", string(bytes))

    err = os.WriteFile(filename, bytes, 0644)
    if err != nil {
        log.Fatalln("Failed to save users to file:", err)
    }

}

func main() {
    filename := "users.dat"
    // Reading users
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatalf("Non-existent filename provided %s\n", filename)
    }

    err = json.Unmarshal(data, &users)
    if err != nil {
        log.Fatalln("Failed unmarshalling data")
    }

    router := gin.Default()
    router.GET("/", status)
    router.GET("/sendUser", getUser)
    endless.ListenAndServe(":8080", router)

    // Saving users
    bytes, err := json.Marshal(users)
    if err != nil {
        log.Fatalln("Failed to marshal users:", err)
    }

    fmt.Println("Marshaled JSON:", string(bytes))

    err = os.WriteFile(filename, bytes, 0644)
    if err != nil {
        log.Fatalln("Failed to save users to file:", err)
    }
}
