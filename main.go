package main

import (
	"log"
	"net/http"

	"testproj/handlers"
	"testproj/models"
	"testproj/routes"

	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_requests_total",
			Help: "Total number of requests to user handler",
		},
		[]string{"method"},
	)

	testsRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tests_requests_total",
			Help: "Total number of requests to tests handler",
		},
		[]string{"method"},
	)
)

func main() {
	// Подключение к базе данных
	dsn := "host=postgres port=5433 user=postgres dbname=testproj password=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	// Автоматическое создание таблицы
	err = db.AutoMigrate(&models.Answer{}, &models.Question{}, &models.Tests{}, &models.Users{}, &models.TestToUsers{})
	if err != nil {
		log.Fatal("Ошибка миграции: ", err)

	}
	err = db.Exec(`ALTER TABLE "questions" ADD FOREIGN KEY ("tests_id") REFERENCES "tests" ("id");`).Error
	if err != nil {
		log.Fatalf("failed to add foreign key for questions: %v", err)
	}

	err = db.Exec(`ALTER TABLE "answers" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");`).Error
	if err != nil {
		log.Fatalf("failed to add foreign key for answers: %v", err)
	}

	err = db.Exec(`ALTER TABLE "test_to_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");`).Error
	if err != nil {
		log.Fatalf("failed to add foreign key for test_to_users: %v", err)
	}

	err = db.Exec(`ALTER TABLE "test_to_users" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");`).Error
	if err != nil {
		log.Fatalf("failed to add foreign key for test_to_users: %v", err)
	}

	prometheus.MustRegister(userRequests)
	prometheus.MustRegister(testsRequests)

	userHandler := &handlers.UserHandler{
		DB:           db,
		UserRequests: userRequests,
	}
	testsHandler := &handlers.TestsHandler{
		DB:            db,
		TestsRequests: testsRequests}

	// Создаем новый маршрутизатор
	// userHandler := &handlers.UserHandler{DB: db}
	// testsHandler := &handlers.TestsHandler{DB: db}
	// userHandler := &handlers.UserHandler{DB: db}
	// testsHandler := &handlers.TestsHandler{DB: db}
	// Определяем маршрут для создания пользователя
	router := routes.InitializeRoutes(userHandler, testsHandler)

	// http.Handle("/metrics", promhttp.Handler())

	// Запуск HTTP-сервера
	log.Println("Сервер запущен на :8000")
	http.ListenAndServe(":8000", router)
}
