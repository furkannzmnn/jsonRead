package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {

	memoryEffectReadJson("https://jsonplaceholder.typicode.com/photos")

}

func memoryEffectReadJson(url string) {
	var photos []Photo
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetcing data:", err)
	}

	decoder := json.NewDecoder(response.Body)

	for decoder.More() {
		err := decoder.Decode(&photos)
		if err != nil {
			fmt.Println(err)
		}
		for _, photo := range photos {
			fmt.Println("title:", photo.Title)
			fmt.Println("ThumbnailURL:", photo.ThumbnailURL)
			fmt.Println("url:", photo.URL)
		}
	}

	if err != nil {
		log.Fatal(err)
	}

}

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func parseJson(fetchUrl string) {

	log.Println("initialize fetch")
	fetchUrl = "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(fetchUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	var fake Fake
	err = json.NewDecoder(response.Body).Decode(&fake)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fake.Title)
	fmt.Println(fake.Completed)
	fmt.Println(fake.UserID)
	fmt.Println(fake.ID)
}

type Fake struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func readJsonLargeFile(name string) {
	jsonFile, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
	}

	encoder := json.NewDecoder(jsonFile)

	_, err = encoder.Token()

	if err != nil {
		fmt.Println(err)
	}

	var pokemon Pokemon
	for encoder.More() {
		err := encoder.Decode(&pokemon)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pokemon.Address)
	}

	_, err = encoder.Token()

	if err != nil {
		log.Fatal(err)
	}
}

func openJson(name string) {
	file, err := os.Open(name)

	if err != nil {
		fmt.Println(err, file.Name())
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	var pokemon []Pokemon
	err = json.Unmarshal(bytes, &pokemon)
	if err != nil {
		return
	}

	for _, yaz := range pokemon {
		fmt.Println(yaz.Greeting)
	}
}

func connect() {
	log.Println("initialize")
	conn, err := net.Listen("tcp", "localhost:8080")

	fmt.Println(conn.Addr())
	if err != nil {
		fmt.Println(err)
	}

	defer func(conn net.Listener) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	for {
		listener, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go proxy(listener)
	}

}

func proxy(listener net.Conn) {
	dial, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		proxy(nil)
	}
	go func() {
		_, err := io.Copy(dial, listener)
		if err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		_, err := io.Copy(listener, dial)
		if err != nil {
			fmt.Println(err)
		}
	}()

}

type Pokemon struct {
	ID         string   `json:"_id"`
	Index      int      `json:"index"`
	GUID       string   `json:"guid"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Picture    string   `json:"picture"`
	Age        int      `json:"age"`
	EyeColor   string   `json:"eyeColor"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Tags       []string `json:"tags"`
	Friends    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"friends"`
	Greeting      string `json:"greeting"`
	FavoriteFruit string `json:"favoriteFruit"`
}
