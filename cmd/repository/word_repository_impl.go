package repository

import (
	"errors"
	"vocab_service/cmd/model"

	"gorm.io/gorm"
)

type WordRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements WordRepository.
func (w *WordRepositoryImpl) Delete(wordId uint32) {
	var word model.Word
	result := w.Db.Where("id = ?", wordId).Delete(&word)
	if result.Error != nil {
		panic(result.Error)
	}
}

// FindByUserId implements WordRepository.
func (w *WordRepositoryImpl) FindByUserId(userId uint32) ([]model.Word, error) {
	var words []model.Word
	result := w.Db.Where("user_id = ?", userId).Find(&words)
	if result != nil {
		return words, nil
	} else {
		return nil, errors.New("words is not found")
	}
}

// Save implements WordRepository.
func (w *WordRepositoryImpl) Save(word model.Word) {
	result := w.Db.Create(&word)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Update implements WordRepository.
func (w *WordRepositoryImpl) Update(word model.Word) {
	var updateWord = &model.Word{
		Definition: word.Definition,
	}

	result := w.Db.Model(&word).Where("id = ?", word.Id).Updates(updateWord)
	if result.Error != nil {
		panic(result.Error)
	}
}

// IsOwnerOfWord implements WordRepository.
func (w *WordRepositoryImpl) IsOwnerOfWord(userId uint32, wordId uint32) bool {
	var word model.Word
	result := w.Db.Where("id = ?", wordId).Find(&word)
	if result.Error != nil {
		panic(result.Error)
	}

	if word.UserId == userId {
		return true
	} else {
		return false
	}
}

func NewWordRepositoryImpl(Db *gorm.DB) WordRepository {
	return &WordRepositoryImpl{Db: Db}
}
