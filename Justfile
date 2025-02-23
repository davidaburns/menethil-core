name := "menethil-core"
package_import := "github.com/davidaburns/menethil-core"
version := "0.0.1"
build_time := `date -u +"%Y-%m-%dT%H:%M:%SZ"`
build_number := "0"

build type="dev" platform="darwin":
	@echo "Building {{name}} ({{type}})..." ;
	@if [ "{{type}}" = "dev" ]; then \
		CGO_ENABLED=0 GOOS={{platform}} GOARCH=amd64 go build -v -ldflags "-X '{{package_import}}/build.name={{name}}-world' -X '{{package_import}}/build.buildTime={{build_time}}' -X '{{package_import}}/build.buildNumber={{build_number}}' -X '{{package_import}}/build.buildType={{type}}'" -o bin/{{name}}-world ./cmd/{{name}}-world; \
		CGO_ENABLED=0 GOOS={{platform}} GOARCH=amd64 go build -v -ldflags "-X '{{package_import}}/build.name={{name}}-auth' -X '{{package_import}}/build.buildTime={{build_time}}' -X '{{package_import}}/build.buildNumber={{build_number}}' -X '{{package_import}}/build.buildType={{type}}'" -o bin/{{name}}-auth ./cmd/{{name}}-auth; \
	elif [ "{{type}}" = "prod" ]; then \
		CGO_ENABLED=0 GOOS={{platform}} GOARCH=amd64 go build -ldflags "-X '{{package_import}}/build.name={{name}}-world' -X '{{package_import}}/build.buildTime={{build_time}}' -X '{{package_import}}/build.buildNumber={{build_number}}' -X '{{package_import}}/build.buildType={{type}}' -s -w" -trimpath -o bin/{{name}}-world ./cmd/{{name}}-world; \
		CGO_ENABLED=0 GOOS={{platform}} GOARCH=amd64 go build -ldflags "-X '{{package_import}}/build.name={{name}}-auth' -X '{{package_import}}/build.buildTime={{build_time}}' -X '{{package_import}}/build.buildNumber={{build_number}}' -X '{{package_import}}/build.buildType={{type}}' -s -w" -trimpath -o bin/{{name}}-auth ./cmd/{{name}}-auth; \
	else \
		echo "Unknown build type {{type}}, please use 'dev' or 'prod'"; \
		exit 1; \
	fi

run type="dev" platform="darwin" app="auth" +args="":
	@just build {{type}}
	@./bin/{{name}}-{{app}} {{args}}

test:
	@echo "Running Tests..."
	@go test ./...

fmt:
    @echo "Formatting files..."
    @go fmt ./...

lint:
	@echo "Linting..."
	@golangci-lint run ./...

echo-arch:
	@echo {{arch()}}

clean:
	@echo "Cleaning..."
	@rm ./{{name}}
