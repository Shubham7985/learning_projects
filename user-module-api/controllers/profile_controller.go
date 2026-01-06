package controllers

import (
	"log"
	"user-module-api/config"
	"user-module-api/models"
	"user-module-api/utils"

	"github.com/gin-gonic/gin"
)

// CreateProfile - insert new profile
func CreateProfile(c *gin.Context) {
	var profile models.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		utils.Error(c, 400, "Invalid request body: "+err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, "Unauthorized: user_id not found")
		return
	}
	profile.UserID = userID.(string)

	query := `
	INSERT INTO profiles (user_id, first_name, last_name, gender, profession, profile_pic, dob)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := config.Supabase.Exec(
		query,
		profile.UserID,
		profile.FirstName,
		profile.LastName,
		profile.Gender,
		profile.Profession,
		profile.ProfilePic,
		profile.DOB,
	)

	if err != nil {
		log.Println("Profile insert error:", err)
		utils.Error(c, 500, "Profile creation failed: "+err.Error())
		return
	}

	utils.Success(c, "Profile created successfully", profile)
}

// ListUsers - list all profiles
func ListUsers(c *gin.Context) {
	rows, err := config.Supabase.Query(`SELECT user_id, first_name, last_name, gender, profession, profile_pic, dob FROM profiles`)
	if err != nil {
		utils.Error(c, 500, "Failed to fetch users: "+err.Error())
		return
	}
	defer rows.Close()

	var users []models.Profile
	for rows.Next() {
		var u models.Profile
		if err := rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Gender, &u.Profession, &u.ProfilePic, &u.DOB); err != nil {
			utils.Error(c, 500, "Row scan failed: "+err.Error())
			return
		}
		users = append(users, u)
	}

	utils.Success(c, "User list fetched successfully", users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id") // URL parameter

	if id == "" {
		utils.Error(c, 400, "User ID is required")
		return
	}

	// Baad me yahan Supabase se user fetch hoga
	utils.Success(c, "User fetched successfully", gin.H{
		"user_id": id,
	})

}
