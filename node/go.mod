module github.com/osyah/go-pletyvo/node

go 1.21.1

replace (
	github.com/osyah/go-pletyvo => ../.
	github.com/osyah/go-pletyvo/protocol/dapp/crypto => ../protocol/dapp/crypto
	github.com/osyah/go-pletyvo/store/pebble => ../store/pebble
)

require (
	github.com/gofiber/fiber/v2 v2.52.6
	github.com/google/uuid v1.6.0
	github.com/osyah/go-pletyvo v0.0.0-00010101000000-000000000000
	github.com/osyah/go-pletyvo/protocol/dapp/crypto v0.0.0-00010101000000-000000000000
	github.com/osyah/go-pletyvo/store/pebble v0.0.0-00010101000000-000000000000
	github.com/osyah/hryzun v0.0.2
	github.com/osyah/hryzun/config v0.0.0-20240910165922-f61b828b2a1e
	github.com/rs/zerolog v1.33.0
)

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/VictoriaMetrics/easyproto v0.1.4 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240606204812-0bbfbd93a7ce // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v1.1.4 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.12.0 // indirect
	github.com/prometheus/client_model v0.2.1-0.20210607210712-147c58e9608a // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.58.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/zeebo/blake3 v0.2.4 // indirect
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
