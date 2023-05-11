package user

type LoginFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func FormatUser(user User) LoginFormatter {
	formatter := LoginFormatter{}
	formatter.Name = user.Name
	formatter.Email = user.Email
	formatter.Role = user.Role

	return formatter
}
