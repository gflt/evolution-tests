package migration

const migration = `
CREATE TABLE "Tests" (
	"Id" uuid PRIMARY KEY NOT NULL,
	"Name" TEXT NOT NULL
  );
  
  CREATE TABLE "Questions" (
	"Id" uuid PRIMARY KEY NOT NULL,
	"Name" TEXT NOT NULL,
	"TestId" uuid NOT NULL
  );
  
  CREATE TABLE "Answers" (
	"Id" uuid PRIMARY KEY NOT NULL,
	"Name" TEXT NOT NULL,
	"IsCorrect" bool NOT NULL,
	"QuestionId" uuid NOT NULL
  );
  
  CREATE TABLE "Users" (
	"Id" uuid PRIMARY KEY NOT NULL,
	"Nickname" TEXT UNIQUE NOT NULL,
	"Password" TEXT NOT NULL
  );
  
  CREATE TABLE "TestsToUsers" (
	"UserId" uuid NOT NULL,
	"TestId" uuid NOT NULL,
	"IsPassed" boolean NOT NULL
  );
  
  COMMENT ON TABLE "Tests" IS 'Таблица для информации о тестах';
  
  COMMENT ON TABLE "Questions" IS 'Таблица для информации о вопросах';
  
  COMMENT ON TABLE "Answers" IS 'Таблица для информации о ответах';
  
  COMMENT ON TABLE "Users" IS 'Таблица для информации о пользователях';
  
  COMMENT ON TABLE "TestsToUsers" IS 'Таблица для информации о всех тестах пользователя';
  
  ALTER TABLE "Questions" ADD FOREIGN KEY ("TestId") REFERENCES "Tests" ("Id");
  
  ALTER TABLE "Answers" ADD FOREIGN KEY ("QuestionId") REFERENCES "Questions" ("Id");
  
  ALTER TABLE "TestsToUsers" ADD FOREIGN KEY ("UserId") REFERENCES "Users" ("Id");
  
  ALTER TABLE "TestsToUsers" ADD FOREIGN KEY ("TestId") REFERENCES "Tests" ("Id");
`

func GetMigration() string {
	return migration
}
