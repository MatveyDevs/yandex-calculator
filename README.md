# yandex-calculator
# Go Calculator Server
## 📜 Description
This repository contains a **Go-based** server application that functions as a **calculator**, designed for a Yandex assignment. The app provides a simple HTTP interface for performing basic arithmetic operations.

## ⭐ Features
- **Basic Operations**: Supports addition, subtraction, multiplication, and division.
- **JSON Format**: Utilizes JSON for request and response for easy client interaction.
- **Error Handling**: Robust error handling and input validation for reliability.

## 🛠️ Technologies
- **Go (Golang)**

## 🚀 Installation and Running

1. Clone the repository:
```
git clone https://github.com/MatveyDevs/yandex-calculator.git 
cd yandex-calculator
```    
2. If port 8080 is busy, create an .env file in the root of the project and change the PORT, for example:
```
PORT=8081
```
3. Ensure Go is installed and set up on your machine.

4. Run the server:

```
cd cmd
go run main.go
```    
5. Send requests to `http://localhost:8080` using a browser or Postman.

## 📋 Example Usage

### Addition:
You can perform addition with the following command:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"2+2\"}"
```
#### Response:
```
{
  "result": 4
}
```
### Call Reslting in 422 Unprocessable Entity:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"incorrect\"}"
```

#### Response:
```
{
  "error": "Expression is not valid"
}
```

### Call Reslting in 500 Internal Server Error:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"invalid\": \"value\"}"
```

#### Response:
```
{
  "error": "Internal Server Error"
}
```
Feel free to submit issues or suggestions!

---

# Русская версия

# Серверный калькулятор на Go

## 📜 Описание
Этот репозиторий содержит серверное приложение на **Go**, которое функционирует как **калькулятор**. Оно было разработано в рамках задания от Яндекса и предоставляет простой HTTP интерфейс для выполнения базовых арифметических операций.

## ⭐ Функциональность
- **Основные операции**: Поддержка сложения, вычитания, умножения и деления.
- **Формат JSON**: Формат запроса и ответа JSON для легкого взаимодействия с клиентами.
- **Обработка ошибок**: Надежная обработка ошибок и валидация вводимых данных.

## 🛠️ Технологии
- **Go (Golang)**

## 🚀 Установка и запуск

1. Клонируйте репозиторий:
```
git clone https://github.com/MatveyDevs/yandex-calculator.git 
cd yandex-calculator
```    
2. Если порт 8080 занят, то создайте файл .env в корне проекта и измените PORT, например:
```
PORT=8081
```
3. Убедитесь, что Go установлен и настроен на вашем компьютере.

4. Запустите сервер:
```
cd cmd
go run main.go
```    
5. Отправляйте запросы на `http://localhost:8080` с помощью браузера или Postman.

## 📋 Примеры использования

### Сложение:
Вы можете выполнить сложение следующим образом:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"2+2\"}"
```
#### Ответ:
```
{
  "result": 4
}
```
### Вызов с ошибкой 422 Unprocessable Entity:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"incorrect\"}"
```

#### Ответ:
```
{
  "error": "Expression is not valid"
}
```

### Вызов с ошибкой 500 Internal Server Error:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"invalid\": \"value\"}"
```

#### Ответ:
```
{
  "error": "Internal Server Error"
}
```
Мы рады получить ваши предложения и комментарии!