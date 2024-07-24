# Todo List Microservice

## Описание проекта

Этот проект представляет собой микросервис для управления списками задач (Todo List). Он реализует RESTful API, позволяющее создавать, обновлять, удалять задачи и помечать их как выполненные. Также он поддерживает фильтрацию задач по статусу и сортировку по дате.

## Требования

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)
- [Render](https://render.com/) для деплоя (или другой облачный сервис)

## Установка

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/gabdylgaziz/go-todo.git
    cd proxy-server
    ```

2. Установите зависимости и запустите сервер:

    ```sh
    go run main.go
    ```

### Docker

1. Соберите и запустите контейнер с помощью Docker Compose:

    ```sh
    make build
   make up
    ```

Проект будет доступен по адресу http://localhost:8080.

**Создание заказа**

Отправьте POST-запрос на `http://localhost:8080/api/todo-list/tasks` с телом запроса в формате JSON:

```json
{
   "title": "Купить книгу",
   "active_at": "2023-08-04"
}
```

Ответ будет в формате JSON:

```json
{
   "id": "123",
   "title": "Купить книгу",
   "active_at": "2023-08-04",
   "done": false
}
```

**Обновление существующей задачи**

Отправьте PUT-запрос на `http://localhost:8080/api/todo-list/tasks/{id}` с телом запроса в формате JSON:

```json
{
   "title": "Купить книгу - Обновлено",
   "active_at": "2023-08-05"
}
```

Ответ:

```204 No Content```

**Удаление задачи**

Отправьте DELETE-запрос на `http://localhost:8080/api/todo-list/tasks/{id}`:

Ответ:

```204 No Content```

**Помечание задачи выполненной**

Отправьте PUT-запрос на `http://localhost:8080/api/todo-list/tasks/{id}/done`:

Ответ:

```204 No Content```



**Список задач по статусу**

Отправьте GET-запрос на `/api/todo-list/tasks?status=active` или `/api/todo-list/tasks?status=done`:

Ответ будет в формате JSON:

```json
[
   {
      "id": "123",
      "title": "Купить книгу",
      "active_at": "2023-08-04",
      "done": false  
   },
   {
      "id": "124",
      "title": "Купить квартиру",
      "active_at": "2023-08-05",
      "done": true
   }
]
```