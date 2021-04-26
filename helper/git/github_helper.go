package git

type Github interface {
	GetContentYaml(owner string, repo string, path string) (contentYaml []byte, err error)
}
