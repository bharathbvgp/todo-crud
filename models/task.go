package models

import (
	"gorm.io/gorm"
	"time"
	"todoapp/enums"
)



type Task struct {
    gorm.Model  
    Title       string         `json:"title"`
    Description string         `json:"description"`
    DueDate     time.Time      `json:"due_date"`
    Priority    enums.PriorityType `json:"priority"`
    Status      enums.StatusType   `json:"status"`
    Category    enums.CategoryType `json:"category"`
}
