pipeline {
    agent any
    def root = tool name: 'Go 1.8', type: 'go'
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