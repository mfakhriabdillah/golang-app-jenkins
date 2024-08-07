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
                echo 'Building application...'
                sh '''
                    go version
                    go help
                '''
            }
        }
    }
}
