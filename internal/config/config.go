package config

var SecretKey = []byte("default-secret-key")

var EmailServerAuth = struct {
	SMTPServer string
	Username   string
	Password   string
}{
	"smtp.gmail.com",
	"resourcepack.notifier@gmail.com",
	"resourcepack1",
}

var FirstUser = struct {
	Usernane string
	Password string
}{
	"admin",
	"root",
}
