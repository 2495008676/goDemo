package model

// User struct
type User struct {
    id           int
    userName     string
    userPassword string
}

func (u *User) setPassword(userPassword string) {
    u.userPassword = userPassword;
}

func (u *User) getUserName(userName string) (*User, error) {
    var user User
    if err := db.Where("username=?", username).Find(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
