SUBDIRS := patientservice patientdetailsservice
.PHONY: all $(SUBDIRS)

target = $(filter-out $@,$(MAKECMDGOALS))
$(SUBDIRS):
	$(MAKE) -C $@ $(target)

.PHONY: test build deploy-local undeploy-local $(SUBDIRS) 

.SILENT: test duild deploy-local undeploy-local

test build deploy-local undeploy-local view-local: 
	@echo # do nothing

	