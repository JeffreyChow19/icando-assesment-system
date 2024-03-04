FROM ubuntu:20.04

ENV REGISTRATION_TOKEN=glrt-yxCKaYMyN_2i1E9HxeYs

RUN apt-get update && \
    apt-get install -y curl git && \
    curl -L https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh | bash && \
    apt-get install -y gitlab-runner

RUN gitlab-runner register \
  --non-interactive \
  --executor "docker" \
  --docker-image alpine:latest \
  --url https://gitlab.informatika.org/ \
  --registration-token $REGISTRATION_TOKEN \
  --description "Docker Runner IF3250_2023_K01_07" \
  --maintenance-note "Rebuild the runner if failed" \
  --tag-list docker-runner-k01-07 \
  --run-untagged="true" \
  --locked="true" \
  --access-level="not_protected"

CMD ["gitlab-runner", "run"]
