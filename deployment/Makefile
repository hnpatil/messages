.PHONY: deploy-local
deploy-local:
	kubectl apply -f local/namespace.yaml
	kubectl apply -f local/database.yaml
	kubectl apply -f local/config.yaml
	kubectl apply -f local/api.yaml

.PHONY: port-forward
port-forward:
	kubectl port-forward service/api-server 8080:8080 -n messages

