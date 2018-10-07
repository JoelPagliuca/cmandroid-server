#!/bin/sh

# thanks to this gist for the advice
# https://gist.github.com/willprice/e07efd73fb7f13f917ea

setup_git() {
	git config --global user.email "travis@travis-ci.org"
	git config --global user.name "Travis CI"
}

commit_docs() {
	git checkout -b docs
	git add . ./docs/*.png
	git commit --message "Travis build: $TRAVIS_BUILD_NUMBER"
}

upload_files() {
	git remote add origin-docs https://${GH_TOKEN}@github.com/JoelPagliuca/cmandroid-server > /dev/null 2>&1
	git push --quiet --set-upstream origin-docs docs
}