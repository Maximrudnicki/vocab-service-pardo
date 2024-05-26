package repository

import "vocab_service/cmd/model"

type WordRepository interface {
	Add(word model.Word) (uint32, error)
	Save(word model.Word) error
	Update(word model.Word) error
	UpdateStatus(word model.Word) error
	Delete(wordId uint32)
	FindByUserId(userId uint32) ([]model.Word, error)
	FindById(wordId uint32) (model.Word, error)
	ManageTrainings(res bool, training string, wordId uint32) error

	// utils
	IsOwnerOfWord(userId uint32, wordId uint32) bool
}
