# Class Monitor CLI

>> A high-performance Command Line Interface (CLI) tool built with **Go (Golang)** to manage student attendance and records efficiently.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active_Development-green?style=for-the-badge)

##  About The Project
This project starts as a simple CLI tool to manage student data using in-memory data structures but will evolve into a containerized application orchestrated by Kubernetes.

**Current Features:**
- **Struct-based Data Modeling:** Clean architecture for student data.
- ** fast Lookups:** Uses `Maps` (Hash Tables) for O(1) data retrieval.
- **Error Handling:** Gracefully handles non-existent records (The "Ghost Student" check).

## Tech Stack
- **Language:** Go (Golang)
- **Concepts:** Structs, Maps, Pointers, Slices

##  How to Run

1. **Clone the repository**
   ```bash
   git clone [https://github.com/ghosthouse7/class-monitor-cli.git](https://github.com/ghosthouse7/class-monitor-cli.git)
