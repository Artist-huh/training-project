package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Pic(dx, dy int) [][]uint8 {
	// Создаём внешний срез длиной dy (строки)
	image := make([][]uint8, dy)

	// Заполняем каждую строку
	for y := 0; y < dy; y++ {
		// Создаём строку длиной dx (столбцы)
		image[y] = make([]uint8, dx)

		// Заполняем каждый пиксель в строке
		for x := 0; x < dx; x++ {
			// Можете менять формулу для разных узоров:
			// (x+y)/2 — градиент
			// x*y     — параболический узор
			// x^y     — фрактальный XOR-узор
			image[y][x] = uint8((x + y) / 2)
		}
	}

	return image
}

func main2() {
	// Размеры изображения
	const width, height = 256, 256

	// Получаем данные пикселей
	pixels := Pic(width, height)

	// Создаём новое изображение RGBA
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Заполняем изображение пикселями
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Преобразуем значение в оттенок синего
			gray := pixels[y][x]
			img.Set(x, y, color.RGBA{
				R: gray, // Немного красного для разнообразия
				G: 0,    // Без зелёного
				B: gray, // Основной синий
				A: 255,  // Полностью непрозрачный
			})
		}
	}

	// Создаём файл для сохранения
	file, err := os.Create("picture.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Кодируем изображение в PNG и сохраняем
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}

	println("Изображение сохранено как picture.png")
}