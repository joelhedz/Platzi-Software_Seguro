package main

import (
	"context"
	"encoding/json"
	"github-tracker/github-tracker/models"
	reporsitory "github-tracker/github-tracker/repository"
	"github-tracker/github-tracker/repository/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDummy(t *testing.T) {
	c := require.New(t)

	result := 22

	c.Equal(22, result)
}

func TestInsert(t *testing.T) {
	c := require.New(t)

	webhook := models.GitHubWebhook{
		Repository: models.Repository{
			FullName: "joelhedz/Platzi-Software_Seguro",
		},
		HeadCommit: models.Commit{
			ID:      "5537ee78e64bce92928ce932fd123683f04a85d3",
			Message: "Fix Checkout version on github action",
			Author: models.CommitUser{
				Email:    "joelhumberto334@gmail.com",
				Username: "joelhedz",
			},
		},
	}

	body, err := json.Marshal(webhook)
	if err != nil {
		c.NoError(err)
	}

	cretedTime := time.Now()

	m := mock.Mock{}

	mockCommit := reporsitory.MockCommit{Mock: &m}

	commit := entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitId:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        string(body),
		CreatedAt:      cretedTime,
		UpdatedAt:      cretedTime,
	}

	ctx := context.Background()

	mockCommit.On("Insert", ctx, &commit).Return(nil)

	err = insertGitHubWebhook(ctx, mockCommit, webhook, string(body), cretedTime)
	c.NoError(err)

	// Verify that the mock expectations were met
	m.AssertExpectations(t)
}
