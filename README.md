# Simple CLI Task Tracker

Простой CLI-инструмент для управления задачами, написанный на Go. Позволяет добавлять, удалять, обновлять и отслеживать статус ваших задач через командную строку.

## 🚀 Возможности

- ✅ **Добавление задач** - создавайте новые задачи с уникальными ID
- 📝 **Обновление задач** - изменяйте текст существующих задач
- 🗑️ **Удаление задач** - удаляйте ненужные задачи
- 📊 **Просмотр списка** - отображение всех задач с их статусами
- 🔄 **Управление статусами** - отмечайте задачи как "в процессе" или "выполнено"
- 💾 **Автосохранение** - все изменения автоматически сохраняются в JSON-файл

## 📦 Установка

### Требования
- Go 1.16 или выше

### Клонирование репозитория
```bash
git clone https://github.com/your-username/Simple-CLI-Tracker.git
cd Simple-CLI-Tracker
```

### Сборка
```bash
go build -o task-tracker main.go
```

### Запуск
```bash
./task-tracker
```

Или напрямую через Go:
```bash
go run main.go
```

## 🎯 Использование

После запуска программы вы увидите интерактивную командную строку:

```
> 
```

### Доступные команды

| Команда | Описание | Пример |
|---------|----------|--------|
| `add <текст>` | Добавить новую задачу | `add Купить молоко` |
| `list` | Показать все задачи | `list` |
| `update <ID> <новый_текст>` | Обновить задачу | `update 1 Купить хлеб и молоко` |
| `delete <ID>` | Удалить задачу | `delete 1` |
| `in-progress <ID>` | Отметить как "в процессе" | `in-progress 1` |
| `done <ID>` | Отметить как "выполнено" | `done 1` |
| `help` | Показать справку | `help` |
| `exit` | Выйти из программы | `exit` |

### Примеры использования

```bash
# Добавить задачу
> add Изучить Go программирование
Task added: Изучить Go программирование

# Просмотреть все задачи
> list
Tasks:
ID: 1, Task: Изучить Go программирование, Status: pending

# Отметить задачу как выполняемую
> in-progress 1
Task with ID 1 marked as in progress.

# Обновить текст задачи
> update 1 Изучить продвинутое Go программирование
Task with ID 1 updated to: Изучить продвинутое Go программирование

# Отметить как выполненную
> done 1
Task with ID 1 marked as done.

# Удалить задачу
> delete 1
Task with ID 1 deleted.
```

## 📁 Структура проекта

```
Simple-CLI-Tracker/
├── main.go              # Точка входа в приложение
├── internal/
│   └── ui.go            # Обработка команд и пользовательский интерфейс
├── task/
│   ├── task.go          # Структура задачи
│   └── storage.go       # Функции сохранения/загрузки
├── tasks.json           # Файл данных (создается автоматически)
├── go.mod              # Go модуль
├── .gitignore          # Git ignore правила
└── README.md           # Этот файл
```

## 💾 Хранение данных

Все задачи сохраняются в файле `tasks.json` в формате:

```json
[
  {
    "id": 1,
    "text": "Изучить Go программирование",
    "status": "pending"
  },
  {
    "id": 2,
    "text": "Создать CLI приложение",
    "status": "done"
  }
]
```

### Статусы задач
- `pending` - ожидает выполнения (по умолчанию)
- `in-progress` - в процессе выполнения
- `done` - выполнено

## 🛠️ Разработка

### Сборка для разных платформ

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o task-tracker.exe main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o task-tracker main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o task-tracker main.go
```

### Запуск тестов
```bash
go test ./...
```

## 🤝 Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для новой функции (`git checkout -b feature/amazing-feature`)
3. Внесите изменения и зафиксируйте их (`git commit -m 'Add amazing feature'`)
4. Отправьте изменения в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request


## 👨‍💻 Автор

**atlkhnv** - [GitHub Profile](https://github.com/atlkhnv)