module GOSICK

go 1.12

replace util => ./util

replace slackUtil => ./util/slack

require (
	github.com/YukiMiyatake/GOSICK v0.0.0-20190724221603-25478903779c
	github.com/aws/aws-sdk-go v1.23.8
	github.com/golangci/golangci-lint v1.17.1 // indirect
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/nlopes/slack v0.5.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/sys v0.0.0-20190826190057-c7b8b68b1456 // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools v0.0.0-20190826234050-71894ab67ee3 // indirect
	util v0.0.0-00010101000000-000000000000
)
