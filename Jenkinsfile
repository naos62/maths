pipeline {
  agent none
  stages {
    stage('Test') {
      agent {
        docker {
        image 'golang'
        }
      }
      steps {
        sh 'go test'
      }
    }
  stage('Build') {
     agent none
    steps {
      script {
        def customImage = docker.build("bdubois/math:${env.BUILD_NUMBER}")
      }
    }
  }
  stage('push') {
    agent none
    steps {
      script {
        docker.withRegistry('https://registry.hub.docker.com', 'Docker') {
            customImage.push("${env.BUILD_NUMBER}")
            customImage.push("latest")
        }
      }
    }
  }
}
}