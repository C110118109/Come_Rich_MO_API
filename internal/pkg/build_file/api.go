package build_file

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Api(outputName string) (filePath string) {

	url := "https://api.pspdfkit.com/build"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("instructions", "{\n  \"parts\": [\n    {\n      \"file\": \"document\"\n    }\n  ]\n}\n")
	file, errFile2 := os.Open("storage/" + outputName + ".xlsx")
	defer file.Close()
	part2,
		errFile2 := writer.CreateFormFile("document", filepath.Base(outputName+".xlsx"))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer pdf_live_bxnx3bbyLOOxkLW59E8yspzKcX1sa0go7J2tWGD5Dgh")
	req.Header.Add("Cookie", "AWSALB=UhKGAVWom3WZ+OUK9PTUqEzj23heYDjsFjUPQ3hM+16A3dKCh4Qlg9ELAKP3GziXbXYl1xFaK5RB2vKsqff0xP0eGsWsrM+lo5V486AgnX4qp8G/PRXdYpNy7yDz; AWSALBCORS=UhKGAVWom3WZ+OUK9PTUqEzj23heYDjsFjUPQ3hM+16A3dKCh4Qlg9ELAKP3GziXbXYl1xFaK5RB2vKsqff0xP0eGsWsrM+lo5V486AgnX4qp8G/PRXdYpNy7yDz")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	output, err := os.Create("storage/" + outputName + ".pdf")
	if err != nil {
		// handle error
	}
	defer output.Close()

	_, err = io.Copy(output, res.Body)
	if err != nil {
		// handle error
	}
	return "http://127.0.0.1:8090/public/" + outputName + ".pdf"
}
