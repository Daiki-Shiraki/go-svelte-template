# syntax=docker/dockerfile:1

################################################################################
# build stage
ARG GO_VERSION=1.23
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /opt/app


# Dockerのキャッシュを利用するために、依存関係を別ステップとしてダウンロードする。
# 以降のビルドを高速化するために、/go/pkg/mod/へのキャッシュマウントを活用する。
# コンテナにコピーする手間を省くために、go.sumとgo.modへのバインドマウントを活用する。
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=/src/go.sum,target=go.sum \
    --mount=type=bind,source=/src/go.mod,target=go.mod \
    go mod download -x 

# これはビルダーから渡されるビルド対象のアーキテクチャです。
# ここに置くことで、前の手順をアーキテクチャ間でキャッシュできるようになります。
ARG TARGETARCH

# アプリケーションをビルドする。
# 以降のビルドを高速化するために、 /go/pkg/mod/へのキャッシュマウントを活用する。
# ソースコードをコンテナにコピーする手間を省くために、カレントディレクトリへのバインドマウントを活用する。
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=/src,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server .

################################################################################
# dev stage(with air)
# Final Stage (Development with Air)
FROM build AS dev


# ビルドステージから実行ファイルをコピーする。
COPY --from=build /bin/server /bin/

RUN go install github.com/air-verse/air@latest

# Copy source code for development
WORKDIR /opt/app
COPY . .

# Air configuration
EXPOSE 3000
ENTRYPOINT ["air", "-c", ".air.toml"]

################################################################################
# final stage

FROM alpine:latest AS final

# アプリケーションの実行に必要なランタイムの依存関係をインストール。
# 以降のビルドを高速化するために、/var/cache/apk/へのキャッシュマウントを活用。
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

# todo これいる？
# アプリを実行する非特権ユーザーを作成します。
# https://docs.docker.com/go/dockerfile-user-best-practices/ を参照。
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# ビルドステージから実行ファイルをコピーする。
COPY --from=build /bin/server /bin/

# ポートを公開する。
EXPOSE 3000

# コンテナ起動時に実行する内容。
ENTRYPOINT [ "/bin/server" ]
