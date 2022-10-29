def Onbranch(branch) {
    echo "Pipeline running on branch: ${branch}"
}

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
            when {
                branch 'master'
            }
           steps {
                sh 'docker build . -t account:latest --build-arg SERVICE_NAME=profile'
           }
        }

        stage("deploy") {
            input {
                message "Deploy to production?"
                id "simple-input"
            }
            steps {
                echo "deploy comming soon..."
            }
        }
    }
}