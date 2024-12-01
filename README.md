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


    git clone https://github.com/MatveyDevs/yandex-calculator.git 
    cd yandex-calculator
2. Ensure Go is installed and set up on your machine.

3. Run the server:


    cd cmd
    go run main.go
4. Send requests to `http://localhost:8080` using a browser or Postman.

## 📋 Example Usage

### Addition:
You can perform addition with the following command:

      curl -X POST http://localhost:8080/calc -H "Content-Type: application/json" -d "{\"expression\": \"2+2\"}"

#### Response:
    {
    "result": 4
    }

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


    git clone https://github.com/MatveyDevs/yandex-calculator.git 
    cd yandex-calculator
2. Убедитесь, что Go установлен и настроен на вашем компьютере.

3. Запустите сервер:


    cd cmd
    go run main.go
4. Отправляйте запросы на `http://localhost:8080` с помощью браузера или Postman.

## 📋 Примеры использования

### Сложение:
Вы можете выполнить сложение следующим образом:

      curl -X POST http://localhost:8080/calc -H "Content-Type: application/json" -d "{\"expression\": \"2+2\"}"

#### Ответ:
    {
    "result": 4
    }

Мы рады получить ваши предложения и комментарии!