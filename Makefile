.PHONY: help test coverage examples example-01 example-02 example-03 example-04 example-05 example-06 example-07 example-08 example-09 example-10 example-11 clean

help:
	@echo "Доступные команды:"
	@echo "  make test          - Запустить все тесты"
	@echo "  make coverage      - Запустить тесты с покрытием"
	@echo "  make examples      - Запустить все примеры"
	@echo "  make example-01    - Запустить пример 01 (Кабель ППГ)"
	@echo "  make example-02    - Запустить пример 02 (Кабель UTP)"
	@echo "  make example-03    - Запустить пример 03 (Оптоволокно)"
	@echo "  make example-04    - Запустить пример 04 (Кабель-канал)"
	@echo "  make example-05    - Запустить пример 05 (Лоток)"
	@echo "  make example-06    - Запустить пример 06 (Дымовые извещатели)"
	@echo "  make example-07    - Запустить пример 07 (Кабельные стяжки)"
	@echo "  make example-08    - Запустить пример 08 (Шурупы)"
	@echo "  make example-09    - Запустить пример 09 (Комплексный склад)"
	@echo "  make example-10    - Запустить пример 10 (Настройка точности)"
	@echo "  make example-11    - Запустить пример 11 (Математические операции)"
	@echo "  make clean         - Очистить бинарные файлы"

test:
	@echo "=== Запуск тестов ==="
	go test -v ./...

coverage:
	@echo "=== Запуск тестов с покрытием ==="
	go test -v -coverprofile=coverage.out ./...
	@echo ""
	@echo "=== Покрытие по файлам ==="
	go tool cover -func=coverage.out

examples: example-01 example-02 example-03 example-04 example-05 example-06 example-07 example-08 example-09 example-10 example-11

example-01:
	@echo "=== Пример 01: Кабель ППГ ==="
	@cd examples/01_cable_ppg && go run main.go
	@echo ""

example-02:
	@echo "=== Пример 02: Кабель UTP ==="
	@cd examples/02_cable_utp && go run main.go
	@echo ""

example-03:
	@echo "=== Пример 03: Оптоволокно ==="
	@cd examples/03_fiber_optic && go run main.go
	@echo ""

example-04:
	@echo "=== Пример 04: Кабель-канал ==="
	@cd examples/04_cable_duct && go run main.go
	@echo ""

example-05:
	@echo "=== Пример 05: Лоток ==="
	@cd examples/05_tray && go run main.go
	@echo ""

example-06:
	@echo "=== Пример 06: Дымовые извещатели ==="
	@cd examples/06_smoke_detector && go run main.go
	@echo ""

example-07:
	@echo "=== Пример 07: Кабельные стяжки ==="
	@cd examples/07_cable_ties && go run main.go
	@echo ""

example-08:
	@echo "=== Пример 08: Шурупы ==="
	@cd examples/08_screws && go run main.go
	@echo ""

example-09:
	@echo "=== Пример 09: Комплексный склад ==="
	@cd examples/09_warehouse_complex && go run main.go
	@echo ""

example-10:
	@echo "=== Пример 10: Настройка точности ==="
	@cd examples/10_precision && go run main.go
	@echo ""

example-11:
	@echo "=== Пример 11: Математические операции ==="
	@cd examples/11_math_operations && go run main.go
	@echo ""

clean:
	@echo "=== Очистка бинарных файлов ==="
	@find examples -name "0*_*" -type f ! -name "*.go" -delete
	@rm -f coverage.out
	@echo "Готово"
