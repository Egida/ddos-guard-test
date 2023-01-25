# ddos-guard-test

## Немного про архитектуру:
За референс брал https://github.com/evrone/go-clean-template/

### cmd/bot cmd/http
Точки входа в аппы

### config/
Инициализация конфига

### docker/
Докерфайлы для проекта

### internal/app
Модули запуска каждого приложения

### controller/
Получение/обработка запросов в http и tg

### entity
Модели данных

### infrastructure
Слой с логикой взаимодействия с бд, сторонними библиотеками и тд.

### infrastructure/mathservice
Сервис получения, обработки и отдачи ответов

### infrastructure/repository
Репозиторий. Работа с БД, в моем случае с PostgreSQL

### usecase
Слой юзкейсов