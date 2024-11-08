package handlers

import (
	"encoding/json"
	"net/http"
	"testproj/models"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

type TestsHandler struct {
	DB            *gorm.DB
	TestsRequests *prometheus.CounterVec
}

// GET
func (h *TestsHandler) GetListTests(w http.ResponseWriter, r *http.Request) {
	h.TestsRequests.WithLabelValues("getListTests").Inc()
	var tests []models.Tests
	if err := h.DB.Find(&tests).Error; err != nil {
		http.Error(w, "Ошибка получения списка тестов", http.StatusInternalServerError)
		return
	}
	response := make([]map[string]interface{}, len(tests))
	for i, test := range tests {
		response[i] = map[string]interface{}{
			"id":   test.Id,
			"name": test.Name,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET
// ?{id=}
func (h *TestsHandler) GetTest(w http.ResponseWriter, r *http.Request) {
	h.TestsRequests.WithLabelValues("getTest").Inc()
	getTestId := r.URL.Query().Get("id")
	var tests []models.Tests
	if err := h.DB.Preload("Questions.Answers").Find(&tests, "id = ?", getTestId).Error; err != nil {
		http.Error(w, "Ошибка получения списка тестов", http.StatusInternalServerError)
		return
	}
	if len(tests) == 0 {
		http.Error(w, "Тест не найден", http.StatusNotFound)
		return
	} else {
		var firstTest = tests[0]
		json.NewEncoder(w).Encode(firstTest)
	}
}

func (h *TestsHandler) GetUnpassedTests(w http.ResponseWriter, r *http.Request) {
	h.TestsRequests.WithLabelValues("getUnpassedTests").Inc()
	userId := r.URL.Query().Get("userId")
	var unpassedTests []models.Result_Tests

	if userId == "" {
		http.Error(w, "Неуказан id пользователя", http.StatusBadRequest)
		return
	}
	if err := h.DB.Model(&models.Tests{}).
		Where("NOT EXISTS (SELECT 1 FROM test_to_users tu WHERE tu.test_id = tests.id AND tu.user_id = ? AND tu.is_passed = true)", userId).Find(&unpassedTests).Error; err != nil {
		http.Error(w, "Ошибка получения непройденных тестов", http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(unpassedTests)
	}

}

// GET
// ?{id=}
func (h *TestsHandler) GetMyTests(w http.ResponseWriter, r *http.Request) {
	h.TestsRequests.WithLabelValues("getMyTests").Inc()
	userId := r.URL.Query().Get("userId")
	var results []map[string]interface{}

	if userId == "" {
		http.Error(w, "Неуказан id пользователя", http.StatusBadRequest)
		return
	}
	err := h.DB.Table("test_to_users tu").
		Select(`tu.test_id AS "Id", t.name AS "Name", COALESCE(tu.is_passed, false) AS "IsPassed"`).
		Joins("LEFT JOIN tests t ON tu.test_id = t.id").
		Where("tu.user_id = ?", userId).
		Scan(&results).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}

// POST
func (h *TestsHandler) AddedUserTest(w http.ResponseWriter, r *http.Request) {
	h.TestsRequests.WithLabelValues("addedUserTest").Inc()
	var addedTest models.TestToUsers

	if err := json.NewDecoder(r.Body).Decode(&addedTest); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	if addedTest.TestId == uuid.Nil || addedTest.UserId == uuid.Nil {
		http.Error(w, "ID теста, ID пользователя и isPassed не могут быть пустыми", http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&addedTest).Error; err != nil {
		http.Error(w, "Ошибка добавления теста в список тестов пользователя"+err.Error(), http.StatusBadRequest)
	}
	response := map[string]interface{}{
		"status": "Тест добавлен к пользователю",
		"id":     addedTest.TestId,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
