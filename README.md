# Вычислитель отличий (Go)

Вычислитель отличий - CLI-утилита, определяющая разницу между двумя структурами данных. Это популярная задача, для решения которой существует множество онлайн сервисов, например [jsondiff](http://www.jsondiff.com/). Подобный механизм используется при выводе тестов или при автоматическом отслеживании изменений в конфигурационных файлах.

## Возможности

- Поддержка разных входных форматов: `yaml`, `json`
- Генерация отчета в виде `plain text`, `stylish` и `json`

## Установка

Сборка из исходников:

```bash
git clone https://github.com/xhrobj-hex/go-project-244.git
cd go-project-244
make build
```

Бинарник появится в `./bin/gendiff`.

## Флаги

- `--format`, `-f` - формат вывода (`stylish` по умолчанию)
- `--help`, `-h` - справка

## Примеры

Показать справку:

```bash
./bin/gendiff --help
```

---

## Hexlet tests and linter status:
[![Actions Status](https://github.com/xhrobj-hex/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/xhrobj-hex/go-project-244/actions)

## Project CI - lint & tests
[![(-_-) go-ci](https://github.com/xhrobj-hex/go-project-244/actions/workflows/go-ci.yml/badge.svg)](https://github.com/xhrobj-hex/go-project-244/actions/workflows/go-ci.yml)
