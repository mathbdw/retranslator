package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/mathbdw/retranslator/internal/app/repo EventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/mathbdw/retranslator/internal/app/sender EventSender
