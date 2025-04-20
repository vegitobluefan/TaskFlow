# TaskFlow

TaskFlow — это RESTful API для управления задачами, построенное на языке Go. 
Приложение предоставляет HTTP-интерфейс для создания, просмотра, редактирования, 
удаления и поиска задач. Все данные сохраняются в JSON-файле, что позволяет 
использовать API как лёгкое хранилище без подключения к базе данных.

Технологии:
- Язык: Go
- Архитектура: MVC-подобная
- API: RESTful
- Хранилище: JSON-файл
- Зависимости: Go Modules

### Установка и запуск:
  ```bash
  git clone https://github.com/vegitobluefan/TaskFlow.git
  ```
  ```bash
  cd TaskFlow
  ```
  ```bash
  go run main.go
  ```

API будет доступен по адресу: http://localhost:8080

### Эндпоинты API:
- POST /tasks         — создать новую задачу
- GET /tasks/{id}     — получить задачу по ID

## Разработчик: [Аринов Данияр](https://github.com/vegitobluefan)
