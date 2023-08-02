package uids

type IRepository interface {
	GetHexString(digits int) string
}
