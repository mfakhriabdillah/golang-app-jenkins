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
                    GOOS=linux GOARCH=amd64 go build -o golang-app
                '''
            }
        }
    }
}
