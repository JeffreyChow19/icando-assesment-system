build-runner-image:
	docker build -t k01-07-runner .

run-runner:
	docker run -d --name docker-runner-k01-07 k01-07-runner
