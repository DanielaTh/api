package repositories

import (
	"github.com/DanielaTh/api/models"
)

// SubjectStore : data access subject interface
type SubjectStore interface {
	GetSubject(id int) (models.Subject, error)
	AddSubject(s models.Subject) (models.Subject, error)
}
