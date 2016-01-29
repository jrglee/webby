def registry = 'registry.revinate.net/techops'
def image = 'webby'
def name = "${registry}/${image}"

stage 'Golang build'
node {
  checkout scm

  sh "docker run --rm -v `pwd`:/usr/src/${image} -w /usr/src/${image} golang:latest make ${image}"

sh 'pwd'
sh 'ls -al'

  stash name: 'binary', includes: 'build/webby'
}

stage 'Docker build and push'
node {
  checkout scm

  unstash 'binary'

  removeRunningContainers()

  try {
    sh 'docker build .'
  } catch (all) {
    echo 'Output from app container'
    sh 'docker logs crawler'

    removeRunningContainers()
    throw all
  }

  sh "git log -1 | head -1 | cut -c 8-10 > git-revision"
  def revision = readFile 'git-revision'
  def version = "1.0.${env.BUILD_NUMBER ?: 0}.${revision ?: 1}"

  sh "docker tag -f ${image} ${name}"
  sh "docker push ${name}"

  def tag = "${name}:${version}"
  sh "docker tag -f ${image} ${tag}"
  sh "docker push ${tag}"
}

def removeRunningContainers() {
  echo 'Remove cleaning containers'
  try {
    sh 'docker rm -f -v `docker ps -aq`'
  } catch(all) {
    echo 'No containers to remove'
  }
}
