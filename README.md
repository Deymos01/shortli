# Shortli — URL Shortener Service

Shortli — это лёгкий сервис для сокращения ссылок.  
Проект позволяет создавать короткие ссылки и получать оригинальные URL, обеспечивая простоту и скорость работы.  
Хранение данных реализовано с использованием локальной базы SQLite.  
Настроена автоматическая сборка через GitHub Actions.

---

## Возможности

- Сокращение длинных URL до коротких идентификаторов
- Получение оригинальной ссылки по сокращённому ID и дальнейшая переадресация
- Проверка корректности введённых данных
- Локальное хранение ссылок в SQLite
- Автоматическая сборка через GitHub Actions

---

## Технологии

- **Golang** — язык разработки
- **SQLite** — локальное хранилище данных
- **net/http**, **go-chi** — для реализации REST API
- **GitHub Actions** — CI/CD для автоматической сборки приложения

---

## Установка и запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/Deymos01/shortli.git
cd shortli
```

### 2. Установить зависимости

```bash
go mod tidy
```

### 3. Запустить приложение

```bash
CONFIG_PATH=config/local.yaml go run ./cmd/shortli/
```

По умолчанию сервер запустится на ```http://localhost:8082```.

---

## Примеры запросов

### Создание короткой ссылки

***POST*** ```/url```

Аутентификация - Basic Auth (данные из файла конфигурации ./config/local.yaml)

#### Пример запроса:

```bash
curl -X POST http://localhost:8082/url \
  -u myuser:mypass \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/"
  }'
```

#### Пример ответа:

```bash
{"status":"OK","alias":"4ckqm0"}
```

### Переход по короткой ссылке

***GET*** ```/{alias}```

#### Пример запроса:

```bash
curl -v http://localhost:8082/4ckqm0
```

#### Пример ответа:

```bash
> GET /4ckqm0 HTTP/1.1
...
< HTTP/1.1 302 Found
...
<a href="https://example.com">Found</a>.
```

### Удаление короткой ссылки

***DELETE*** ```/url/{alias}```

Аутентификация — Basic Auth (данные из файла конфигурации ./config/local.yaml)

#### Пример запроса:

```bash
curl -X DELETE http://localhost:8082/url/4ckqm0 \
    -u myuser:mypass
```

#### Пример ответа:

```bash
{"status":"OK"}
```
