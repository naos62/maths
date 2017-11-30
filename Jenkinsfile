pipeline {
  agent none
  stages {
    stage('') {
     agent {
      docker {
          image 'golang'
      } 
     steps {
        sh 'go test'
      }
    }
  }
}
