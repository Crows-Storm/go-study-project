package counters

// AlertCounter is an exported named type that contains an integer counter for alerts.
// The first character is in upper-case format so it is considered to be exported.
type AlertCounter int

// alertCounter is an unexported named type that contains an integer counter for alerts.
// The first character is in lower-case format so it is considered to be unexported.
// It is not accessible for other packages, unless they are part of the package counters themselves.
type alertCounter int

// Declare an exported function called New - a factory function that knows how to create and
// initialize the value of an unexported type.
// It returns an unexported value of alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}

// The compiler is okay with this because exporting and unexporting is not about the value like
// private and public mechanism, it is about the identifier itself.
// However, we don't do this since there is no encapsulation here. We can just make the type
// exported.

// Exported type User represents information about a user.
// It has 2 exported fields: Name and ID and 1 unexported field: password.
type User struct {
	Name string
	ID   int

	password string
}

func InitPwd(u *User) {
	u.password = "123456"
}

type student struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
// Exported type embedded the unexported field user.
type Manager struct {
	Title string

	student
}
