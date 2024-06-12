package main

import "fmt"

/*
Дан набор пар городов:
- между каждой парой городов сотрудник совершил прямой перелёт;
- информация, в каком направлении был совершён перелёт, утеряна.

Известно, что все указанные перелёты относятся к одному путешествию.
Каждый следующий перелёт сотрудник начинал из того города, в котором закончил предыдущий.
Никакой город не был посещён сотрудником дважды.
Город начала путешествия также отличается от конечной точки.
Выведите города в порядке следования по маршруту.

Примеры:
[("Moscow", "Belgrade")] -> ["Moscow", "Belgrade"]
[("Moscow", "Belgrade"), ("Moscow", "Erevan")] -> ["Erevan", "Moscow", "Belgrade"]
[("Moscow", "Erevan"), ("Erevan", "Tokio"), ("Moscow", "Belgrade")] -> ["Tokio", "Erevan", "Moscow", "Belgrade"]
*/

func main() {
	cities := []Flight{
		{"Moscow", "Belgrade"},
		{"Moscow", "Erevan"},
		{"Erevan", "Tokio"},
	}

	fmt.Println(GetRoute(cities))
}

// Flight перелёт между двумя городами, направление неизвестно
type Flight = [2]string

func GetRoute(flights []Flight) []string {
	cityPairs := make(map[string][]string, len(flights)+1)

	// Для каждого города собираем список городов, в которые можно полететь из него.
	for _, flight := range flights {
		// Прямой перелет.
		sities, exists := cityPairs[flight[0]]
		if !exists {
			sities = []string{}
		}
		sities = append(sities, flight[1])
		cityPairs[flight[0]] = sities

		// Обратный перелет.
		sities, exists = cityPairs[flight[1]]
		if !exists {
			sities = []string{}
		}
		sities = append(sities, flight[0])
		cityPairs[flight[1]] = sities
	}

	firstCity := ""

	// Ищем первый город.
	for city, pairs := range cityPairs {
		if len(pairs) == 1 {
			firstCity = city
			break
		}
	}

	result := make([]string, 0, len(flights)+1)
	result = append(result, firstCity)
	currCity := firstCity
	prevCity := firstCity

	for {
		// Берем текущий город, ищем куда был осуществлен перелет.
		pairs := cityPairs[currCity]

		for _, city := range pairs {
			if city != prevCity {
				// Текущий город отмечаем как предыдущий и исключаем его из поиска на следующем этапе.
				// Текущий город обновляем.
				result = append(result, city)
				prevCity = currCity
				currCity = city
			}
		}

		if len(result) == len(flights)+1 {
			break
		}
	}

	return result
}
