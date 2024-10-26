module github.com/osyah/go-pletyvo/client

go 1.21.1

replace (
	github.com/osyah/go-pletyvo => ../.
	github.com/osyah/go-pletyvo/protocol/dapp/crypto => ../protocol/dapp/crypto/.
)

require (
	github.com/google/uuid v1.6.0
	github.com/osyah/go-pletyvo v0.0.0-00010101000000-000000000000
	github.com/osyah/go-pletyvo/protocol/dapp/crypto v0.0.0-00010101000000-000000000000
	github.com/osyah/hryzun v0.0.1
)

require (
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
)
