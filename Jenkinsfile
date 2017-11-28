node {
  def root = tool name: 'Go 1.8', type: 'go'
  pipeline {
    agent any
    stages {
      stage('test') {
        steps {
          withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
           sh 'go version'
          }
          
        }
      }
    }
  }
}