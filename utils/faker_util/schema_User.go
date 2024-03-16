package faker_util

type FakeUser struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Email     string `faker:"email"`
	Password  string `faker:"password"`
	Role      string `faker:"oneof: Administrator, Editor, Author, Contributor, Subscriber"`
	CreatedAt string `faker:"timestamp"`
	UUID      string `faker:"uuid_hyphenated"`
}