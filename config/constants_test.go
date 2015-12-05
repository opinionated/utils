package config_test
import (
	"github.com/opinionated/utils/config"
	"fmt"
	"testing"
	"os"


) //test if reading constants work
func TestReadConst(t *testing.T) {
 constant := config.Constants{} 		//struct to pull in the constants
 file, err := os.Open("test_const.json")		// open file and error checking
 if err != nil {
 	t.Errorf("oh nsoe, err reading json file")
 	return
 }
 defer file.Close()
 err = config.LoadConst(file, &constant)		//loads all the constants from the file into constant
 if err != nil{
 	t.Errorf("Decode = :(")						//error checking
 }
 fmt.Println("constants are: ", constant)


}