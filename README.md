# LLM Voice Agent API

A lightweight Go API that provides a simple HTTP interface to OpenAI's language models.  
Send a POST request with a `message` to the `/query` endpoint and get a response from the OpenAI model.

---

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoint](#api-endpoint)
- [Testing](#testing)
- [License](#license)

---

## Features

- HTTP server written in Go.
- `/query` endpoint for sending messages to OpenAI.
- JSON request/response format.
- Easily configurable via environment variables.
- Simple and clean codebase with modular design and doc blocks.
- Automated tests with `httptest`.

---

## Requirements

- Go 1.25 or higher
- OpenAI API key

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/jm-borges/llm-voice-agent-api.git
cd llm-voice-agent-api
````

2. Install dependencies:

```bash
go mod tidy
```

---

## Configuration

Copy the example environment file and set your OpenAI API key:

```bash
cp .env.example .env
```

Edit `.env`:

```
PORT=8090
OPENAI_API_KEY=sk-your-key-here
OPENAI_MODEL=gpt-5-nano
```

* `PORT` – Port for the HTTP server (default: 8080)
* `OPENAI_API_KEY` – Your OpenAI API key
* `OPENAI_MODEL` – OpenAI model to use (default: `gpt-4`)

---

## Usage

Start the server:

```bash
go run main.go
```

You should see:

```
Server listening on :8090
```

Send a test request:

```bash
curl -X POST http://localhost:8090/query \
     -H "Content-Type: application/json" \
     -d '{"message":"Hello, world!"}'
```

Example response:

```json
{
    "response": "Hi! How can I help you today?"
}
```

---

## API Endpoint

### POST `/query`

**Request body:**

```json
{
    "message": "Hello!"
}
```

**Response:**

```json
{
    "response": "Hi there! I’m here and ready to help with whatever you need."
}
```

---

## Testing

Run automated tests:

```bash
go test ./handlers -v
```

> Make sure your `.env` file contains a valid `OPENAI_API_KEY` before running tests.

---

## Project Structure

```
.
├── README.md
├── config
│   └── env.go          # Environment loading & validation
├── go.mod
├── go.sum
├── handlers
│   ├── llm.go          # Main handler code
│   └── llm_test.go     # Handler tests
└── main.go             # Entry point
```

