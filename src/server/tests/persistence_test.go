package vanilla

import (

	"testing"
	DB "server/persistence"
	M "server/models"
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

	cat := M.NewCategory("PruebaAdd")
	_ , err := DB.AddCategory(cat)

	if (err != nil){
		t.Errorf("Error while trying to add category.")
	}

	err2 := DB.RemoveCategory(cat)

	if (err2 != nil){
		t.Errorf("Error while trying to remove category.")
	}
	
}

func TestAddCategoryWhenItAlreadyExists(t *testing.T) {

	t.Log("Testing: AddCategory when it exists...")

	cat := M.NewCategory("TestAddCategory")
	DB.AddCategory(cat)

	_ , err := DB.AddCategory(cat)

	if (err == nil){
		t.Errorf("System let me add the same category twice.")
	}

	DB.RemoveCategory(cat)
	
}

func TestUpdateCategory(t *testing.T) {
	
	t.Log("Testing: AddCategory when it doesnt exist...")

	cat := M.NewCategory("xD")
	_ , err := DB.AddCategory(cat)

	if (err != nil){
		t.Errorf("Error while trying to add category.")
	}

	err2 := DB.UpdateCategory(cat, "Caca")

	if (err2 != nil){
		log.Print(err2.Error())
		t.Errorf("Error while trying to update category.")
	}

	_ , errDE := DB.GetCategoryByName("Caca")

	if (errDE != nil){
		t.Errorf("Error, category not updated")
	}

	DB.RemoveCategory(cat)

}

func TestGetAllCategories(t *testing.T) {
	
	t.Log("Testing: AddCategory when it doesnt exist...")

	cat := M.NewCategory("PruebaAdd")
	_ , err := DB.AddCategory(cat)

	if (err != nil){
		t.Errorf("Error while trying to add category.")
	}

	err2 := DB.RemoveCategory(cat)

	if (err2 != nil){
		t.Errorf("Error while trying to remove category.")
	}
	
}
