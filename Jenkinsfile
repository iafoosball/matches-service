pipeline {
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
    }
    post {
        always {
            sh "docker-compose down -v"
        }
    }
}