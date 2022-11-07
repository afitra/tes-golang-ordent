package user

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
type EmailCheckFormatter struct {
	Validate bool `json:"validate"`
}

func FormatUser(user User, token string) UserFormatter {

	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}
	return formatter
}

func CheckEmail(validate bool) EmailCheckFormatter {

	formatter := EmailCheckFormatter{
		Validate: validate,
	}

	return formatter
}
