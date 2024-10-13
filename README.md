# World capitals quiz CLI with REST API

This project is a world capitals quiz application built with Golang. It consists of two components:

**Backend**: A REST API that serves quiz questions and processes user submissions.
**CLI**: A CLI using Cobra to interact with the quiz, fetch questions and submit answers.

## Installation

1. **Clone the repository**

```bash
git clone https://github.com/NicolasBors/world-capitals-quiz
cd world-capitals-quiz
```

2. **Install dependencies**

Make sure you have Go installed, then run the following command to intsall Cobra:

```bash
go get -u github.com/spf13/cobra
```

## Running the API

To start the REST API, run:

```bash
go run api/*.go
```

This will start the API on http://localhost:8080, serving the quiz questions and processing user submissions.

## Running the CLI

The CLI allows users to take the quiz by fetching questions from the API and submitting their answsers.

To start the CLI:

```bash
go run main.go start
```

The CLI will prompt you to answer the questions and submit your answers. At the end, you'll get feedback on your performance.