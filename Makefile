YEAR=2024

.PHONY: all
all:
	@echo "Usage: make X (where X is the day number)"

define CREATE_DAY_FOLDER
	mkdir -p $(YEAR)/day$(1); \
	touch $(YEAR)/day$(1)/main.go; \
	touch $(YEAR)/day$(1)/input.txt
endef

# Rule to dynamically handle day creation
$(shell seq 1 31): 
	@echo "Creating folder for day $@"
	$(call CREATE_DAY_FOLDER,$@)
	@echo "Created $(YEAR)/day$@ with main.go and input.txt"