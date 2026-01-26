# 🎓 Class Monitor API

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green)

A lightweight **REST API** built with Golang to manage student attendance and behavior data. Originally a CLI tool, now upgraded to a web microservice.

## 🚀 Features

* **RESTful Architecture**: Serves data over HTTP.
* **JSON Response**: Standard data format for web integration.
* **In-Memory Database**: Fast data retrieval using Go Maps.
* **Concurrency Ready**: Built on Go's robust `net/http` package.

## 🛠️ Tech Stack

* **Language**: Go (Golang)
* **Standard Libs**: `net/http`, `encoding/json`

## 🏃‍♂️ How to Run

1.  **Clone the repository**
    ```bash
    git clone [https://github.com/ghosthouse7/class-monitor-cli.git](https://github.com/ghosthouse7/class-monitor-cli.git)
    cd class-monitor-cli
    ```

2.  **Start the Server**
    ```bash
    go run main.go
    ```
    *You will see: `Server starting on localhost:8080...`*

3.  **Test the API**
    Open your browser or Postman and visit:
    ```
    http://localhost:8080/students
    ```

## 🔌 API Endpoints

| Method | Endpoint    | Description             |
| :----- | :---------- | :---------------------- |
| `GET`  | `/students` | Fetch all students JSON |

---
*Built with ❤️ by ghost_hunter*
