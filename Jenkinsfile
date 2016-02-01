def registry = 'registry.revinate.net/techops'
def image = 'webby'
def name = "${registry}/${image}"

stage 'Golang build'
node {
  checkout scm

  sh "docker run --rm -v `pwd`:/usr/src/${image} -w /usr/src/${image} golang:1.5 make ${image}"

  stash name: 'binary', includes: 'webby'
}

stage 'Docker build and push'
node {
  checkout scm

  unstash 'binary'

  sh "docker build -t ${name} ."

  sh "git log -1 | head -1 | cut -c 8-10 > git-revision"
  def revision = readFile 'git-revision'
  def version = "1.0.${env.BUILD_NUMBER ?: 0}.${revision ?: 1}"

  sh "docker push ${name}"

  def tag = "${name}:${version}"
  sh "docker tag -f ${name} ${tag}"
  sh "docker push ${tag}"
}

