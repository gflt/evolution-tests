package handlers

import (
	"encoding/json"
	"net/http"
	"testproj/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Декодирование JSON из тела запроса
	var user models.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	if user.Nickname == "" || user.Password == "" {
		http.Error(w, "Имя пользователя и пароль не могут быть пустыми", http.StatusBadRequest)
		return
	}

	user.Id = uuid.New()

	if err := h.DB.Create(&user).Error; err != nil {
		http.Error(w, "Ошибка создания пользователя: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var userChangePassword models.UpdatePassword

	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&userChangePassword); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// Проверяем, что nickname и пароль предоставлены
	if userChangePassword.Nickname == "" || userChangePassword.Password == "" {
		http.Error(w, "Nickname и password должны быть указаны", http.StatusBadRequest)
		return
	}

	// Поиск пользователя по nickname
	var user models.Users
	if err := h.DB.Where("nickname = ?", userChangePassword.Nickname).First(&user).Error; err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	user.Password = userChangePassword.Password

	// Сохраняем обновленного пользователя в базе данных
	if err := h.DB.Save(&user).Error; err != nil {
		http.Error(w, "Ошибка при обновлении пользователя: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пароль успешно обновлен"))
}

func (h *UserHandler) AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	var Authorize models.Authorize

	if err := json.NewDecoder(r.Body).Decode(&Authorize); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	if Authorize.Nickname == "" || Authorize.Password == "" {
		http.Error(w, "Имя пользователя и пароль должны быть указаны", http.StatusBadRequest)
		return
	}
	var user models.Users
	if err := h.DB.Where("nickname = ? AND password = ?", Authorize.Nickname, Authorize.Password).First(&user).Error; err != nil {
		http.Error(w, "Ошибка авторизации", http.StatusUnauthorized)
	} else {
		response := map[string]interface{}{
			"status": "Авторизован",
			"id":     user.Id,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
