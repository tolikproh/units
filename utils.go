package units

import (
	"errors"
	"fmt"
	"strconv"
)

// FormatWithDecimals форматирует число в строку с указанным количеством знаков после запятой
func FormatWithDecimals(q Quantiter, value uint64, divisor Prefix, decimals uint) (string, error) {
	sname := q.ShortName(divisor)
	if sname == "" {
		return "", errors.New("указанный префикс к данному типу применить нельзя")
	}

	quotient := value / uint64(divisor)
	remainder := value % uint64(divisor)

	// Подготавливаем дробную часть
	decimalStr := fmt.Sprintf("%d", remainder)
	// Добавляем ведущие нули если нужно
	for len(decimalStr) < 3 {
		decimalStr = "0" + decimalStr
	}

	// Если нужно округление
	if decimals < uint(len(decimalStr)) {
		// Получаем следующую цифру после места округления
		nextDigit := 0
		if decimals < uint(len(decimalStr)) {
			nextDigit = int(decimalStr[decimals] - '0')
		}

		// Округляем если следующая цифра >= 5
		if nextDigit >= 5 {
			// Конвертируем обрезанную часть в число для округления
			numStr := decimalStr[:decimals]
			num, _ := strconv.ParseUint(numStr, 10, 64)
			num++ // увеличиваем на 1 для округления

			// Проверяем переполнение
			if len(fmt.Sprint(num)) > int(decimals) {
				quotient++
				num = 0
			}

			// Форматируем обратно в строку с ведущими нулями
			decimalStr = fmt.Sprintf("%0*d", decimals, num)
		} else {
			decimalStr = decimalStr[:decimals]
		}
	}

	if decimals == 0 {
		return fmt.Sprintf("%d %s", quotient, sname), nil
	}
	return fmt.Sprintf("%d.%s %s", quotient, decimalStr, sname), nil
}
