package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Artist struct {
	ID    string   `json:"id"`    // id коллектива
	Name  string   `json:"name"`  // название группы
	Born  string   `json:"born"`  // год основания группы
	Genre string   `json:"genre"` // жанр
	Songs []string `json:"songs"` // популярные песни, это слайс строк, так как песен может быть несколько
}

var artists = map[string]Artist {
    "1": {
            ID: "1",
            Name: "30 Seconds To Mars",
            Born: "1998",
            Genre: "alternative",
            Songs: []string {
                "The Kill",
                "A Beautiful Lie",
                "Attack",
                "Live Like A Dream",
        },
    },
    "2": {
            ID: "2",
            Name: "Garbage",
            Born: "1994",
            Genre: "alternative",
            Songs: []string {
                "Queer",
                "Shut Your Mouth",
                "Cup of Coffee",
                "Til the Day I Die",
        },
    },
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	// сериализуем данные из слайса artists
	resp, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, "json не сериализовался", http.StatusInternalServerError)
		return
	}

	// в заголовок записываем тип контента, у нас это данные в формате JSON
	w.Header().Set("Content-Type", "application/json")
	// так как все успешно, то статус OK
	w.WriteHeader(http.StatusOK)
	// записываем сериализованные в JSON данные в тело ответа
	w.Write(resp)
}

func main() {
	r := chi.NewRouter()

	r.Get("/artists", getArtists)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
