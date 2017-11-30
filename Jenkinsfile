pipeline {
  agent none
  stages {
    stage('Testing') {
     agent {
      docker {
          image 'golang'
      }
     } 
     steps {
        sh 'go test'
      }
    }
   stage('Deploy') {
        agent none
        steps {
            script {
                def customimage = docker.build("naos62/maths:${env.BUILD_NUMBER}")
            }
         }
        }
   }
  }
}
