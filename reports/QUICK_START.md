# 🚀 Cluely MVP - Quick Start Guide

## ⚡ Самый быстрый способ

### 1️⃣ Собрать и запустить одной командой:

```bash
make run
```

Это сделает всё автоматически:
- ✅ Очистит старую сборку
- ✅ Скомпилирует приложение
- ✅ Скопирует конфиг
- ✅ Запустит приложение

### 2️⃣ Открыть браузер:

```
http://localhost:8080
```

---

## 🔨 Отдельные команды

Если хотите контролировать каждый шаг:

```bash
# 1. Собрать приложение
make build

# 2. Запустить (если уже собрано)
make run-binary

# или просто запустить бинарник напрямую
bin/cluely.exe
```

---

## 📋 Все доступные команды

```bash
make help       # Показать все команды
make build      # Собрать (bin/cluely.exe + конфиг)
make run        # Собрать и запустить
make run-binary # Запустить уже собранное
make clean      # Удалить bin папку
make fmt        # Отформатировать код
make lint       # Проверить код (go vet)
make test       # Запустить тесты
make deps       # Скачать зависимости
make dev-setup  # Полная подготовка для разработки
```

---

## 🎯 Результат

После `make run` вы должны увидеть:

```
🚀 Starting Cluely Agent...
✅ UI Server started
✅ Audio Module started (transcriber: mock)
✅ Vision Module started (OCR engine: mock)
✅ AI Module initialized (provider: mock, model: mock)
✅ Cluely Agent started successfully!
📝 Press Ctrl+C to stop...
🌐 UI server listening on http://localhost:8080
```

Затем каждые 5-10 секунд:
- 🎤 Mock аудиотранскрипция
- 📸 Mock OCR
- 🤖 AI анализ
- 💬 Подсказки в UI

---

## 🌐 UI Interface

Откройте: **http://localhost:8080**

Вы увидите:
- ✅ Status: "Connected"
- 💬 Подсказки, обновляющиеся в реальном времени
- ⏰ Timestamp для каждого сообщения
- 📊 Последние 10 сообщений

---

## 🛑 Остановка

Нажмите **Ctrl+C** в консоли

---

## 📁 Что было собрано

```
bin/
├── cluely.exe       # Скомпилированное приложение
└── default.toml     # Конфигурация (автоматически скопирована)
```

---

## ⚙️ Как это работает

1. **Makefile:**
   - Создает папку `bin/`
   - Компилирует Go код в `bin/cluely.exe`
   - Копирует `configs/default.toml` в `bin/default.toml`

2. **main.go:**
   - Ищет конфиг в текущей папке или рядом
   - Находит `default.toml` автоматически
   - Запускает все модули

3. **Модули:**
   - **Audio** - генерирует транскрипции каждые 5 сек
   - **Vision** - генерирует OCR каждые 10 сек
   - **AI** - анализирует контекст и выдает подсказки
   - **UI** - обновляет браузер через WebSocket

---

## 🔍 Troubleshooting

### "command not found: make"
Используйте вместо этого:
```bash
go run ./cmd/cluely/main.go
```

### "failed to load config"
Убедитесь, что в `configs/` папке есть `default.toml`:
```bash
ls configs/default.toml
```

### Порт 8080 занят
Измените в `configs/default.toml`:
```toml
[ui]
port = 8081  # или другой свободный порт
```

---

## 📚 Дополнительно

- **README.md** - Полная документация
- **MVP_COMPLETION_REPORT.md** - Детальный отчет
- **docs/roadmap.md** - Спецификация продукта
- **Makefile** - Все команды сборки

---

## 🎉 Готово!

```bash
make run
```

И наслаждайтесь Cluely MVP! 🚀
