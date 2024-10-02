package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nilpntr/gitdesk-forwarder/internal/config"
	"github.com/nilpntr/gitdesk-forwarder/internal/messengers"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type WebhookRequest struct {
	User struct {
		Username string `json:"username"`
	} `json:"user"`
	ObjectAttributes ObjectAttributes  `json:"object_attributes"`
	Changes          map[string]Change `json:"changes"`
}

type ObjectAttributes struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Change struct {
	Previous any `json:"previous"`
	Current  any `json:"current"`
}

func HandleWebhook(c *gin.Context, webhook config.Webhook) {
	var dto WebhookRequest
	if err := c.ShouldBind(&dto); err != nil {
		zap.L().Sugar().Debugf("failed to bind webhook request: %v", err)
		c.AbortWithStatus(400)
		return
	}

	// Check if we need to validate the provider secret token against header X-Gitlab-Token
	if len(webhook.SecretToken) > 0 {
		if webhook.SecretToken != c.GetHeader("X-Gitlab-Token") {
			zap.L().Sugar().Debugf("webhook secret token mismatch")
			c.AbortWithStatus(500)
			return
		}
	}

	// Determine if the issue is created by the support bot
	if dto.User.Username != viper.Get("botUsername") {
		zap.L().Sugar().Debug("username mismatch")
		c.Status(200)
		return
	}

	// Determine if the issue is new
	if !isNew(dto.Changes) {
		zap.L().Sugar().Debug("received webhook is not a new issue")
		c.Status(200)
		return
	}

	if webhook.SlackWebhookUrl != nil {
		if slack, ok := messengers.Messengers["slack"]; ok {
			if err := slack.SendMessage(*webhook.SlackWebhookUrl, dto.ObjectAttributes.Title, dto.ObjectAttributes.Description, dto.ObjectAttributes.Url); err != nil {
				zap.L().Sugar().Errorf("failed to send slack message: %v", err)
				c.Status(500)
				return
			}
		}
	}

	c.Status(200)
}

func isNew(changes map[string]Change) bool {
	idChange, idExists := changes["id"]
	createdAtChange, createdAtExists := changes["created_at"]

	return idExists && createdAtExists && idChange.Previous == nil && createdAtChange.Previous == nil
}
