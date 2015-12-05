package config
				//Loads constants from a file
import (
	"encoding/json"
	"io"
)
type Constants struct{
	RssDelay	int	`json:"Rss Delay"`	//Rss Delay time
	ArticleDelay int	`json:"Article Delay"`	//task delay
	RequeueTime int	`json:"Requeue Time"`	//waitTime 
}
//function to load constants
func LoadConst(file io.Reader,v interface{}) error {
	dec := json.NewDecoder(file)   	//to read the json file correctly
 	err := dec.Decode(&v) 			// stores the values into the interface
 	if err != nil{		//error checking
 		return err
 	}
 	return nil
}