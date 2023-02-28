package pkg

var (
	OG IOpenGauss
)

func Init(shell string) {
	OG = NewOpenGauss(shell)
}
