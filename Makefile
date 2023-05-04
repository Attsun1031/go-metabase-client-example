.PHONY: generate
generate:
	openapi-generator generate -i resources/openapi.yaml -g go -o ./ --package-name gometabase --git-user-id Attsun1031 --git-repo-id go-metabase-client-example
