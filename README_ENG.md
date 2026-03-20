# Shortli — URL Shortener Service

Shortli is a lightweight URL shortening service.
It lets you create short links and retrieve original URLs, focusing on simplicity and speed.
Data is stored locally using SQLite.
Automated builds are configured via GitHub Actions.

---

## Features

- Shorten long URLs to compact identifiers
- Retrieve the original URL by its alias and redirect
- Input validation
- Local link storage with SQLite
- Automated builds via GitHub Actions

---

## Tech Stack

- **Golang** — primary language
- **SQLite** — local data storage
- **net/http**, **go-chi** — REST API implementation
- **GitHub Actions** — CI/CD for automated builds

---

## Installation & Setup

### 1. Clone the repository
```bash
git clone https://github.com/Deymos01/shortli.git
cd shortli
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Run the application
```bash
CONFIG_PATH=config/local.yaml go run ./cmd/shortli/
```

By default, the server starts at `http://localhost:8082`.

---

## API Examples

### Create a short link

***POST*** `/url`

Authentication — Basic Auth (credentials from `./config/local.yaml`)

#### Request:
```bash
curl -X POST http://localhost:8082/url \
  -u myuser:mypass \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/"
  }'
```

#### Response:
```bash
{"status":"OK","alias":"4ckqm0"}
```

---

### Follow a short link

***GET*** `/{alias}`

#### Request:
```bash
curl -v http://localhost:8082/4ckqm0
```

#### Response:
```bash
> GET /4ckqm0 HTTP/1.1
...
< HTTP/1.1 302 Found
...
<a href="https://example.com">Found</a>.
```

---

### Delete a short link

***DELETE*** `/url/{alias}`

Authentication — Basic Auth (credentials from `./config/local.yaml`)

#### Request:
```bash
curl -X DELETE http://localhost:8082/url/4ckqm0 \
    -u myuser:mypass
```

#### Response:
```bash
{"status":"OK"}
```