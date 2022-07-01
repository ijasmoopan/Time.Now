package user

// import (
// 	"errors"

// 	"github.com/ijasmoopan/Time.Now/models"
// 	"github.com/ijasmoopan/Time.Now/repository"
// )

// func ValidateUser(userForm models.User) (models.User, error){

// 	var user models.User
	
// 	db := repository.ConnectDB()
// 	defer repository.CloseDB(db)

// 	if result := db.First(&user, "user_email = ? ", userForm.User_email); result.Error != nil {
// 		return user, result.Error
// 	} 
// 	if user.User_password != userForm.User_password {
// 		return user, errors.New("Incorrect Password")
// 	}
// 	return user, nil
// }

// func RegisteringUser(newUser models.User) (models.User, error){

// 	db := repository.ConnectDB()
// 	defer repository.CloseDB(db)

// 	if result := db.Create(&newUser); result.Error != nil {
// 		return newUser, result.Error
// 	}
// 	return newUser, nil
// }