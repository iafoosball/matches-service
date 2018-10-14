pipeline {

    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage ("Prepare environment") {
            steps {
            sh "rm docker-compose.yml && rm Dockerfile"
            sh "cp ../iaf-configs/matches-service/stag/docker-compose.yml . && cp ../iaf-configs/stag/matches-service/Dockerfile ."
            }
        }
        stage ("Staging") {
            steps {
            sh "docker-compose build --pull"
            sh "docker-compose up --force-recreate -d"
            sh "cd matches/ && go test"
            sh "cd ../ && docker-compose down"
            }
        }
        stage ("Deploy") {
            steps {
            sh "docker-compose up --force-recreate"
            }
        }
    }
    post {
        always {
            sh "docker-compose down -v"
        }
    }
}