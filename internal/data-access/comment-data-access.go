package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateNewComment(comment *databasemodels.Comment) (*databasemodels.Comment, error) {
	// Generate ID
	comment.ID = api.GenerateUUID()

	// store to database
	err := database.Database.Create(&comment).Error

	if err != nil {
		return &(databasemodels.Comment{}), err
	}

	return comment, nil
}

func FindCommentById(id uuid.UUID) (databasemodels.Comment, error) {
	var comment databasemodels.Comment // garbage collected once no reference
	err := database.Database.Preload("Thread").Where("id=?", id).Find(&comment).Error

	if err != nil {
		return databasemodels.Comment{}, err
	}

	return comment, nil
}

func FindCommentByIdLocked(id uuid.UUID, tx *gorm.DB) (databasemodels.Comment, error) {
	var comment databasemodels.Comment // garbage collected once no reference
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id=?", id).Find(&comment).Error

	if err != nil {
		return databasemodels.Comment{}, err
	}

	return comment, nil
}

func UpdateComment(comment *databasemodels.Comment) (*databasemodels.Comment, error) {
	err := database.Database.Save(&comment).Error

	if err != nil {
		return &(databasemodels.Comment{}), err
	}

	return comment, nil
}

func DeleteComment(comment *databasemodels.Comment) error {
	err := database.Database.Delete(&comment).Error

	if err != nil {
		return err
	}

	return nil
}

func AddUsersLikedComment(comment *databasemodels.Comment, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&comment).Association("UsersLiked").Append(user)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsersLikedComment(comment *databasemodels.Comment, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&comment).Association("UsersLiked").Delete(user)

	if err != nil {
		return err
	}

	return nil
}

func AddUsersDislikedComment(comment *databasemodels.Comment, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&comment).Association("UsersDisliked").Append(user)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsersDislikedComment(comment *databasemodels.Comment, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&comment).Association("UsersDisliked").Delete(user)

	if err != nil {
		return err
	}

	return nil
}

func FindCommentUsersLikedByIdsLocked(comment *databasemodels.Comment, ids []uuid.UUID, tx *gorm.DB) ([]databasemodels.User, error) {
	var users []databasemodels.User
	// locks row until tx is committed or rolledback
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&comment).Where(ids).Association("UsersLiked").Find(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindCommentUsersDislikedByIdsLocked(comment *databasemodels.Comment, ids []uuid.UUID, tx *gorm.DB) ([]databasemodels.User, error) {
	var users []databasemodels.User
	// locks row until tx is committed or rolledback
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&comment).Where(ids).Association("UsersDisliked").Find(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
