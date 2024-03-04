build-runner-image:
	docker build -t k01-07-runner .

start-runner:
	docker run -d --name docker-runner-k01-07 -v /var/run/docker.sock:/var/run/docker.sock k01-07-runner

remove-runner:
	docker rm -f docker-runner-k01-07
