package helper

type MailList struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	From       string     `json:"from"`
	To         string     `json:"to"`
	Message    string     `json:"msg"`
	Attachment Attachment `json:"attachment"`
}

type Attachment struct {
	File  string `json:"file"`
	Link  string `json:"link"`
	Video string `json:"video"`
}

type MailCreate struct {
	Title   string `json:"title"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"msg"`
}

type MailUpdate struct {
	Title   string `json:"title"`
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"msg"`
	Type    string `json:"type"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type UserLogin struct {
	Token string `json:"token"`
}

type ResponseModule struct {
	Package interface{} `json:"package"`
	Message string      `json:"message"`
}

type Image struct {
	File string `json:"file"`
	Link string `json:"link"`
}

type File struct {
	File string `json:"file"`
	Link string `json:"link"`
}
