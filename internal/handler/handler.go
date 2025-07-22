package handler

import (
	"net/http"
	"time"

	"golang-subscriptions-api/internal/model"
	"golang-subscriptions-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	svc *service.Service
	log *logrus.Logger
}

func NewHandler(svc *service.Service) *Handler {
	logger := logrus.New()
	return &Handler{svc: svc, log: logger}
}

// CreateSubscription godoc
// @Summary Создать подписку
// @Description Создает новую запись о подписке
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body model.Subscription true "Subscription info"
// @Success 201 {object} model.Subscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
func (h *Handler) Create(c *gin.Context) {
	var input model.Subscription

	// Привязка JSON к структуре
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input: " + err.Error()})
		return
	}

	// Проверка валидности UUID user_id
	if input.UserID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	// Генерируем новый UUID для новой подписки (если не задан)
	if input.ID == uuid.Nil {
		input.ID = uuid.New()
	}

	// Вызов сервиса для сохранения
	err := h.svc.Create(&input)
	if err != nil {
		h.log.Error("Failed to create subscription: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	// Возвращаем созданный объект с кодом 201 Created
	c.JSON(http.StatusCreated, input)
}

// ListSubscriptions godoc
// @Summary Получить список подписок
// @Description Возвращает список всех подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {array} model.Subscription
// @Failure 500 {object} map[string]string
// @Router /subscriptions [get]
func (h *Handler) List(c *gin.Context) {
	subs, err := h.svc.GetAll()
	if err != nil {
		h.log.Error("Failed to list subscriptions: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list subscriptions"})
		return
	}
	c.JSON(http.StatusOK, subs)
}

// GetSubscription godoc
// @Summary Получить подписку по ID
// @Description Возвращает подписку по её ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} model.Subscription
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	sub, err := h.svc.GetByID(id)
	if err != nil {
		h.log.Error("Failed to get subscription: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get subscription"})
		return
	}
	if sub == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}
	c.JSON(http.StatusOK, sub)
}

// UpdateSubscription godoc
// @Summary Обновить подписку
// @Description Обновляет существующую подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param subscription body model.Subscription true "Subscription info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	var input model.Subscription
	if err := c.ShouldBindJSON(&input); err != nil {
		h.log.Warn("Invalid update request: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Update(id, &input); err != nil {
		h.log.Error("Failed to update subscription: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription updated"})
}

// DeleteSubscription godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.Delete(id); err != nil {
		h.log.Error("Failed to delete subscription: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete subscription"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription deleted"})
}

// TotalSum godoc
// @Summary Получить общую сумму подписок по фильтрам
// @Description Возвращает сумму цен подписок с возможностью фильтрации по user_id, service_name и диапазону дат
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service Name"
// @Param from query string false "Дата начала (MM-YYYY)"
// @Param to query string false "Дата конца (MM-YYYY)"
// @Success 200 {object} map[string]int64
// @Failure 500 {object} map[string]string
// @Router /subscriptions/total [get]
func (h *Handler) TotalSum(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	fromStr := c.Query("from")
	toStr := c.Query("to")

	parseDate := func(s string) *time.Time {
		if s == "" {
			return nil
		}
		t, err := time.Parse("01-2006", s)
		if err != nil {
			return nil
		}
		return &t
	}

	from := parseDate(fromStr)
	to := parseDate(toStr)

	total, err := h.svc.TotalSum(userID, serviceName, from, to)
	if err != nil {
		h.log.Error("Failed to calculate total sum: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}
