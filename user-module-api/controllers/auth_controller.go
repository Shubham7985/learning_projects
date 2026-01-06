package controllers

import (
	"user-module-api/config"
	"user-module-api/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Mobile   string `json:"mobile" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.Error(c, 500, "Password hashing failed")
		return
	}

	_, err = config.Supabase.Exec(
		`INSERT INTO users (email, mobile, password)
		 VALUES ($1, $2, $3)`,
		req.Email,
		req.Mobile,
		hashedPassword,
	)

	if err != nil {
		utils.Error(c, 500, "User creation failed")
		return
	}

	utils.Success(c, "User registered successfully")
}

func SignIn(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	var user struct {
		ID       string
		Password string
	}

	err := config.Supabase.QueryRow(
		"SELECT id, password FROM users WHERE email=$1",
		req.Email,
	).Scan(&user.ID, &user.Password)

	if err != nil {
		utils.Error(c, 401, "User not found")
		return
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		utils.Error(c, 401, "Incorrect password")
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	utils.Success(c, "Login success", gin.H{
		"token": token,
	})
}

func ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	// 🔹 Demo: normally yahan email send hoga
	utils.Success(c, "Reset link sent to "+req.Email)
}
