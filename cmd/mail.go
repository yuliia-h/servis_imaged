package main

import (
	"fmt"
	"net/http"
	"projectPackage/servis_images/infrastructure"
	_ "projectPackage/servis_images/infrastructure"
	_ "projectPackage/servis_images/user_case"
)

func main() {

	http.HandleFunc("/struct/", infrastructure.HandleResizeImage)

	http.HandleFunc("/getimageId/{id}/", infrastructure.GetImageId)

	http.HandleFunc("/getimages/", infrastructure.GetImages)

	http.HandleFunc("/updateimage/{id}/", infrastructure.UpdateImage)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":45998", nil)
}
