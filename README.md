# 🎓 EnrollMate

<div align="center">
  
![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen?style=for-the-badge)

</div>

<p align="center">
  <strong>A lightweight student enrollment management system built with Go and the Gin web framework</strong>
</p>

## 📋 Overview

EnrollMate provides a clean, modern interface for managing student enrollments. The system is designed with simplicity and performance in mind, offering essential functionality for educational institutions to track and manage their student data.

## ✨ Features

- 👥 **Student Management**: Complete CRUD operations for student records
- 🎨 **Clean Web Interface**: Modern, responsive design with an intuitive user experience  
- 🔗 **RESTful API**: Well-structured endpoints for seamless integration
- ⚡ **Lightweight Architecture**: Minimal dependencies for optimal performance
- 📄 **JSON Support**: Full JSON serialization for API responses

## 🛠️ Tech Stack

- 🔧 **Backend**: Go 1.24.6
- 🌐 **Web Framework**: Gin-Gonic
- 🏗️ **Architecture**: Clean architecture with separate handlers, models, and routes
- 📊 **Data Format**: JSON for API communication

## 📁 Project Structure

```
enrollmate/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── model/          # Data models and structures
│   └── routes/         # Route definitions
├── go.mod              # Go module configuration
└── go.sum              # Dependency checksums
```

## 🚀 Getting Started

### 📋 Prerequisites

- Go 1.24.6 or higher
- Git

### 💻 Installation

1. **Clone the repository**
```bash
git clone https://github.com/alberto-debug/enrollmate.git
cd enrollmate
```

2. **Install dependencies**
```bash
go mod download
```

3. **Run the application**
```bash
go run cmd/server/server.go
```

🌐 The server will start on `http://localhost:8080`

### 🏗️ Building for Production

```bash
go build -o enrollmate cmd/server/server.go
./enrollmate
```

## 🔨 Development

### 🆕 Adding New Features

1. **📁 Models**: Add new data structures in `internal/model/`
2. **⚙️ Handlers**: Implement business logic in `internal/handlers/`
3. **🛣️ Routes**: Register new endpoints in `internal/routes/`

### 📝 Code Style

This project follows Go's standard formatting guidelines. Run `go fmt` before committing changes.

## 🤝 Contributing

1. 🍴 Fork the repository
2. 🌟 Create a feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'Add amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing-feature`)
5. 🔄 Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 📧 Contact

**Alberto** - [@alberto-debug](https://github.com/alberto-debug)

Project Link: [https://github.com/alberto-debug/enrollmate](https://github.com/alberto-debug/enrollmate)

---

*Built with ❤️ using Go and Gin*
