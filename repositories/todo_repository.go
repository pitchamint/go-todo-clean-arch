package repositories

//adapter
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // use to connect db
	"github.com/jinzhu/gorm"
	"github.com/krittawatcode/go-todo-clean-arch/domains"
	"github.com/krittawatcode/go-todo-clean-arch/models"
	"github.com/krittawatcode/go-todo-clean-arch/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type todoRepository struct {
	conn *gorm.DB
}

// NewToDoRepository ...
func NewToDoRepository(conn *gorm.DB) domains.ToDoRepository {
	return &todoRepository{conn}
}

func (t *todoRepository) GetAllTodo(todo *[]models.Todo) (err error) {
	//ไม่ใส่ตัว log error เพราะว่าไม่ใช่หน้าที่ของ repo ที่ต้องจัดการเป็นส่วนของ business logi
	if err = t.conn.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) CreateATodo(todo *models.Todo) (err error) {
	if err = t.conn.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) GetATodo(todo *models.Todo, id string) (err error) {
	if err := t.conn.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) UpdateATodo(todo *models.Todo, id string) (err error) {
	// fmt.Println(todo)
	t.conn.Save(todo) // save all field
	return nil
}

func (t *todoRepository) DeleteATodo(todo *models.Todo, id string) (err error) {
	t.conn.Where("id = ?", id).Delete(todo)
	return nil
}
func (t *todoRepository) CreateUserTodo(login *models.LoginUser) (err error) {
	if err = t.conn.Create(login).Error; err != nil {
		return err
	}
	return nil
}
func (t *todoRepository) LoginTodo(login *models.Login) (err error) {
	user := new(models.LoginUser)
	if err = t.conn.Where("email = ?", login.Email).Find(&user).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return err
	}
	token, exp, err := jwt.CreateJWTToken(*user)
	if err != nil {
		return err
	}
	fmt.Println("token", token)
	fmt.Println("exp", exp)
	return nil
}
