NAME=template
VERSION=0.0.0

TARGET_OS=linux windows darwin
TARGET_ARCH=amd64 arm64 386 arm

DOCKER_PREFIX=

SHELL='bash'

run:
	@echo 'Starting backend...'
	@go run src/backend/cmd/main.go

front:
	@echo 'Starting frontend'
	@cd ./src/frontend && yarn && yarn dev

build: build_front build_back
	
build_front: create_build_dir
	@echo "$$(tput setaf 4)Preparing frontend...$$(tput sgr0)"
	@cd ./src/frontend && yarn install > /dev/null
	@echo "$$(tput setaf 2)Creating frontend bundle...$$(tput sgr0)"
	@cd ./src/frontend && yarn build --outDir ../../build/$(VERSION)/frontend --emptyOutDir > /dev/null

build_back: create_build_dir
	@for os in $(TARGET_OS) ; do \
	for arch in $(TARGET_ARCH) ; do \
	if [ $$os = "windows" ] ; then \
		ending='.exe' ; \
	else \
		ending='' ; \
	fi ;\
	echo "$$(tput setaf 2)Building $$arch/$$os binary...$$(tput sgr0)" ;\
	GOOS=$$os GOARCH=$$arch go build -o build/$(VERSION)/$$arch/$$os/$(NAME)$$ending src/backend/cmd/main.go || true; \
	done ; \
	done ; \

major: MAJOR_INC=1
major: PATCH_INC=-${PATCH} 
major: MINOR_INC=-${MINOR}
major: release
minor: MINOR_INC=1 
minor: PATCH_INC=-${PATCH}
minor: release
patch: PATCH_INC=1
patch: release

PATCH_INC=0
MINOR_INC=0
MAJOR_INC=0
PATCH=$(shell echo ${VERSION} | cut -d'.' -f3 )
MINOR=$(shell echo ${VERSION} | cut -d'.' -f2 )
MAJOR=$(shell echo ${VERSION} | cut -d'.' -f1 )

bump:
	@echo "$$(tput setaf 3)Bumping version...$$(tput sgr0)"
	$(eval FULLNEW=$(shell echo `expr ${MAJOR} + ${MAJOR_INC}`.`expr ${MINOR} + ${MINOR_INC}`.`expr ${PATCH} + ${PATCH_INC}`))
	@sed -i 's/VERSION=${VERSION}/VERSION=${FULLNEW}/' Makefile
	$(eval VERSION=${FULLNEW})	

release: bump build dist git_full docker docker_push gh_release clean
	@echo "$$(tput setaf 2)Version ${VERSION} is ready$$(tput sgr0)"

git_full: git_commit git_tag

git_commit:
	@git add Makefile
	@git commit -m "Release v${VERSION}"
	@git push

git_tag:
	@git tag -a v${VERSION} -m "Release v${VERSION}"
	@git push origin v${VERSION}

dist: clean_dist
	@{ \
	mkdir -p dist ;\
	echo "$$(tput setaf 4)Creating distributive...$$(tput sgr0)" ;\
	for a in ${TARGET_ARCH}; do \
	pushd build/${VERSION}/$$a > /dev/null ;\
			for o in $$(ls); do \
				name=${NAME}-${VERSION}-$$o-$$a.zip ;\
				echo "$$(tput setaf 2)$$name$$(tput sgr0)" ;\
				(cd ../../.. && zip -j dist/$$name build/${VERSION}/$$a/$$o/${NAME}* LICENSE README.md > /dev/null) ;\
				(cd ../ && zip -r ../../dist/$$name frontend > /dev/null) ;\
			done ; \
	popd > /dev/null ;\
	done; \
	echo "$$(tput setaf 5)Distributive created$$(tput sgr0)" ;\
	}

gh_release:
	@{\
		if [ $$(which gh) ]; then \
			gh release create v${VERSION} --notes "" || true;\
			gh release upload v${VERSION} dist/*.zip || true;\
		else \
			echo "$$(tput setaf 1)Github CLI tool is not installed - skiping artifact upload$$(tput sgr0)" ;\
		fi;\
	}

dock:
	@docker run ${DOCKER_PREFIX}${NAME}

docker: build_front
	@echo "$$(tput setaf 4)Dockerizing...$$(tput sgr0)"
	@docker build --build-arg version=${VERSION} -t ${DOCKER_PREFIX}${NAME}:${VERSION} . > /dev/null
	@docker tag ${DOCKER_PREFIX}${NAME}:${VERSION} ${DOCKER_PREFIX}${NAME}:latest > /dev/null
	@echo "$$(tput setaf 5)Built ${DOCKER_PREFIX}${NAME}:${VERSION} image$$(tput sgr0)"

docker_push:
	@{\
		if [ ${DOCKER_PREFIX} ]; then \
			docker push ${DOCKER_PREFIX}${NAME}:${VERSION} || true ;\
			docker push ${DOCKER_PREFIX}${NAME} || true ;\
		else\
			echo "$$(tput setaf 1)DOCKER_PREFIX is not set - skiping docker registry push$$(tput sgr0)" ;\
		fi;\
	}

create_build_dir:
	@mkdir -p build

clean: clean_dist
	@rm -rf build

clean_dist:
	@rm -rf dist