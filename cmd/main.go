package cmd

// Подключаемся к бд

func CinnectDB() {

}

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	// Создаем конфигурацию приложения
	conf := config.NewAppConf()

	// Создаем таблицу

	// Инициализируем слои и запускаем хэндлер
}
