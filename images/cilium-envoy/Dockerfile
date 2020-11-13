# syntax=docker/dockerfile:1.1-experimental

# Copyright 2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG COMPILERS_IMAGE=docker.io/cilium/image-compilers:07ffd7516e5af3a4b369f676b893be6305ad75f6

FROM --platform=linux/amd64 ${COMPILERS_IMAGE} as builder

ARG BAZEL_BUILD_EXTRA_OPTS
ARG BAZEL_BUILD_OPTS="--jobs=auto --experimental_local_memory_estimate --verbose_failures --disk_cache=/tmp/bazel-cache/cas --repository_cache=/tmp/bazel-cache/repository --spawn_strategy=local --copt=-v ${BAZEL_BUILD_EXTRA_OPTS}"

WORKDIR /src

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS}" \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/amd64

# otionally cleanup cached objects from disk since the amd64 objects cannot be re-used for cross-compilation
ARG PRUNE_DISK_CACHE=false
RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache \
  test "${PRUNE_DISK_CACHE}" = "true" && ( df -h ; du -h /tmp/bazel-cache ; rm -rf /tmp/bazel-cache/cas )

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json --mount=type=cache,id=bazel-cache,target=/tmp/bazel-cache \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS} --incompatible_enable_cc_toolchain_resolution --platforms=//bazel/platforms:aarch64_cross --define=cross=aarch64" \
  STRIP=aarch64-linux-gnu-strip \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/arm64

FROM scratch
ARG TARGETPLATFORM
LABEL maintainer="maintainer@cilium.io"
COPY images/cilium-envoy/empty /tmp/empty
COPY --from=builder /out/${TARGETPLATFORM} /
