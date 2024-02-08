package unittests

import (
	octopus "backend/app"
	"backend/pkg/models"
	"testing"

	"github.com/google/uuid"
)

func TestPostObject(t *testing.T) {

	post := models.Post{
		Title:   "hello world",
		Content: "hello senegal",
		Privacy: models.PrivacyPublic,
		//ect ........
	}
	if err := post.Create(octopus.AppTest.Db.Conn, uuid.New()); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	clonedpost := models.Post{}

	if err := clonedpost.Get(octopus.AppTest.Db.Conn, post.ID); err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	results := []bool{
		post.Title == clonedpost.Title,
		post.Content == clonedpost.Content,
		post.Privacy == clonedpost.Privacy,
	}
	for _, result := range results {
		if !result {
			t.Errorf("post  created is different from user cloned by the methode Ged ")
		}
	}

	// fmt.Println("succes✅:  user created is the same as user cloned")

}
