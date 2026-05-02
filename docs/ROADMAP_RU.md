# Roadmap проекта Shuffle

## 1) Что уже есть в проекте (краткий аудит)

- CLI с базовыми командами: `ls`, `view`, `download`.
- Базовая доменная модель (`Artifact`, `Locator`, `Storage`, `Source`).
- Слой application/use-case для list/read/download.
- Поддержка только GitHub как source type в конфиге.
- Конфиг с одним типом источника (`github`) и привязкой к owner/repo/ref.

---

## 2) Продуктовая цель

Сделать единый CLI-клиент для **поиска, просмотра и скачивания скриптов** из удалённых источников:

- GitHub
- GitLab
- Nexus Repository (Sonatype)
- (позже) другие registry/repository endpoints

С фокусом на:

- единый UX команд;
- расширяемость через адаптеры провайдеров;
- воспроизводимость (ref/version/hash);
- безопасность (валидация, checksum, trust policy).

---

## 3) Детальный roadmap (эпики → задачи → подзадачи)

## Эпик A. Платформа источников (multi-source core)

### Задача A1. Унифицировать модель источников
- [ ] Расширить `SourceType` для `gitlab`, `nexus`.
- [ ] Добавить provider-specific поля в конфиг (base_url/project_id/group/repository/raw_api и т.д.).
- [ ] Ввести строгую валидацию конфига (обязательные поля, enum, формат ref).
- [ ] Добавить версионирование формата конфига (например `version: 1`).

### Задача A2. Source factory / registry
- [ ] Реализовать фабрику провайдеров по `type`.
- [ ] Вынести общие интерфейсы transport/auth/pagination.
- [ ] Добавить интеграционные моки для provider contract tests.

### Задача A3. Ошибки и диагностика
- [ ] Ввести типизированные ошибки (`ErrNotFound`, `ErrUnauthorized`, `ErrRateLimited`).
- [ ] Добавить машиночитаемые коды ошибок и дружелюбные сообщения для CLI.
- [ ] Добавить `--debug` режим с trace-id запроса.

## Эпик B. GitLab адаптер

### Задача B1. API-клиент GitLab
- [ ] Поддержать token auth (`PRIVATE-TOKEN` / OAuth token).
- [ ] Реализовать list tree (path/ref, recursive=false по умолчанию).
- [ ] Реализовать чтение raw file и metadata.

### Задача B2. Маппинг доменной модели
- [ ] Преобразование GitLab file/tree сущностей в `domain.Artifact`.
- [ ] Нормализация путей и размеров.
- [ ] Обработка edge-cases: symlink/submodule/empty dirs.

### Задача B3. Тесты GitLab
- [ ] Unit-тесты mapper и adapter.
- [ ] Контрактные тесты на фикстурах API.
- [ ] Smoke-тест на публичном тестовом репозитории (опционально по env-флагу).

## Эпик C. Nexus (Sonatype) адаптер

### Задача C1. Поддержка browse/search/download
- [ ] Поддержать browsing компонентов/ассетов в hosted/proxy репозиториях.
- [ ] Поддержать фильтрацию по path/prefix/extension.
- [ ] Реализовать скачивание бинарных и текстовых скриптов.

### Задача C2. Аутентификация и endpoint-профили
- [ ] Basic auth + bearer (где применимо).
- [ ] Поддержка кастомного base URL на инстанс.
- [ ] Нормализация URL и retry policy.

### Задача C3. Тестирование Nexus
- [ ] Набор mock-ответов для API search/assets.
- [ ] E2E-тест в docker-compose окружении (по возможности).
- [ ] Проверка корректной работы больших файлов.

## Эпик D. CLI/UX и сценарии пользователя

### Задача D1. Единый CLI UX
- [ ] Добавить команду `sources` (list/add/remove/validate).
- [ ] Добавить `--source`, `--ref`, `--format=json|table`.
- [ ] Поддержать постраничный вывод и сортировку.

### Задача D2. Поиск скриптов
- [ ] Команда `search` по имени/маске/расширению.
- [ ] Полнотекст (если доступен у провайдера) или fallback по path.
- [ ] Вывод score + origin (source, repo, ref).

### Задача D3. Скачивание и кэш
- [ ] Локальный кэш метаданных и файлов.
- [ ] Флаги `--force`, `--no-cache`, `--checksum`.
- [ ] Атомарная запись на диск и resume (если возможно).

## Эпик E. Надежность и безопасность

### Задача E1. Сетевой слой
- [ ] Backoff/retry для 429/5xx.
- [ ] Таймауты по операциям list/read/download.
- [ ] Circuit-breaker/лимиты конкурентных запросов.

### Задача E2. Безопасность цепочки поставки
- [ ] Проверка checksum (SHA256).
- [ ] Опциональная верификация подписи/attestation (позже).
- [ ] Правила trust policy (разрешённые source/repo/path).

### Задача E3. Логи и аудит
- [ ] Структурированные логи (json mode).
- [ ] Маскирование секретов в логах.
- [ ] Audit trail: кто/что/откуда было скачано.

## Эпик F. Качество, релизы и DX

### Задача F1. Тестовая стратегия
- [ ] Unit coverage для domain/app/adapters.
- [ ] Контрактные тесты для каждого провайдера.
- [ ] Golden tests для CLI-вывода.

### Задача F2. CI/CD
- [ ] Линтеры (`go vet`, `staticcheck`, `golangci-lint`).
- [ ] Матрица тестов на версиях Go.
- [ ] Сборка release-артефактов для Linux/macOS/Windows.

### Задача F3. Документация
- [ ] Быстрый старт + examples для всех source types.
- [ ] Справочник конфигурации.
- [ ] Troubleshooting (auth/rate limit/network).

---

## 4) Спринт на 10 рабочих дней (2 недели)

## Sprint Goal
За 10 рабочих дней вывести MVP v0.2: **GitHub + GitLab**, унифицированный конфиг, базовый поиск/листинг/скачивание, тесты и CI-checks.

## Scope (commitment)

### День 1: Архитектура и планирование
- [ ] Уточнить требования MVP (обязательные сценарии).
- [ ] Зафиксировать schema конфига v1.
- [ ] Спроектировать source factory + interfaces.

### День 2: Core refactoring
- [ ] Внести изменения в `config` и валидацию.
- [ ] Реализовать factory/registry провайдеров.
- [ ] Добавить типизированные ошибки.

### День 3: GitHub hardening
- [ ] Привести GitHub-клиент к production базовым практикам (timeouts/retry/error mapping).
- [ ] Выровнять output/listing contract.
- [ ] Закрыть unit-тестами mapper + adapter.

### День 4: GitLab client
- [ ] Реализовать базовый API-клиент GitLab.
- [ ] Поддержать list/read/download.
- [ ] Поддержать token auth и base URL.

### День 5: GitLab adapter + tests
- [ ] Доделать маппинг в domain.Artifact.
- [ ] Написать unit/contract tests.
- [ ] Провести smoke-проверки CLI-команд.

### День 6: CLI UX улучшения
- [ ] Добавить `--source`, `--format`.
- [ ] Добавить команду `sources list`.
- [ ] Улучшить тексты ошибок и help.

### День 7: Поиск и фильтрация
- [ ] Добавить команду `search` (MVP: path/name filter).
- [ ] Добавить фильтры по extension/path prefix.
- [ ] Добавить сортировку и ограничение выдачи.

### День 8: Кэш и надежность
- [ ] Реализовать локальный кэш метаданных.
- [ ] Добавить `--no-cache`/`--force`.
- [ ] Добавить backoff для 429/5xx.

### День 9: CI + документация
- [ ] Настроить CI pipeline: lint + tests.
- [ ] Добавить quickstart и конфиг-примеры.
- [ ] Подготовить changelog драфт.

### День 10: Стабилизация и релиз
- [ ] Bugfix day и регрессия по сценариям.
- [ ] Freeze scope, tag `v0.2.0`.
- [ ] Релизные заметки и план следующего спринта.

## Definition of Done для спринта
- [ ] GitHub и GitLab источники работают end-to-end.
- [ ] `ls/view/download/search` работают по единому контракту.
- [ ] Покрытие критических use-case тестами.
- [ ] CI зелёный, релизный артефакт собирается.

## Риски и буферы
- Риск: различия API GitLab/GitHub в tree/list semantics.
- Риск: rate limits и нестабильные integration-тесты.
- Буфер: 15–20% времени (день 10 + часть дня 9).

## Stretch goals (если успеете)
- Базовая заготовка Nexus adapter (read-only list).
- Проверка checksum при скачивании.
- JSON-вывод для всех команд.
