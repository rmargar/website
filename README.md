
[![codecov](https://codecov.io/gh/rmargar/website/branch/main/graph/badge.svg)](https://codecov.io/gh/rmargar/website) ![build](https://github.com/rmargar/website/actions/workflows/deploy.yaml/badge.svg)
# rmargar.net

This repository contains the source code for my personal website: [rmargar.net](http://rmargar.net)

It is basically a Go web server that serves static files with my personal dev portolio and a blog section with posts that are stored in a PostgreSQL database. The posts are served as HTML templates and are rendered using Markdown, which allows for easy formatting. It is a kinda over-engineered way of hosting a blog site, but I used it to learn more about the Go language :)

## Development

1. Run `make build` to build from source.
2. Run the binary from `./bin/server`, or `make run`

## Testing

Run `make test` to run both unit and integration tests. Make sure to have docker running as the integration tests use a PostgreSQL container.

## Deployment

Currently I am hosting the website using [CapRover](https://caprover.com/), which makes really easy to deploy containerized applications on any cloud server or similar. A new build of the app is currently triggered after merge to main branch.

## Dependencies

- Golang 1.19 or higher
- PostgreSQL
- To build the blog, a [Disqus](https://disqus.com/) account is needed.
