
# ==============================================================================
# Makefile helper functions for deploy to developer env
#

KUBECTL := kubectl
NAMESPACE ?= default
#CONTEXT ?= rebirthmonkey.dev
#DEPLOYS=apiserver authz
DEPLOYS=apiserver authz

.PHONY: deploy.run.all
deploy.run.all:
	@echo "===========> Deploying all"
	@$(MAKE) deploy.run

.PHONY: deploy.run
deploy.run: $(addprefix deploy.run., $(DEPLOYS))

.PHONY: deploy.run.%
deploy.run.%:
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> Deploying $* $(VERSION)-$(ARCH)"
#    $(KUBECTL) -n $(NAMESPACE) apply -f manifests/$*.yaml set image deployment/$* $*=$(REGISTRY_PREFIX)/$*-$(ARCH):$(VERSION)
	$(KUBECTL) -n $(NAMESPACE) apply -f manifests/$*.yaml

.PHONY: deploy.clean
deploy.clean: $(addprefix deploy.clean., $(DEPLOYS))

.PHONY: deploy.clean.%
deploy.clean.%:
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> Undeploying $* $(VERSION)-$(ARCH)"
	@$(KUBECTL) -n $(NAMESPACE) delete -f manifests/$*.yaml
