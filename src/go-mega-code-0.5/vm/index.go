package vm

import "go-mega-code-0.5/model"

// IndexViewModel struct
type IndexViewModel struct {
    BaseViewModel
    model.User
    Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
    u1, _ := model.GetUserByUsername("rene")
    posts, _ := model.GetPostsByUserID(u1.ID)
    v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
    return v
}
