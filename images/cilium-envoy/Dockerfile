# syntax=docker/dockerfile:1.1-experimental

# Copyright 2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG COMPILERS_IMAGE=docker.io/cilium/image-compilers:07ffd7516e5af3a4b369f676b893be6305ad75f6

FROM --platform=linux/amd64 ${COMPILERS_IMAGE} as builder
ARG TARGETPLATFORM

ARG BAZEL_BUILD_EXTRA_OPTS
ARG BAZEL_BUILD_OPTS="--jobs=auto --experimental_local_memory_estimate --verbose_failures --spawn_strategy=local ${BAZEL_BUILD_EXTRA_OPTS}"

WORKDIR /src

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json \
  make cilium-envoy install \
  BAZEL_BUILD_OPTS="${BAZEL_BUILD_OPTS}" \
  PKG_BUILD=1 \
  DESTDIR=/out/linux/amd64 || sleep 15000

RUN --mount=type=bind,readwrite,target=/src --mount=type=secret,id=gcp-service-account-key,target=/src/gcp-service-account-key.json \
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
