package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kajtekajtek/forum/backend/internal/sse"
	"github.com/kajtekajtek/forum/backend/internal/utils"
)

/*
	SSE connection handler
*/
func StreamMessages(manager *sse.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get channel ID from URL parameters
		channelID, err := utils.ParseUintParam(c, "channelID")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid channel ID"})
			return
		}

		// subscribe the channel
		msgCh := manager.Subscribe(channelID)
		defer manager.Unsubscribe(channelID, msgCh)

		// set required headers
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Writer.Flush()

		/* 
			notify channel recieves a true value 
			when the client connection is gone
		*/
		notify := c.Writer.CloseNotify()

		for {
			select {
			// send SSEvent to client on message recieval
			case msg := <-msgCh:
				c.SSEvent("message", msg)
				c.Writer.Flush()
			// return on client's connection closure
			case <-notify:
				return
			// return on server's connection closure
			case <-c.Request.Context().Done():
				return
			}
		}
	}
}
