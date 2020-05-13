package infrastructure

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleResizeImage(w http.ResponseWriter, r *http.Request) {

	//считываем весь реквест в body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	//создаем структуру
	image := &Image{}

	//парсим json в эту структуру
	err = json.Unmarshal(body, image)

	//формируем ответ передаем в метод структуру и возвращаем ошибку
	err = Resize(*image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Resize(image Image) error {

	if image.Id == "" || image.Height == 0 || image.Width == 0 || len(image.Buffer) == 0 {
		return errors.New("error: data is not correct")
	}
	return nil
}

// выводит весь массив картинок на экран
func GetImages(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(&images)
}

func GetImageId(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for _, item := range images {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Image{})
}

func AddImage(image Image) {

	//рандомим id картинки
	//image.Id = rand.Intn(10000)

	//добавляем картинку в массив картинок
	images = append(images, image)
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for index, item := range images {
		if item.Id == params["id"] {
			images = append(images[:index], images[index+1:]...)
			var image Image
			_ = json.NewDecoder(r.Body).Decode(&image)
			image.Id = params["id"]
			images = append(images, image)
			json.NewEncoder(w).Encode(image)
			return
		}
	}
	json.NewEncoder(w).Encode(images)
}
