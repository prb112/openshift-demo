REGISTRY=quay.io/pbastide_rh/openshift-demo

# ARM64
cross-build-user-linux-arm64:
	+@echo "Building Image - 'user'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:user-arm64 -f automation/Dockerfile-user
	+@echo "Done Image - 'user'"
.PHONY: cross-build-user-linux-arm64

cross-build-front-end-linux-arm64:
	+@echo "Building Image - 'front-end'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:front-end-arm64 -f automation/Dockerfile-front-end
	+@echo "Done Image - 'front-end'"
.PHONY: cross-build-front-end-linux-arm64

cross-build-payment-linux-arm64:
	+@echo "Building Image - 'payment'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:payment-arm64 -f automation/Dockerfile-payment
	+@echo "Done Image - 'payment'"
.PHONY: cross-build-payment-linux-arm64

cross-build-orders-linux-arm64:
	+@echo "Building Image - 'orders'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:orders-arm64 -f automation/Dockerfile-orders
	+@echo "Done Image - 'orders'"
.PHONY: cross-build-orders-linux-arm64

cross-build-catalogue-linux-arm64:
	+@echo "Building Image - 'catalogue'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:catalogue-arm64 -f automation/Dockerfile-catalogue
	+@echo "Done Image - 'catalogue'"
.PHONY: cross-build-catalogue-linux-arm64

cross-build-catalogue-db-linux-arm64:
	+@echo "Building Image - 'catalogue-db'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:catalogue-db-arm64 -f automation/Dockerfile-catalogue-db
	+@echo "Done Image - 'catalogue-db'"
.PHONY: cross-build-catalogue-db-linux-arm64

cross-build-carts-linux-arm64:
	+@echo "Building Image - 'carts'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:carts-arm64 -f automation/Dockerfile-carts
	+@echo "Done Image - 'carts'"
.PHONY: cross-build-carts-linux-arm64

cross-build-shipping-linux-arm64:
	+@echo "Building Image - 'shipping'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:shipping-arm64 -f automation/Dockerfile-shipping
	+@echo "Done Image - 'shipping'"
.PHONY: cross-build-shipping-linux-arm64

cross-build-queue-master-linux-arm64:
	+@echo "Building Image - 'queue-master'"
	+@podman build --platform linux/arm64 -t ${REGISTRY}:queue-master-arm64 -f automation/Dockerfile-queue-master
	+@echo "Done Image - 'queue-master'"
.PHONY: cross-build-queue-master-linux-arm64

cross-build-linux-arm64: cross-build-user-linux-arm64 cross-build-front-end-linux-arm64 cross-build-payment-linux-arm64 cross-build-orders-linux-arm64 cross-build-catalogue-linux-arm64 cross-build-catalogue-db cross-build-carts-linux-arm64 cross-build-shipping-linux-arm64 
.PHONY: cross-build-linux-arm64

# s390x

cross-build-user-linux-s390x:
	+@echo "Building Image - 'user'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:user-s390x -f automation/Dockerfile-user
	+@echo "Done Image - 'user'"
.PHONY: cross-build-user-linux-s390x

cross-build-front-end-linux-s390x:
	+@echo "Building Image - 'front-end'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:front-end-s390x -f automation/Dockerfile-front-end
	+@echo "Done Image - 'front-end'"
.PHONY: cross-build-front-end-linux-s390x

cross-build-payment-linux-s390x:
	+@echo "Building Image - 'payment'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:payment-s390x -f automation/Dockerfile-payment
	+@echo "Done Image - 'payment'"
.PHONY: cross-build-payment-linux-s390x

cross-build-orders-linux-s390x:
	+@echo "Building Image - 'orders'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:orders-s390x -f automation/Dockerfile-orders
	+@echo "Done Image - 'orders'"
.PHONY: cross-build-orders-linux-s390x

cross-build-catalogue-linux-s390x:
	+@echo "Building Image - 'catalogue'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:catalogue-s390x -f automation/Dockerfile-catalogue
	+@echo "Done Image - 'catalogue'"
.PHONY: cross-build-catalogue-linux-s390x

cross-build-catalogue-db-linux-s390x:
	+@echo "Building Image - 'catalogue-db'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:catalogue-db-s390x -f automation/Dockerfile-catalogue-db
	+@echo "Done Image - 'catalogue-db'"
.PHONY: cross-build-catalogue-db-linux-s390x

cross-build-carts-linux-s390x:
	+@echo "Building Image - 'carts'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:carts-s390x -f automation/Dockerfile-carts
	+@echo "Done Image - 'carts'"
.PHONY: cross-build-carts-linux-s390x

cross-build-shipping-linux-s390x:
	+@echo "Building Image - 'shipping'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:shipping-s390x -f automation/Dockerfile-shipping
	+@echo "Done Image - 'shipping'"
.PHONY: cross-build-shipping-linux-s390x

cross-build-queue-master-linux-s390x:
	+@echo "Building Image - 'queue-master'"
	+@podman build --platform linux/s390x -t ${REGISTRY}:queue-master-s390x -f automation/Dockerfile-queue-master
	+@echo "Done Image - 'queue-master'"
.PHONY: cross-build-queue-master-linux-s390x

cross-build-linux-s390x: cross-build-user-linux-s390x cross-build-front-end-linux-s390x cross-build-payment-linux-s390x cross-build-orders-linux-s390x cross-build-catalogue-linux-s390x cross-build-catalogue-db cross-build-carts-linux-s390x cross-build-shipping-linux-s390x 
.PHONY: cross-build-linux-s390x

# amd64

cross-build-user-linux-amd64:
	+@echo "Building Image - 'user'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:user-amd64 -f automation/Dockerfile-user
	+@echo "Done Image - 'user'"
.PHONY: cross-build-user-linux-amd64

cross-build-front-end-linux-amd64:
	+@echo "Building Image - 'front-end'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:front-end-amd64 -f automation/Dockerfile-front-end
	+@echo "Done Image - 'front-end'"
.PHONY: cross-build-front-end-linux-amd64

cross-build-payment-linux-amd64:
	+@echo "Building Image - 'payment'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:payment-amd64 -f automation/Dockerfile-payment
	+@echo "Done Image - 'payment'"
.PHONY: cross-build-payment-linux-amd64

cross-build-orders-linux-amd64:
	+@echo "Building Image - 'orders'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:orders-amd64 -f automation/Dockerfile-orders
	+@echo "Done Image - 'orders'"
.PHONY: cross-build-orders-linux-amd64

cross-build-catalogue-linux-amd64:
	+@echo "Building Image - 'catalogue'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:catalogue-amd64 -f automation/Dockerfile-catalogue
	+@echo "Done Image - 'catalogue'"
.PHONY: cross-build-catalogue-linux-amd64

cross-build-catalogue-db-linux-amd64:
	+@echo "Building Image - 'catalogue-db'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:catalogue-db-amd64 -f automation/Dockerfile-catalogue-db
	+@echo "Done Image - 'catalogue-db'"
.PHONY: cross-build-catalogue-db-linux-amd64

cross-build-carts-linux-amd64:
	+@echo "Building Image - 'carts'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:carts-amd64 -f automation/Dockerfile-carts
	+@echo "Done Image - 'carts'"
.PHONY: cross-build-carts-linux-amd64

cross-build-shipping-linux-amd64:
	+@echo "Building Image - 'shipping'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:shipping-amd64 -f automation/Dockerfile-shipping
	+@echo "Done Image - 'shipping'"
.PHONY: cross-build-shipping-linux-amd64

cross-build-queue-master-linux-amd64:
	+@echo "Building Image - 'queue-master'"
	+@podman build --platform linux/amd64 -t ${REGISTRY}:queue-master-amd64 -f automation/Dockerfile-queue-master
	+@echo "Done Image - 'queue-master'"
.PHONY: cross-build-queue-master-linux-amd64

cross-build-linux-amd64: cross-build-user-linux-amd64 cross-build-front-end-linux-amd64 cross-build-payment-linux-amd64 cross-build-orders-linux-amd64 cross-build-catalogue-linux-amd64 cross-build-catalogue-db cross-build-carts-linux-amd64 cross-build-shipping-linux-amd64 
.PHONY: cross-build-linux-amd64

# ppc64le


cross-build-user-linux-ppc64le:
	+@echo "Building Image - 'user'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:user-ppc64le -f automation/Dockerfile-user
	+@echo "Done Image - 'user'"
.PHONY: cross-build-user-linux-ppc64le

cross-build-front-end-linux-ppc64le:
	+@echo "Building Image - 'front-end'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:front-end-ppc64le -f automation/Dockerfile-front-end
	+@echo "Done Image - 'front-end'"
.PHONY: cross-build-front-end-linux-ppc64le

cross-build-payment-linux-ppc64le:
	+@echo "Building Image - 'payment'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:payment-ppc64le -f automation/Dockerfile-payment
	+@echo "Done Image - 'payment'"
.PHONY: cross-build-payment-linux-ppc64le

cross-build-orders-linux-ppc64le:
	+@echo "Building Image - 'orders'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:orders-ppc64le -f automation/Dockerfile-orders
	+@echo "Done Image - 'orders'"
.PHONY: cross-build-orders-linux-ppc64le

cross-build-catalogue-linux-ppc64le:
	+@echo "Building Image - 'catalogue'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:catalogue-ppc64le -f automation/Dockerfile-catalogue
	+@echo "Done Image - 'catalogue'"
.PHONY: cross-build-catalogue-linux-ppc64le

cross-build-catalogue-db-linux-ppc64le:
	+@echo "Building Image - 'catalogue-db'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:catalogue-db-ppc64le -f automation/Dockerfile-catalogue-db
	+@echo "Done Image - 'catalogue-db'"
.PHONY: cross-build-catalogue-db-linux-ppc64le

cross-build-carts-linux-ppc64le:
	+@echo "Building Image - 'carts'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:carts-ppc64le -f automation/Dockerfile-carts
	+@echo "Done Image - 'carts'"
.PHONY: cross-build-carts-linux-ppc64le

cross-build-shipping-linux-ppc64le:
	+@echo "Building Image - 'shipping'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:shipping-ppc64le -f automation/Dockerfile-shipping
	+@echo "Done Image - 'shipping'"
.PHONY: cross-build-shipping-linux-ppc64le

cross-build-queue-master-linux-ppc64le:
	+@echo "Building Image - 'queue-master'"
	+@podman build --platform linux/ppc64le -t ${REGISTRY}:queue-master-ppc64le -f automation/Dockerfile-queue-master
	+@echo "Done Image - 'queue-master'"
.PHONY: cross-build-queue-master-linux-ppc64le

cross-build-linux-ppc64le: cross-build-user-linux-ppc64le cross-build-front-end-linux-ppc64le cross-build-payment-linux-ppc64le cross-build-orders-linux-ppc64le cross-build-catalogue-linux-ppc64le cross-build-catalogue-db cross-build-carts-linux-ppc64le cross-build-shipping-linux-ppc64le 
.PHONY: cross-build-linux-ppc64le

# cross build all arches.

cross-build: cross-build-linux-arm64 cross-build-linux-amd64 cross-build-linux-ppc64le cross-build-linux-s390x
.PHONY: cross-build