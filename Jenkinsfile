pipeline {

/*
agent {
        docker {
            image 'pdmlab/jenkins-node-docker-agent:6.11.1'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }
    */



    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage ("Build") {
            steps {
            sh "docker run hello-world"
            }
        }

        stage ("Deploy") {
            steps {
            sh "docker-compose up"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v"
        }
    }
}