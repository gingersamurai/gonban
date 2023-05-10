# gonban

REST API сервис для работы с kanban доской, написанный на языке *Go*.

##  Установка и конфигурация
Конфигурация сервиса происходит с помощью *yaml* файла.
Также, некоторые данные в целях безопасности 
должны храниться в переменных окружения 
на *production* сервере.\
Для начала нужно запустить базу данных,
потом накатить на неё миграции.

## Использование

### Сервис на текущий поддерживает 4 эндпоинта:
+ `GET /tasks` возвращает список всех задач
+ `GET /tasks/$id` возращает задачу по индексу `$id`
+ `POST /tasks` добавляет новую задачу по её *json* схеме
+ `DELETE /tasks/$id` удаляет задачу по индексу `$id`

### Схема задачи:

```go
type Task struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Performer   string    `json:"performer"`
	Deadline    time.Time `json:"deadline"`
}
```

# Архитектура проекта
На текущий момент на архитекуру проекта можно посмотреть по [ссылке](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1#G1P_K1QbEO4mQN1rD4mX4H2Hf9h5R8DbL-)