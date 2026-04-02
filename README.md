# Вычислитель отличий (Go)

Вычислитель отличий - CLI-утилита, определяющая разницу между двумя структурами данных. Это популярная задача, для решения которой существует множество онлайн-сервисов, например [jsondiff](http://www.jsondiff.com/). Подобный механизм используется при выводе тестов или при автоматическом отслеживании изменений в конфигурационных файлах.

## Возможности

- Поддержка разных входных форматов: `yaml`, `json`
- Генерация отчета в виде `stylish`, `plain`,  и `json`

## Структура проекта

```text
cmd/gendiff/         -> точка входа CLI
internal/parser/     -> парсинг JSON/YAML
internal/diff/       -> построение дерева различий
internal/formatters/ -> форматирование результата
testdata/fixture/    -> тестовые файлы
```

## Установка

Сборка из исходников:

```bash
git clone https://github.com/xhrobj-hex/go-project-244.git
cd go-project-244
make build
```

Бинарник появится в `./bin/gendiff`.

## Использование

```text
./bin/gendiff <filepath1> <filepath2> [flags]
```

### Флаги

- `--help`, `-h` - показать справку
- `--format`, `-f` - формат вывода

Поддерживаются форматы:

- `stylish` — формат по умолчанию
- `plain` — плоское текстовое описание изменений
- `json` — вывод результата в формате JSON

### Примеры

Показать справку:

```bash
./bin/gendiff --help
```

Сравнить два JSON-файла:

```bash
./bin/gendiff testdata/fixture/file5.json testdata/fixture/file6.json
```

Сравнить два YAML-файла:

```bash
./bin/gendiff testdata/fixture/file5.yml testdata/fixture/file6.yml
```

Сравнить JSON и YAML:

```bash
./bin/gendiff testdata/fixture/file5.json testdata/fixture/file6.yml
```

Вывести результат в формате `plain`:

```bash
./bin/gendiff --format plain testdata/fixture/file5.json testdata/fixture/file6.json
```

Вывести результат в формате `json`:

```bash
./bin/gendiff --format json testdata/fixture/file5.json testdata/fixture/file6.json
```

## Демонстрация

🎬 Видео с примером работы программы записано с помощью **asciinema**:

[![asciicast](https://asciinema.org/a/XocMUnE2Aq4ZAaVY.svg)](https://asciinema.org/a/XocMUnE2Aq4ZAaVY)

---

## Hexlet tests and linter status

[![Actions Status](https://github.com/xhrobj-hex/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/xhrobj-hex/go-project-244/actions)

## Project CI - lint & tests

[![(-_-) go-ci](https://github.com/xhrobj-hex/go-project-244/actions/workflows/go-ci.yml/badge.svg)](https://github.com/xhrobj-hex/go-project-244/actions/workflows/go-ci.yml)

## SonarQube statuses

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=xhrobj-hex_go-project-244&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=xhrobj-hex_go-project-244)

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=xhrobj-hex_go-project-244&metric=coverage)](https://sonarcloud.io/summary/new_code?id=xhrobj-hex_go-project-244)
