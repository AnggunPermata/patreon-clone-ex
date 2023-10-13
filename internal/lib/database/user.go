package database

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrUpdateOneUser(c echo.Context, DB *gorm.DB, user models.User) (*models.User, error) {
	// create if record is not exists
	if DB.Model(&user).Where("username=?", user.Username).Updates(&user).RowsAffected == 0 {
		DB.Create(&user)
	}

	return &user, nil
}

// Get only one User by username
func GetOneUserByUsername(c echo.Context, DB *gorm.DB, username string) (models.User, error) {
	var user models.User

	if err := DB.Where("username=?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmailAndPassword(c echo.Context, cfg *configs.Config, DB *gorm.DB, email string, password string) (models.User, error) {
	var user models.User
	var err error

	if err = DB.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return user, fmt.Errorf("failed to sign user with email=%s, error_message=%v", email, err)
	}

	return user, nil
}

// // Get Multiple agent by IDs (multichannel_id), ordered by total_customer_daily and all_time_total_customer
// func GetMultipleAgentByIdsOrderedByTC(AgentIds []string) ([]models.Agent, error) {
// 	var agent []models.Agent

// 	if err := config.DB.Order("total_customer_daily asc, all_time_total_customer asc").Where("multichannel_id IN ?", AgentIds).Find(&agent); err != nil {
// 		return agent, err.Error
// 	}
// 	return agent, nil
// }

// // Get All agent by role (as agent), ordered by total_customer_daily and all_time_total_customer
// func GetAllAgentByRoleOrderedByTC(role string) ([]models.Agent, error) {
// 	var agent []models.Agent

// 	if err := config.DB.Order("total_customer_daily asc, all_time_total_customer asc").Where("role=?", role).Find(&agent); err != nil {
// 		return agent, err.Error
// 	}
// 	return agent, nil
// }

// Update Total Customer by 1
// func UpdateTotalCustomer(agentID string) (models.Agent, error) {
// 	var agent models.Agent
// 	config.DB.Find(&agent, "multichannel_id=?", agentID)
// 	agent.TotalCustomerDaily += 1
// 	agent.AllTimeTotalCustomer += 1

// 	updatedAgent, err := CreateOrUpdateOneAgent(agent)
// 	if err != nil {
// 		return *updatedAgent, err
// 	}
// 	return *updatedAgent, nil
// }

// // Delete one agent by multichannel id on database
// // Currently Unused
// func DeleteOneAgent(agentID string) (models.Agent, error) {
// 	var deleteAnAgent models.Agent
// 	if err := config.DB.Find(&deleteAnAgent, "multichannel_id=?", agentID).Error; err != nil {
// 		return deleteAnAgent, err
// 	}

// 	if err := config.DB.Delete(&deleteAnAgent).Error; err != nil {
// 		return deleteAnAgent, err
// 	}
// 	return deleteAnAgent, nil
// }
