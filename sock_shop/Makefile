ARC=ppc64le   #amd64 s390x arm64 ppc64le
REGISTRY=quay.io/cbade_cs/openshift-demo

cross-build-user:
	+@echo "Building Image - 'user'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:user-${ARC} -f automation/Dockerfile-user
	+@echo "Done Image - 'user'"
.PHONY: cross-build-user

cross-build-front-end:
	+@echo "Building Image - 'front-end'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:front-end-${ARC} -f automation/Dockerfile-front-end
	+@echo "Done Image - 'front-end'"
.PHONY: cross-build-front-end

cross-build-payment:
	+@echo "Building Image - 'payment'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:payment-${ARC} -f automation/Dockerfile-payment
	+@echo "Done Image - 'payment'"
.PHONY: cross-build-payment

cross-build-orders:
	+@echo "Building Image - 'orders'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:orders-${ARC} -f automation/Dockerfile-orders
	+@echo "Done Image - 'orders'"
.PHONY: cross-build-orders

cross-build-catalogue:
	+@echo "Building Image - 'catalogue'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:catalogue-${ARC} -f automation/Dockerfile-catalogue
	+@echo "Done Image - 'catalogue'"
.PHONY: cross-build-catalogue

cross-build-catalogue-db:
	+@echo "Building Image - 'catalogue-db'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:catalogue-db-${ARC} -f automation/Dockerfile-catalogue-db
	+@echo "Done Image - 'catalogue-db'"
.PHONY: cross-build-catalogue-db

cross-build-carts:
	+@echo "Building Image - 'carts'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:carts-${ARC} -f automation/Dockerfile-carts
	+@echo "Done Image - 'carts'"
.PHONY: cross-build-carts

cross-build-shipping:
	+@echo "Building Image - 'shipping'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:shipping-${ARC} -f automation/Dockerfile-shipping
	+@echo "Done Image - 'shipping'"
.PHONY: cross-build-shipping

cross-build-queue-master:
	+@echo "Building Image - 'queue-master'"
	+@podman build --platform linux/${ARC} -t ${REGISTRY}:queue-master-${ARC} -f automation/Dockerfile-queue-master
	+@echo "Done Image - 'queue-master'"
.PHONY: cross-build-queue-master


cross-build-all: cross-build-user cross-build-front-end cross-build-payment cross-build-orders cross-build-catalogue cross-build-catalogue-db cross-build-carts cross-build-shipping cross-build-queue-master

.PHONY: cross-build-all

