# evolution-tests
Для запуска требуется запущенная БД postgres на 5433 порту

Команда запуска go run main.go

Сервер запускает на 8000 порту

Описание api эндпоинтов
**/createuser Methods("POST")**
Создание пользователя, на вход принимает raw body, например:
{
    "Nickname": "testUser123",
    "Password": "111222333"
}

Ответ 200:
{"Id":"some uuid...","Nickname":"testUser1234","Password":"supersecurepassword"


**/authorize Methods("GET")**
Авторизация пользователя на вход принимает raw body, например:
{
    "Nickname": "testUser123",
    "Password": "111222333"
}

Ответ 200:
{"id":"some uuid...","status":"Авторизован"}**


**/unpassedtest?userId Methods("GET")**
Получение списка всех тестов пользователя со статусом is_passed=false
На вход принимает user_id, в качестве параметра

Ответ 200:

[{"Id":"f235a434-6a3c-46f3-9327-b84c1846ea99","Name":"Кто ты из Ворониных?","Questions":[]},{"Id":"9926ab11-2bf3-4731-b04b-31118be156e6","Name":"ХТО ТЫ?","Questions":[{"Id":"550e8400-e29b-41d4-a716-446655440000","Name":"Какой ваш любимый цвет?","TestsId":"9926ab11-2bf3-4731-b04b-31118be156e6","Answers":[{"Id":"c7eeb0a4-6965-4ac0-87ed-9947590a38e2","Name":"Да","IsCorrect":true,"QuestionId":"550e8400-e29b-41d4-a716-446655440000"},{"Id":"557d1b9e-ffa0-4b0a-99e4-602a990e1f2d","Name":"Нет","IsCorrect":false,"QuestionId":"550e8400-e29b-41d4-a716-446655440000"}]},{"Id":"550e8400-e29b-41d4-a716-446655440001","Name":"1+1?","TestsId":"9926ab11-2bf3-4731-b04b-31118be156e6","Answers":[{"Id":"853b585d-23a9-49a0-a5dc-0e46f8e18e8c","Name":"Да","IsCorrect":false,"QuestionId":"550e8400-e29b-41d4-a716-446655440001"},{"Id":"54a0e0df-60cb-47f0-a9f7-b2e2ae459937","Name":"Нет","IsCorrect":true,"QuestionId":"550e8400-e29b-41d4-a716-446655440001"}]},{"Id":"550e8400-e29b-41d4-a716-446655440002","Name":"Выиграет ли Ростислав на спектре?","TestsId":"9926ab11-2bf3-4731-b04b-31118be156e6","Answers":[{"Id":"dde0754b-0856-4368-98b9-baba74749864","Name":"Нет","IsCorrect":true,"QuestionId":"550e8400-e29b-41d4-a716-446655440002"},{"Id":"a3b8d48b-342a-4b72-8a4d-4f503c2440ee","Name":"Да","IsCorrect":false,"QuestionId":"550e8400-e29b-41d4-a716-446655440002"}]}]},{"Id":"32438c43-a727-452a-98dd-330ec6751a0b","Name":"ХТО Я?","Questions":[]},{"Id":"95465742-c58d-4846-8598-30c1af1c97a5","Name":"Как правильно керить на Луне","Questions":[]},{"Id":"29b85601-f563-49e3-b8ef-83660bc1f3d2","Name":"Как правильно керить на Троле","Questions":[]}]


**/getmytests?userId Methods("GET")**
Получение тестов из таблицы test_to_users (то есть тесты привязанные к пользователю)
На вход принимает user_id, в качестве параметра

Ответ 200:

[{"is_passed":true,"test_id":"9926ab11-2bf3-4731-b04b-31118be156e6","test_name":"ХТО ТЫ?"}]


**/addedtesttouser Methods("POST")**
Привязка теста к конкретному пользователю, на вход принимает raw body, например:
{
    "TestId" : "id_теста",
    "UserId" : "id_пользователя",
    "IsPassed" : false
}

 Ответ 200:
 
 {"id":"9926ab11-2bf3-4731-b04b-31118be156e6","status":"Тест добавлен к пользователю"}

**/gettest?id Methods("GET")**
Получение списка вопросов/ответов для конкретного теста
На вход принимает id, в качестве параметра

Ответ 200:

{"Id":"9926ab11-2bf3-4731-b04b-31118be156e6","Name":"ХТО ТЫ?","Questions":[{"Id":"550e8400-e29b-41d4-a716-446655440000","Name":"Какой ваш любимый цвет?","TestsId":"9926ab11-2bf3-4731-b04b-31118be156e6","Answers":[{"Id":"c7eeb0a4-6965-4ac0-87ed-9947590a38e2","Name":"Да","IsCorrect":true,"QuestionId":"550e8400-e29b-41d4-a716-446655440000"},{"Id":"557d1b9e-ffa0-4b0a-99e4-602a990e1f2d","Name":"Нет","IsCorrect":false,"QuestionId":"550e8400-e29b-41d4-a716-446655440000"}]},{"Id":"550e8400-e29b-41d4-a716-446655440001","Name":"1+1?","TestsId":"9926ab11-2bf3-4731-b04b-31118be156e6","Answers":[{"Id":"853b585d-23a9-49a0-a5dc-0e46f8e18e8c","Name":"Да","IsCorrect":false,"QuestionId":"550e8400-e29b-41d4-a716-446655440001"},{"Id":"54a0e0df-60cb-47f0-a9f7-b2e2ae459937","Name":"Нет","IsCorrect":true,"QuestionId":"550e8400-e29b-41d4-a716-446655440001"}]},{"Id":"550e8400-e29b-41d4-a716-446655440002","Name":"Выиграет ли Ростислав на спектре?","TestsId":"9926ab11-2bf3-4731-b04b-31118be156e6","Answers":[{"Id":"dde0754b-0856-4368-98b9-baba74749864","Name":"Нет","IsCorrect":true,"QuestionId":"550e8400-e29b-41d4-a716-446655440002"},{"Id":"a3b8d48b-342a-4b72-8a4d-4f503c2440ee","Name":"Да","IsCorrect":false,"QuestionId":"550e8400-e29b-41d4-a716-446655440002"}]}]}

