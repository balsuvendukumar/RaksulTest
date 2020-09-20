package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	/* As the Purpose is not clear of the code, I beleive ,the below code can be used to find Interactions of a particular user
	from multiple files based on date. Thus instead of providing the word during while
	executing the main function, we can provide the word in the Console Input*/
	//Provide Input for User ID to be flagged
	var flagie string
	fmt.Scanf("Provide the User Name to be Seached in log files: %s", &flagie)
	word := flag.String("word", "", "a word")
	flag.Parse()
	//Open the Zip Folder which contains all log files.
	r, err := zip.OpenReader("foc-slack-export.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close() //do nt forget to close the file

	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			if strings.Index(f.Name, ".json") == len(f.Name)-5 {
				rc, err := f.Open() //Open each json file
				if err != nil {
					panic(err)
				}

				b, err := ioutil.ReadAll(rc) //read the content
				if err != nil {
					panic(err)
				}
				rc.Close()

				var all []map[string]interface{}
				err = json.Unmarshal(b, &all) //save the JSON formatted data in Map:all
				if err != nil {
					panic(err)
				}

				for _, m := range all {
					//validate if the required word is present against text
					if text, ok := m["text"].(string); ok {
						//if yes, Print the File name along with the Test Message
						if strings.Contains(text, *word) {
							fmt.Printf("File %s : %s\n ", f.Name, text)
						}
					}
				}
			}
		}
	}
}
