package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"project-Deteksiwajah/models"
)

//Getdataface ambil data jumlah gambar
func Getdataface() int {
	// Open the file
	file, err := os.Open("./images/1.jpeg")

	if err != nil {
		log.Fatalln(err)
	}
	// Close the file later
	defer file.Close()

	// Buffer to store our request body as bytes
	var requestBody bytes.Buffer

	// Create a multipart writer
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Initialize the file field
	fileWriter, err := multiPartWriter.CreateFormFile("imagefile", "./images/1.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	// Copy the actual file content to the field field's writer
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatalln(err)

	}

	// Populate other fields
	// fieldWriter, err := multiPartWriter.CreateFormField("normal_field")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// _, err = fieldWriter.Write([]byte("Value"))

	if err != nil {
		log.Fatalln(err)
	}

	// We completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	multiPartWriter.Close()

	// By now our original request body should have been populated, so let's just use it with our custom request
	req, err := http.NewRequest("POST", "https://api.cloudmersive.com/image/face/locate-with-landmarks", &requestBody)

	req.Header.Set("Apikey", "406ae12f-86db-4178-9e71-d406b87af1d6")

	if err != nil {
		log.Fatalln(err)
	}
	// We need to set the content type from the writer, it includes necessary boundary as well
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	NewFaceWithLandmarks := models.FaceWithLandmarks{}

	json.NewDecoder(response.Body).Decode(&NewFaceWithLandmarks)

	// fmt.Println(NewFaceWithLandmarks.Faces[0].LeftEye[0].X)

	var Jumlahka, Jumlahki int
	for _, value := range NewFaceWithLandmarks.Faces[0].LeftEye {

		Jumlahka += value.X
	}

	for _, value := range NewFaceWithLandmarks.Faces[0].RightEye {

		Jumlahki += value.X
	}

	ratakiri := Jumlahka / 6
	ratakanan := Jumlahki / 6
	jarak := ratakanan - ratakiri
	if jarak > 300 {
		jarak = 300
	} else if jarak > 200 {

		jarak = 200
	} else if jarak > 100 {

		jarak = 100
	} else {
		jarak = 0
	}
	return jarak
}

//GetQutes untuk mengambil Qutes
func GetQutes(keyword, personvalue string) models.Personality {

	var requestBody bytes.Buffer
	req, _ := http.NewRequest("GET", "https://quote-garden.herokuapp.com/quotes/search/"+keyword, &requestBody)

	//responseData, _ := ioutil.ReadAll(response.Body)
	//defer response.Body.Close()

	NewQutes := models.Qutes{}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(response.Body).Decode(&NewQutes)

	hasil := (NewQutes.Results[0].QuoteText)
	var personal models.Personality

	personal.UserPersonality = personvalue
	personal.UserQutes = hasil

	return personal

}
