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
                    GOCACHE=/tmp/ GOOS=linux GOARCH=amd64 go build -o golang-app
                '''
                stash includes: 'golang-app', name: 'GOAPP_ARTIFACT'
            }
        }
        stage('Deploy') {
            environment {
                SSH_KEY = credentials('SSHCRED')
            }
            steps {
                echo 'Deploying Application...'
                unstash name: 'GOAPP_ARTIFACT'
                sh '''
                    scp -o StrictHostKeyChecking=no -i $SSH_KEY golang-app faseero0@10.128.0.14:~/golang-app/

                '''
            }
        }
    }
}
