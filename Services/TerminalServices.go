package Services

import "fmt"

func createTerminals(terms []Terminal)  {
	for _,term := range terms{
		test := Terminal{}
		db.Where("url = ? ", fmt.Sprintf(term.Url)).First(&test)
		if (test.Url == term.Url) {
			test.Timeout 	= term.Timeout
			test.Threshold 	= term.Threshold
			test.Frequency 	= term.Frequency
			db.Save(&test)
			fmt.Println("Updated one terminal")
		}else {
			db.Save(&term)
			fmt.Println("added a new terminal")
		}
	}
}