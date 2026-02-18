package repositories

import (
	"github.com/ferryku8/project-management/config"
	"github.com/ferryku8/project-management/models"
)

type BoardMemberRepository interface {
	GetMembers(boardPublicID string) ([]models.User, error)
}

type boardMemberRepository struct {
}

func NewBoardMemberRepository() BoardMemberRepository {
	return &boardMemberRepository{}
}

func (r *boardMemberRepository) GetMembers(boardPublicID string) ([]models.User, error) {
	var users []models.User
	err := config.DB.Joins("Join board_members ON board_members.user_internal_id = users.interanl_id").
		Joins("Join boards ON boards.internal_id = board_members.board_internal_id").
		Where("boards.public_id = ?", boardPublicID).
		Find(&users).Error
	return users, err
}
