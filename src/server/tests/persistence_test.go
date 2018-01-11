package vanilla

import (

	"testing"
	DB "server/persistence"
	"log"
	"os"
	"github.com/joho/godotenv"
) 

func setup(){
	log.Printf("Iniciando entorno de pruebas...")
	godotenv.Load(os.ExpandEnv("$GOPATH/src/server/.env"))
}

func tearDown(){
	log.Printf("Pruebas terminadas.")
}

func TestMain(m *testing.M) { 
    setup()
    retCode := m.Run()
    tearDown()
    os.Exit(retCode)
}

func TestOpenDB(t *testing.T) {
	
	t.Log("Testing: Open and ping database...")

	db := DB.DB()
	defer db.Close()
	
}

//CategoryPersistence

func TestAddCategoryAndDeleteIt(t *testing.T) {
	
	t.Log("Testing: AddCategory when it doesnt exist...")

	err := DB.AddCategory("TestAddCategory")

	if (err != nil){
		t.Errorf("Error while trying to add category.")
	}

	err2 := DB.RemoveCategory("TestAddCategory")

	if (err2 != nil){
		t.Errorf("Error while trying to remove category.")
	}
	
}

func TestAddCategoryWhenItAlreadyExists(t *testing.T) {

	t.Log("Testing: AddCategory when it exists...")

	DB.AddCategory("TestAddCategory")

	err := DB.AddCategory("TestAddCategory")

	if (err == nil){
		t.Errorf("System let me add the same category twice.")
	}

	DB.RemoveCategory("TestAddCategory")
	
}
