# 🚀 Running the Application with Air 🚀

This project utilizes **Air** for live reloading in a Golang development environment. ⚡

## 🔧 Prerequisites 🔧

Ensure you have the following installed on your system:

- **Go** 🛠️
- **Air** (version **v1.60.0**) 🌬️

## 📥 Installation 📥

### Install Air version **v1.60.0**:
```sh
go install github.com/air-verse/air@v1.60.0
```

### Ensure dependencies are installed:
```sh
go mod download
```

## 🚀 Running the Application 🚀

Start the application using Air:
```sh
air
```
Or run it manually using:
```sh
go run .
```

## 📂 Project Structure 📂

```
tripatra-dct-service-config/
├── context/        # Middleware and shared context utilities 🌍
├── database/       # Database configurations, models, and repositories 🗄️
├── manifests/      # Deployment manifests 📜
├── resolver/       # Resolvers 📡
├── routes/         # API route definitions 🛤️
├── services/       # Business logic services ⚙️
├── utils/          # Utility functions 🔧
├── validators/     # Input validation logic ✅
├── .air.toml       # Air configuration file 💨
├── .env            # Environment variables 🌱
├── app.go          # Application entry point 🚀
├── main.go         # Main execution file 🎯
├── Dockerfile      # Docker configuration 🐳
└── README.md       # Project documentation 📖
```

## 🔄 Live Reloading with Air 🔄

- Any changes made to the source code will automatically restart the application. 🔥
- Configurations for Air are stored in `.air.toml`. 📄
- Logs will display file changes and restarts in real time. ⏳

## ⏹️ Stopping the Application ⏹️

To stop the application running with Air, press **Ctrl + C** in the terminal. ❌

## 📌 Additional Notes 📌

- Ensure your `.env` file is properly configured before running the application. 📝
- Modify the Air configuration in `.air.toml` if needed. 🛠️

### 🎉 Happy Coding! 🚀🚀🚀🚀🚀