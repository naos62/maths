pipeline {
    agent any
    
    stages {
      stage('test') {
        node {
          def root = tool name: 'Go 1.8', type: 'go'
          steps {
            withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
          }
          
        }
      }
    }
  }
}