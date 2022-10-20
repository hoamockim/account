package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tipee/account/db/models"
	"os"
	"testing"
)

func Test_Connection(t *testing.T) {
	var user models.UserInfo
	if err := GetById(1, "user_profile", &user); err != nil {
		t.Fatal(err)
	}
	js, _ := json.Marshal(&user)
	fmt.Fprintf(os.Stdout, "%s", js)

}
