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
                branch 'develop'
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

/*
    - run pipeline with when conditionn -> branch
    - agent Docker, Unitest
    - Intergration with Sonaquer, selenium, K6
*/