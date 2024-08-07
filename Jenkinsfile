pipeline {
    agent any

    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:alpine'
                    label 'docker'
                }
            }
            steps {
                echo 'Hello World'
                echo 'Building application...'
            }
        }
    }
}
