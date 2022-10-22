pipeline {
    agent any
    environment {
        dev = "develop"
        prod = "master"
    } 
    stages{
        stage("prepare env") {
            steps {
                sh 'make update'
            }
        }

        stage("build") {
           steps {
                sh 'docker build . -t account:latest --build-arg SERVICE_NAME=profile'
           }
        }

        stage("deploy") {
            steps {
                echo "deploy comming soon..."
            }
        }
    }
}