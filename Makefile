docker/build:
	docker build -t cperd .

docker/run:
	docker run --rm -it --entrypoint bash -e RUN_LOCAL=true -e GITHUB_WORKSPACE=/cperd -v $(pwd):/cperd cperd