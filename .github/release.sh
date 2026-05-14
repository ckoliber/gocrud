#!/usr/bin/env bash
set -e

function get_ref() {
  TYPE="branch"
  if [ "${GITHUB_REF_TYPE:-}" = "tag" ]; then
    TYPE="tag"
  elif [ -n "${CI_COMMIT_TAG:-}" ]; then
    TYPE="tag"
  elif git describe --tags --exact-match >/dev/null 2>&1; then
    TYPE="tag"
  fi

  if [ -n "${GITHUB_REF_NAME:-}" ]; then
    echo "$TYPE/$GITHUB_REF_NAME"
  elif [ -n "${CI_COMMIT_REF_NAME:-}" ]; then
    echo "$TYPE/$CI_COMMIT_REF_NAME"
  else
    echo "$TYPE/$(git symbolic-ref --quiet --short HEAD || git describe --tags --exact-match 2>/dev/null || git rev-parse --short HEAD)"
  fi
}

function is_release() {
  [[ "$(get_ref)" =~ ^branch/(main|master)$ ]]
}

function is_prerelease() {
  [[ "$(get_ref)" =~ ^branch/(alpha|beta|rc|next)$ ]]
}

if is_release || is_prerelease; then
  VERSION=$(git cliff --unreleased --bumped-version)
  git cliff --unreleased --tag $VERSION --output CHANGELOG.md
  gh release create "$VERSION" --title "$VERSION" --target "$(git rev-parse HEAD)" --notes-file CHANGELOG.md
else
  echo "Not a release or pre-release branch, skipping release."
fi
