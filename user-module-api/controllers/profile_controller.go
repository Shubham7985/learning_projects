package controllers

import (
	"log"

	"github.com/Shubham7985/user-module-api/config"
	"github.com/Shubham7985/user-module-api/models"
	"github.com/Shubham7985/user-module-api/utils"
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
	id := c.Param("id")

	query := `
		SELECT user_id, first_name, last_name, gender, profession, profile_pic, dob
		FROM profiles
		WHERE user_id = $1
	`

	var profile models.Profile

	err := config.Supabase.QueryRow(query, id).Scan(
		&profile.UserID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Gender,
		&profile.Profession,
		&profile.ProfilePic,
		&profile.DOB,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Profile not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User profile fetched successfully",
		"data":    profile,
	})
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	query := `
		UPDATE profiles
		SET
			first_name = $1,
			last_name = $2,
			gender = $3,
			profession = $4,
			profile_pic = $5,
			dob = $6
		WHERE user_id = $7
	`

	result, err := config.Supabase.Exec(
		query,
		profile.FirstName,
		profile.LastName,
		profile.Gender,
		profile.Profession,
		profile.ProfilePic,
		profile.DOB,
		userID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to update profile",
		})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Profile not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Profile updated successfully",
	})
}
