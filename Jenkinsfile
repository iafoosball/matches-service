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
            sh "cp ../iaf-configs/matches-service/stag/docker-compose.yml . && cp ../iaf-configs/matches-service/stag/Dockerfile ."
            sh "docker-compose rm -f"
            }
        }
        stage ("Build") {
            sh "docker-compose build --pull"
            sh "docker cp matches-service:/root/matches.test ."
            sh "docker cp matches-service:/root/main ."
        }
        stage ("Test") {
            steps {
            sh "docker-compose up --force-recreate -d"
            sh "sleep 30s"
            sh "./matches.test"
            sh "docker-compose down"
            }
        }
        stage ("Production") {
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