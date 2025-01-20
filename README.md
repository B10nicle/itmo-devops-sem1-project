# Финальный проект 1 семестра

REST API сервис для загрузки и выгрузки данных о ценах.

## Требования к системе

- Go 1.23
- PostgreSQL 15

## Алгоритм взаимодействия с приложением

1. Выдача необходимых прав для исполнения скриптов
   ```bash
   chmod +x ./scripts/make_executable.sh
   ./scripts/make_executable.sh
    ```

2. Выполнение миграций и компиляция приложения
    ```bash
    ./scripts/prepare.sh
    ```

3. Запуск приложения
    ```bash
    ./scripts/run.sh
    ```

4. Запуск тестов
    ```bash
    ./scripts/all_tests.sh
    ```

## Контакты

Telegram: @B10nicle (Хилько Олег)