package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"vocab_service/cmd/config"
	"vocab_service/cmd/model"
	"vocab_service/cmd/repository"
)

func TestWordRepository(t *testing.T) {
	lc, err := config.LoadConfig("..")
	if err != nil {
		t.Fatalf("🚀 Could not load environment variables: %v", err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		lc.DBHost, lc.DBPort, lc.DBUsername, lc.DBPassword, lc.DBTestName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to connect to the database:", err)
	}
	db.AutoMigrate(&model.Word{})

	wr := repository.NewWordRepositoryImpl(db)

	testWord := model.Word{
		Id:              1,
		Word:            "water",
		Definition:      "вода",
		UserId:          1,
		CreatedAt:       time.Now(),

		IsLearned:       false,
		Cards:           false,
		WordTranslation: false,
		Constructor:     false,
		WordAudio:       false,
	}

	t.Run("Test Save Word", func(t *testing.T) {
		err := wr.Save(testWord)
		assert.NoError(t, err, "Expected no error while saving the word")

	})

	t.Run("Test Update Word", func(t *testing.T) {
		updatedWord := model.Word{
			Id:              testWord.Id,
			Definition:      "водичка",
			UserId:          testWord.UserId,
		}
		err := wr.Update(updatedWord)
		assert.NoError(t, err, "Expected no error while updating the word")

		words, _ := wr.FindByUserId(1)
		assert.Equal(t, words[0].Definition, "водичка", "Expected to be \"водичка\"")
	})

	t.Run("Test ManageTrainings Word", func(t *testing.T) {
		err := wr.ManageTrainings(true, "cards", testWord.Id)
		assert.NoError(t, err, "Expected no error while manage cards")

		err = wr.ManageTrainings(true, "word_translation", testWord.Id)
		assert.NoError(t, err, "Expected no error while manage word_translation")

		err = wr.ManageTrainings(true, "constructor", testWord.Id)
		assert.NoError(t, err, "Expected no error while manage constructor")

		words, _ := wr.FindByUserId(1)
		assert.False(t, words[0].IsLearned, "should be not learned yet")

		err = wr.ManageTrainings(true, "word_audio", testWord.Id)
		assert.NoError(t, err, "Expected no error while manage")

		err = wr.ManageTrainings(true, "eavcerwrfc", testWord.Id)
		assert.Error(t, err, "Expected error while manage unknown training")

		words, _ = wr.FindByUserId(1)
		assert.True(t, words[0].IsLearned, "should be already learned")
	})

	t.Run("Test IsOwnerOfWord Word", func(t *testing.T) {
		assert.True(t, wr.IsOwnerOfWord(testWord.UserId, testWord.Id), "Expected to be the owner of the word")
		assert.False(t, wr.IsOwnerOfWord(5, testWord.Id), "Expected to not be the owner of the word")
	})

	t.Run("Test FindByUserId Word", func(t *testing.T) {
		words, err := wr.FindByUserId(1)
		assert.NoError(t, err, "Expected no error while finding words by id")
		assert.NotEmpty(t, words, "Should return words")

		words, err = wr.FindByUserId(6)
		assert.NoError(t, err, "Expected no error while finding words by id")
		assert.Equal(t, words, []model.Word{}, "Should not return words")
	})

	t.Run("Test Delete Word", func(t *testing.T) {
		wr.Delete(testWord.Id)
		words, err := wr.FindByUserId(testWord.UserId)
		assert.Equal(t, words, []model.Word{}, "Should not return words")
		assert.NoError(t, err, "Expected no error while finding words by id")
	})

	db.Migrator().DropTable(&model.Word{})
}
