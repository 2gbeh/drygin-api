package faker_util

type FakePerson struct {
	FirstName string  `faker:"first_name"`
	LastName  string  `faker:"last_name"`
	Gender    string  `faker:"gender"`
	BirthDate string  `faker:"date"`
	Email     string  `faker:"email"`
	Phone     string  `faker:"phone_number"`
	Bio       string  `faker:"sentence"`
	Wallet    float64 `faker:"amount"`
	IP        string  `faker:"ipv4"`
	Latitude  float32 `faker:"lat"`
	Longitude float32 `faker:"long"`
	CreatedAt string  `faker:"timestamp"`
	UUID      string  `faker:"uuid_hyphenated"`
}