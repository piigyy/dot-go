package repository

var (
	QueryCreateUser = `INSERT INTO public."user"
	(fullname, email, birth_date)
	VALUES($1, $2, $3) RETUNING "id";`
)
