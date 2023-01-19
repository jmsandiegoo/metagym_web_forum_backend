package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateNewThread(thread *databasemodels.Thread) (*databasemodels.Thread, error) {
	// Generate ID
	thread.ID = api.GenerateUUID()

	// store to database
	err := database.Database.Create(&thread).Error
	if err != nil {
		return &(databasemodels.Thread{}), err
	}
	return thread, nil
}

func FindThreadById(id uuid.UUID) (databasemodels.Thread, error) {
	var thread databasemodels.Thread
	err := database.Database.Preload("Comments").Preload("UsersLiked").Preload("UsersDisliked").Preload("Interests").Where("id=?", id).Find(&thread).Error

	if err != nil {
		return databasemodels.Thread{}, err
	}

	return thread, nil
}

func FindThreadByIdLocked(id uuid.UUID, tx *gorm.DB) (databasemodels.Thread, error) {
	var thread databasemodels.Thread // garbage collected once no reference
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id=?", id).Find(&thread).Error
	if err != nil {
		return databasemodels.Thread{}, err
	}

	return thread, nil
}

func FindThreadByInterestAndTitle(interest_ids []uuid.UUID, title string) ([]databasemodels.Thread, error) {
	chain := database.Database.Debug()
	// Join // add select
	if len(interest_ids) > 0 {
		chain = chain.Select("*").Joins("inner join thread_interests ti on ti.thread_id = threads.id").Where("ti.interest_id IN ?", interest_ids)
	}

	// title keyword
	if title != "" {
		chain = chain.Where("title LIKE ?", "%"+title+"%")
	}

	// find the specific thread
	var threads []databasemodels.Thread
	err := chain.Order("threads.created_at desc").Preload("Interests").Find(&threads).Error

	if err != nil {
		return nil, err
	}

	return threads, nil
}

func UpdateThread(thread *databasemodels.Thread) (*databasemodels.Thread, error) {
	err := database.Database.Save(&thread).Error

	if err != nil {
		return &(databasemodels.Thread{}), err
	}

	return thread, nil
}

func DeleteThread(thread *databasemodels.Thread) error {
	tx := database.Database.Begin()

	// Delete associations
	err := tx.Where("thread_id = ?", thread.ID).Delete(&databasemodels.PostLike{}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("thread_id = ?", thread.ID).Delete(&databasemodels.PostDislike{}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("thread_id = ?", thread.ID).Delete(&databasemodels.ThreadInterest{}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("thread_id = ?", thread.ID).Delete(&databasemodels.Comment{}).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete the thread record
	err = tx.Delete(&thread).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error

	if err != nil {
		return err
	}

	return nil
}

// Vote related functions requires a gorm transaction variable

func AddUsersLikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersLiked").Append(user)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsersLikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersLiked").Delete(user)

	if err != nil {
		return err
	}

	return nil
}

func AddUsersDislikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersDisliked").Append(user)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsersDislikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersDisliked").Delete(user)

	if err != nil {
		return err
	}

	return nil
}

func FindThreadUsersLikedByIdsLocked(thread *databasemodels.Thread, ids []uuid.UUID, tx *gorm.DB) ([]databasemodels.User, error) {
	var users []databasemodels.User
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&thread).Where(ids).Association("UsersLiked").Find(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindThreadUsersDislikedByIdsLocked(thread *databasemodels.Thread, ids []uuid.UUID, tx *gorm.DB) ([]databasemodels.User, error) {
	var users []databasemodels.User
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&thread).Where(ids).Association("UsersDisliked").Find(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
