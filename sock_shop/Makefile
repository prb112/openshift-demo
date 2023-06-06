REGISTRY=quay.io/pbastide_rh/openshift-demo

cross-build-front-end-linux-amd64:
	+@echo test
	+@podman build --platform linux/amd64 -t ${REGISTRY}:user-amd64 -f automation/Dockerfile-front-end
.PHONY: cross-build-front-end-linux-amd64

cross-build-linux-arm64:
	+@podman build --platform linux/arm64 -t ${REGISTRY}:user-arm64 -f automation/Dockerfile-user
	+@podman build --platform linux/arm64 -t ${REGISTRY}:front-end-arm64 -f automation/Dockerfile-front-end
	+@podman build --platform linux/arm64 -t ${REGISTRY}:payment-arm64 -f automation/Dockerfile-payment
	+@podman build --platform linux/arm64 -t ${REGISTRY}:orders-arm64 -f automation/Dockerfile-orders
	+@podman build --platform linux/arm64 -t ${REGISTRY}:catalogue-arm64 -f automation/Dockerfile-catalogue
	+@podman build --platform linux/arm64 -t ${REGISTRY}:catalogue-db-arm64 -f automation/Dockerfile-catalogue-db
	+@podman build --platform linux/arm64 -t ${REGISTRY}:carts-arm64 -f automation/Dockerfile-carts
	+@podman build --platform linux/arm64 -t ${REGISTRY}:shipping-arm64 -f automation/Dockerfile-shipping
	+@podman build --platform linux/arm64 -t ${REGISTRY}:queue-master-arm64 -f automation/Dockerfile-queue-master

.PHONY: cross-build-linux-arm64





cross-build-linux-amd64:
	+@podman build --platform linux/amd64 -t ${REGISTRY}:user-amd64 -f automation/Dockerfile-user
.PHONY: cross-build-linux-amd64

cross-build-linux-ppc64le:
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:user-ppc64le -f automation/Dockerfile-user
.PHONY: cross-build-linux-ppc64le

cross-build: cross-build-linux-amd64 cross-build-linux-ppc64le
.PHONY: cross-build