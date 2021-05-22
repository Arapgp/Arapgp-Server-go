package jwt

// JwtArab is a struct used when calling json.Unmarshal.
// used to store / send / receive some user info.
type JwtArab struct {
	Username string
	Email    string
}
