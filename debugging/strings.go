package debugging

import "strings"

func concat(values []string) string {
	sb := strings.Builder{} // Создается strings.Builder
	for _, value := range values {
		_, _ = sb.WriteString(value) // Добавляется строка
	}
	return sb.String() // Возвращается результирующая строка
}
