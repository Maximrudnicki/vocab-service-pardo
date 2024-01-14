package repository

import "vocab_service/cmd/model"

type WordRepository interface {
	Save(word model.Word)
	Update(word model.Word)
	Delete(wordId uint32)
	FindByUserId(userId uint32) ([]model.Word, error)
	ManageTrainings(res bool, training string, wordId uint32) error

	// utils
	IsOwnerOfWord(userId uint32, wordId uint32) bool
}
